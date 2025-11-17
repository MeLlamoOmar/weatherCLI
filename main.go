package main

import (
	"fmt"
	"log"
	"math"
	"omar/sun/internals"
	"omar/sun/internals/services"
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

	service := services.NewWeatherService()
	weather, err := service.GetWeatherData(config.ApiKey, location)
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Printf("%s, %s: %.1fC, %s\n", weather.Location.Name, weather.Location.Country, weather.Current.TempC, weather.Current.Condition.Text)
	for day, hours := range weather.Forecast.ForecastDay {
		for _, hour := range hours.Hour {
			date := time.Unix(hour.TimeEpoch, 0)
			t := ""
			redChance := int(255 * hour.ChanceOfRain / 100)
			red := color.RGB(redChance, 0, 0).SprintFunc()

			if hour.TempC < 30 {
				tBlue := int(255 * math.Abs(hour.TempC/100-1))
				blue := color.RGB(0, 0, tBlue).SprintFunc()
				t = blue(hour.TempC, "C")
			} else {
				tYellow := int(255 * math.Abs(hour.TempC/100-1))
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
