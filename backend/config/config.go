package config

import (
	"os"
)

type Config struct {
	Port         string
	DatabaseURL  string
	Environment  string
}

func LoadConfig() (*Config, error) {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", "memory"), // Use "memory" for local development
		Environment:  getEnv("ENVIRONMENT", "development"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}