package internals

import "os"

type Config struct {
	ApiKey string
}

func LoadEnv() *Config {
	apiKey := GetEnv("WEATHER_API_KEY", "weatherApiKey")

	return &Config{
		ApiKey: apiKey,
	}
}

func GetEnv(key, fallback string) string {
	k := os.Getenv(key)
	if k == "" {
		return fallback
	}

	return k
}
