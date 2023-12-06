package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NewsAPIKey string
	NewsAPIURL string
}

func NewConfig() (*Config, error) {
	godotenv.Load(".env")

	return &Config{
		NewsAPIKey: getEnv("NEWS_API_KEY", ""),
		NewsAPIURL: getEnv("NEWS_API_URL", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
