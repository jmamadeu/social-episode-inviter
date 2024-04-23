package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Email struct {
	SendGridApiKey string
}

type AppConfig struct {
	DatabaseUrl string
	Email       Email
}

func New() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading config vars")
	}

	return &AppConfig{
		DatabaseUrl: os.Getenv("DB_URL"),
		Email: Email{
			SendGridApiKey: os.Getenv("SENDGRID_API_KEY"),
		},
	}
}
