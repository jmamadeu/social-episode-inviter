package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmamadeu/episode-inviter.com/internal/service"
)

type Auth struct {
	userService *service.User
}

func NewAuth(us *service.User) *Auth {
	return &Auth{
		us,
	}
}

type authenticateRequest struct {
	Email string `json:"email" binding:"email"`
}

func (authHandler *Auth) Authenticate(ctx *gin.Context) {
	var requestBody authenticateRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	fmt.Print(requestBody.Email)

	_, err := authHandler.userService.GetUserByEmail(ctx, requestBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		os.Exit(1)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "everything ok",
	})
}
