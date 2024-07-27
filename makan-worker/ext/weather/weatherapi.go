package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func (w WeatherAPIClient) GetWeather(ctx context.Context, latitude, longitude float64) (WeatherData, error) {
	// Format the URL with the provided API key and coordinates
	var weatherResp WeatherData
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%f,%f&key=%s", latitude, longitude, w.apikey)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := w.client.Do(req)
	if err != nil {
		return weatherResp, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return weatherResp, err
	}

	return weatherResp, nil
}
