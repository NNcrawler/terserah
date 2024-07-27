package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmadnaufal/recommender-worker/model"
)

type WeatherAPIClient struct {
	client *http.Client
	host   string
	apikey string
}

func New(host, apikey string) WeatherAPIClient {
	return WeatherAPIClient{
		client: &http.Client{},
		host:   host,
		apikey: apikey,
	}
}

func (w WeatherAPIClient) GetWeather(ctx context.Context, latitude, longitude float64) (model.CurrentWeather, error) {
	// Format the URL with the provided API key and coordinates
	var weatherModel model.CurrentWeather
	url := fmt.Sprintf("%s/current.json?q=%f,%f&key=%s", w.host, latitude, longitude, w.apikey)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := w.client.Do(req)
	if err != nil {
		return weatherModel, err
	}
	defer resp.Body.Close()

	var weatherResp WeatherData
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return weatherModel, err
	}

	return weatherRespToWeatherModel(weatherResp), nil
}

func weatherRespToWeatherModel(w WeatherData) model.CurrentWeather {
	return model.CurrentWeather{
		Time:        time.Unix(w.Location.LocaltimeEpoch, 0),
		Temperature: w.Current.TempC,
		FeelsLike:   w.Current.FeelslikeC,
	}
}
