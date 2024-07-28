package recommender

import "github.com/ahmadnaufal/recommender-pipeline/model"

type RecommendationRequest struct {
	Places           []model.Place
	WeatherCondition model.CurrentWeather
}
