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

	mediaChannelService := service.NewMediaChannelService(db, userService)
	mediaChannelHandler := handler.NewMediaChannelHandler(mediaChannelService)

	authService := service.NewAuth(userService, tokenService)
	authHandler := handler.NewAuth(authService)

	v1Router := router.Group("/api/v1")
	{
		v1Router.POST("/login", authHandler.Login)
		v1Router.POST("/media-channels", mediaChannelHandler.CreateNewMediaChannel)
		v1Router.GET("/media-channels", mediaChannelHandler.FetchMediaChannels)
	}

	router.Run(":3333")
}
