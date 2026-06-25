package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	Port             string
	DATABASE_URL     string
	SHORT_URL_DOMAIN string
}

func LoadEnvData() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config := &AppConfig{
		Port:             os.Getenv("Port"),
		DATABASE_URL:     os.Getenv("DATABASE_URL"),
		SHORT_URL_DOMAIN: os.Getenv("SHORT_URL_DOMAIN"),
	}

	return config
}
