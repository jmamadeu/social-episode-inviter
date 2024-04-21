package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DatabaseUrl string
}

func New() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading config vars")
	}

	return &AppConfig{
		DatabaseUrl: os.Getenv("DB_URL"),
	}
}
