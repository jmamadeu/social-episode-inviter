package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmamadeu/episode-inviter.com/internal/config"
	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/handler"
	"github.com/jmamadeu/episode-inviter.com/internal/service"
)

type RequestBody struct {
	Email string `json:"email" binding:"required"`
}

func main() {
	appConfig := config.New()
	ctx := context.Background()
	db, err := data.NewDatabase(ctx, appConfig.DatabaseUrl)
	if err != nil {
		slog.Error("Error initializing the database connections", "", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected with the database")

	router := gin.Default()

	userService := service.NewUser(db)
	tokenService := service.NewTokenService(db, userService)

	authService := service.NewAuth(userService, tokenService)
	authHandler := handler.NewAuth(authService)
	router.POST("/api/v1/login", authHandler.Login)

	router.Run(":3333")
}
