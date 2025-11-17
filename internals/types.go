package internals

type Weather struct {
	Location WeatherLocation `json:"location"`
	Current  WeatherCurrent  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	Hour []WeatherForecastHour `json:"hour"`
}

type WeatherLocation struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type WeatherCondition struct {
	Text string `json:"text"`
}

type WeatherCurrent struct {
	TempC     float64 `json:"temp_c"`
	Condition WeatherCondition `json:"condition"`
}

type WeatherForecastHour struct {
	TimeEpoch int64   `json:"time_epoch"`
	TempC     float64 `json:"temp_c"`
	Condition WeatherCondition `json:"condition"`
	ChanceOfRain float64 `json:"chance_of_rain"`
}
