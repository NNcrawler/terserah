package recommender

import "github.com/ahmadnaufal/recommender-worker/model"

type RecommendationRequest struct {
	Places           []model.Place
	WeatherCondition model.CurrentWeather
}
