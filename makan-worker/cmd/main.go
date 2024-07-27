package main

import (
	"context"
	"fmt"

	"github.com/ahmadnaufal/recommender-worker/ext/location"
	"github.com/ahmadnaufal/recommender-worker/ext/weather"
	"github.com/ahmadnaufal/recommender-worker/server"
)

func main() {
	cfg, err := server.LoadConfig()
	if err != nil {
		panic(err)
	}

	locationProv := location.New(cfg.Google.Host, cfg.Google.APIKey)
	weatherProv := weather.New(cfg.Weather.Host, cfg.Weather.APIKey)

	latitude := -6.291456
	longitude := 106.7840557
	radius := 500.0
	var numOfRecommendation uint = 10
	ctx := context.Background()
	places, err := locationProv.GetNearby(ctx, latitude, longitude, radius, numOfRecommendation)
	if err != nil {
		panic(err)
	}

	for _, p := range places {
		fmt.Println(p.String())
	}
}
