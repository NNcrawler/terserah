package makanworker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ahmadnaufal/recommender-worker/ext/location"
	"github.com/ahmadnaufal/recommender-worker/ext/openai"
	"github.com/ahmadnaufal/recommender-worker/ext/weather"
	"github.com/ahmadnaufal/recommender-worker/model"
	"github.com/ahmadnaufal/recommender-worker/recommender"
	"github.com/ahmadnaufal/recommender-worker/server"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetRecommendations", GetRecommendations)
}

type PlaceResponse struct {
	Name       string        `json:"name"`
	Tags       []string      `json:"tags"`
	DishType   []string      `json:"dishType"`
	PriceLevel string        `json:"priceLevel"`
	Location   PlaceLocation `json:"location"`
	Reviews    []string      `json:"reviews"`
}

type PlaceLocation struct {
	GoogleMaps string `json:"googleMaps"`
	Address    string `json:"address"`
}

type BaseResponse struct {
	Data any `json:"data"`
}

func GetRecommendations(w http.ResponseWriter, r *http.Request) {
	shouldReturn := handleCors(w, r)
	if shouldReturn {
		return
	}

	cfg, err := server.LoadConfig()
	if err != nil {
		panic(err)
	}

	locationProv := location.New(cfg.Google.Host, cfg.Google.APIKey)
	weatherProv := weather.New(cfg.Weather.Host, cfg.Weather.APIKey)
	openAiProv := openai.New(cfg.OpenAI.Host, cfg.OpenAI.APIKey)
	recommenderEngine := recommender.New("test")

	ctx := r.Context()

	q := r.URL.Query()
	latitude, _ := strconv.ParseFloat(q.Get("latitude"), 64)
	longitude, _ := strconv.ParseFloat(q.Get("longitude"), 64)

	places, err := locationProv.GetNearby(ctx, latitude, longitude, 500.0, 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	currentWeather, err := weatherProv.GetWeather(ctx, latitude, longitude)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	// populate dishes
	for i := 0; i < len(places); i++ {
		dishType, err := openAiProv.GetPossibleFoodsFromPlace(ctx, places[i])
		if err != nil {
			fmt.Fprint(w, "Error:", err.Error())
			return
		}
		places[i].DishType = dishType
	}

	placesToRecommend, err := recommenderEngine.GenerateRecommendations(r.Context(), recommender.RecommendationRequest{
		Places:           places,
		WeatherCondition: currentWeather,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	var placesResponse []PlaceResponse
	for _, placeToRecommend := range placesToRecommend {
		placesResponse = append(placesResponse, placeToResponse(placeToRecommend))
	}

	response := BaseResponse{Data: placesResponse}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func placeToResponse(place model.Place) PlaceResponse {
	return PlaceResponse{
		Name:       place.PlaceName,
		DishType:   place.DishType,
		Tags:       place.Tags,
		PriceLevel: place.PriceLevel,
		Location: PlaceLocation{
			GoogleMaps: place.GoogleMapsURI,
			Address:    place.Address,
		},
		Reviews: place.Reviews,
	}
}

func handleCors(w http.ResponseWriter, r *http.Request) bool {
	// Set CORS headers
	// Allow all origins
	// Allow specific methods
	// Allow specific headers
	// Handle preflight requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}
