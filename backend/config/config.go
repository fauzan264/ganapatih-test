package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey string
	AppPort string
	AppHost string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		SecretKey: os.Getenv("SECRET_KEY"),
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
	}

	if cfg.SecretKey == "" {
		return nil, fmt.Errorf("SECRET_KEY is required")
	}

	if cfg.AppHost == "" {
		cfg.AppHost = "localhost"
	}

	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}

	return cfg, nil
}