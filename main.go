package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"omar/sun/internals"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config := internals.LoadEnv()
	location := "Santo+Domingo"
	if len(os.Args) >= 2 {
		location = strings.Join(os.Args[1:], "+")
	}

	apiURL := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=yes&alerts=no", config.ApiKey, strings.TrimRight(location, " "))
	res, err := http.Get(apiURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch weather data: %v", err))
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read response body: %v", err))
	}

	if res.StatusCode != 200 {
		panic("Failed to get valid weather data: " + res.Status)
	}

	var weather internals.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse weather data: %v", err))
	}

	l, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour
	localTime := time.Unix(l.LocalTime, 0)

	fmt.Printf("%s, %s: %.1fC, %s\n", l.Name, l.Country, current.TempC, current.Condition.Text)
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(localTime) {
			continue
		}

		if hour.ChanceOfRain > 50 {
			color.Red(" %s - %.1fC %s %.0f%%\n", date.Format("15:04"), hour.TempC, hour.Condition.Text, hour.ChanceOfRain)
			continue
		}

		fmt.Printf(" %s - %.1fC %s %.0f%%\n", date.Format("15:04"), hour.TempC, hour.Condition.Text, hour.ChanceOfRain)
	}
}
