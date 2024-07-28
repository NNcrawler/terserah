package makanworker

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ahmadnaufal/recommender-pipeline/ext/location"
	"github.com/ahmadnaufal/recommender-pipeline/ext/openai"
	"github.com/ahmadnaufal/recommender-pipeline/model"
	"github.com/ahmadnaufal/recommender-pipeline/recommender"
	"github.com/ahmadnaufal/recommender-pipeline/repo"
	"github.com/ahmadnaufal/recommender-pipeline/server"
	"github.com/google/uuid"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("BatchInsertRecommendations", BatchInsertRecommendations)
}

type PlaceResponse struct {
	Name       string            `json:"name"`
	Tags       []string          `json:"tags"`
	DishType   []string          `json:"dishType"`
	PriceLevel string            `json:"priceLevel"`
	Location   PlaceLocation     `json:"location"`
	Reviews    map[string]string `json:"reviews"`
}

type PlaceLocation struct {
	GoogleMaps string `json:"googleMaps"`
	Address    string `json:"address"`
}

type BaseResponse struct {
	Data any `json:"data"`
}

func BatchInsertRecommendations(w http.ResponseWriter, r *http.Request) {
	shouldReturn := handleCors(w, r)
	if shouldReturn {
		return
	}

	cfg, err := server.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := server.ConnectToDB(cfg.Database)
	locationProv := location.New(cfg.Google.Host, cfg.Google.APIKey)
	recommenderEngine := recommender.New("test")
	locationRepo := repo.New(db)

	ctx := r.Context()

	q := r.URL.Query()
	latitude, _ := strconv.ParseFloat(q.Get("latitude"), 64)
	longitude, _ := strconv.ParseFloat(q.Get("longitude"), 64)
	numOfRecommendationsBatch, _ := strconv.ParseUint(q.Get("n"), 10, 64)

	places, err := locationProv.GetNearby(ctx, latitude, longitude, 500.0, uint(numOfRecommendationsBatch))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	err = enrichWithDishTypes(ctx, cfg, places)
	if err != nil {
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	err = enrichWithReviews(ctx, cfg, places)
	if err != nil {
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	placesToRecommend, err := recommenderEngine.GenerateRecommendations(r.Context(), recommender.RecommendationRequest{
		Places: places,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	// save places
	for _, p := range placesToRecommend {
		p.ID = uuid.NewString()
		err = locationRepo.InsertPlace(ctx, p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error:", err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d places processed.", len(placesToRecommend))
}

func handleCors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

func enrichWithDishTypes(ctx context.Context, cfg server.Config, places []model.Place) error {
	openAiProv := openai.NewDishExtractor(cfg.OpenAI.APIKey)
	for i := 0; i < len(places); i++ {
		dishType, err := openAiProv.GetPossibleFoodsFromPlace(ctx, places[i].Reviews)
		if err != nil {
			return err
		}
		places[i].DishType = dishType
	}
	return nil
}

func enrichWithReviews(ctx context.Context, cfg server.Config, places []model.Place) error {
	reviewSummarizer := openai.ReviewSummarizer{
		ApiKey: cfg.OpenAI.APIKey,
	}
	for i := 0; i < len(places); i++ {
		reviewSummary, err := reviewSummarizer.AsReviewer(places[i].Reviews)
		if err != nil {
			return err
		}
		places[i].SummaryReviewFood = reviewSummary.Food
		places[i].SummaryReviewPlace = reviewSummary.Place
	}
	return nil
}
