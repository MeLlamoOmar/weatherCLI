package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"omar/sun/internals"
	"strings"
)

type WeatherService struct {}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s WeatherService) GetWeatherData(apiKey, location string) (*internals.Weather, error) {
	apiURL := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=yes&alerts=no", apiKey, strings.TrimRight(location, " "))
	res, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get valid weather data: %s", res.Status)
	}

	var weather internals.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %v", err)
	}

	return &weather, nil
}