package recommender

import (
	"context"

	"github.com/ahmadnaufal/recommender-worker/model"
)

type RecommenderEngine struct {
	// stub tiktok provider
	tiktokProvider string
}

func New(tiktokProvider string) RecommenderEngine {
	return RecommenderEngine{tiktokProvider: tiktokProvider}
}

func (r RecommenderEngine) GenerateRecommendations(ctx context.Context, req RecommendationRequest) ([]model.Place, error) {
	placesToSort := req.Places
	// currentWeather := req.WeatherCondition

	return placesToSort, nil
}
