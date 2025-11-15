package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
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

	l, current, forecastDay := weather.Location, weather.Current, weather.Forecast.ForecastDay

	fmt.Printf("%s, %s: %.1fC, %s\n", l.Name, l.Country, current.TempC, current.Condition.Text)
	for day, hours := range forecastDay {
		for _, hour := range hours.Hour {
			date := time.Unix(hour.TimeEpoch, 0)
			t := ""
			redChance := int(255*hour.ChanceOfRain/100)
			red := color.RGB(redChance,0,0).SprintFunc()

			if hour.TempC < 30 {
				tBlue := int(255 * math.Abs(hour.TempC / 100 - 1))
				blue := color.RGB(0, 0, tBlue).SprintFunc()
				t = blue(hour.TempC, "C")
			} else {
				tYellow := int(255 * math.Abs(hour.TempC / 100 - 1))
				yellow := color.RGB(tYellow, tYellow, 0).SprintFunc()
				t = yellow(hour.TempC, "C")
			}
			
			if date.Before(time.Now()) {
				continue
			}

			if day == 0 {
				fmt.Printf("%s - %s %s %s \n", date.Format("15:04"), t, hour.Condition.Text, red(hour.ChanceOfRain, "%"))
				continue
			} else if date.Format("15:04") == "00:00" {
				fmt.Printf("%s - %s %s %s \n", date.Format("15:04"), t, hour.Condition.Text, red(hour.ChanceOfRain, "%"))
				continue
			}
		}
	} 
}
