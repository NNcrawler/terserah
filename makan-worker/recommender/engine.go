package recommender

import (
	"context"
	"sort"

	"github.com/ahmadnaufal/recommender-worker/model"
)

type RecommenderEngine struct {
	ratingWeight            float64
	numberOfReviewWeight    float64
	distanceWeight          float64
	currentWeatherWeight    float64
	numberOfReviewMaxAmount int
}

func New() RecommenderEngine {
	return RecommenderEngine{
		ratingWeight:            0.3,
		numberOfReviewWeight:    0.3,
		distanceWeight:          0.1,
		currentWeatherWeight:    0.3,
		numberOfReviewMaxAmount: 300,
	}
}

// recommendation should follow this rule:
// google maps:
// rating = 0.3
// number of review = 0.3
// distance = 0.1
// current weather = 0.3
func (r RecommenderEngine) GenerateRecommendations(ctx context.Context, req RecommendationRequest) ([]model.Place, error) {
	placeSorted := req.Places
	currentWeather := req.WeatherCondition
	n := len(placeSorted)

	curCondition := currentWeather.DetermineCondition()
	var foodChoices []string

	// Add temperature-based food choices
	if choices, ok := temperatureToFoodChoiceMap[curCondition.Temperature]; ok {
		foodChoices = append(foodChoices, choices...)
	}

	// Add time of day-based food choices
	if choices, ok := timeOfDayToFoodChoiceMap[curCondition.Time]; ok {
		foodChoices = append(foodChoices, choices...)
	}

	// Add condition-based food choices if applicable
	if choices, ok := conditionToFoodChoiceMap[curCondition.Condition]; ok {
		foodChoices = append(foodChoices, choices...)
	}

	// Deduplicate food choices
	foodChoices = deduplicate(foodChoices)

	for i := 0; i < n; i++ {
		// by default, places are sorted by distance, so we apply the distance score
		placeSorted[i].TotalScore += (float64(n-i) / float64(n)) * r.distanceWeight

		// apply rating score
		placeSorted[i].TotalScore += (placeSorted[i].Rating * 2) * r.ratingWeight

		// apply number of review score,
		numberOfReviewNormalized := placeSorted[i].UserRatingCount
		if numberOfReviewNormalized > r.numberOfReviewMaxAmount {
			numberOfReviewNormalized = r.numberOfReviewMaxAmount
		}
		divider := float64(r.numberOfReviewMaxAmount) / 10

		placeSorted[i].TotalScore += (float64(numberOfReviewNormalized) / divider) * r.ratingWeight

		// TODO: apply scoring on current weather
		// a bit complex since we need to map each dishType into its perfect weather
		placeSorted[i].TotalScore += r.calculateCurrentWeatherScore(placeSorted[i], foodChoices)
	}

	// in the end, the response should be sorted by the calculated score
	sort.Slice(placeSorted, func(i, j int) bool {
		return placeSorted[i].TotalScore > placeSorted[j].TotalScore
	})

	return placeSorted, nil
}

func deduplicate(items []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// FindIntersection returns the intersection of two slices
func findIntersection(slice1, slice2 []string) []string {
	elementMap := make(map[string]bool)
	for _, elem := range slice1 {
		elementMap[elem] = true
	}

	var intersection []string
	for _, elem := range slice2 {
		if _, found := elementMap[elem]; found {
			intersection = append(intersection, elem)
		}
	}

	return intersection
}

func (r RecommenderEngine) calculateCurrentWeatherScore(place model.Place, foodChoices []string) float64 {
	matchedFoods := findIntersection(place.DishType, foodChoices)
	return float64(len(matchedFoods)) / float64(len(foodChoices)) * 10 * r.currentWeatherWeight
}
