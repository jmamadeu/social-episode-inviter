package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmamadeu/episode-inviter.com/internal/service"
)

type Auth struct {
	authService *service.Auth
}

func NewAuth(us *service.Auth) *Auth {
	return &Auth{
		us,
	}
}

type authenticateRequest struct {
	Email string `json:"email" binding:"email"`
}

func (authHandler *Auth) Login(ctx *gin.Context) {
	var requestBody authenticateRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	user, err := authHandler.authService.Token(ctx, requestBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(
			err.Error(),
		))

		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Payload: user,
		Message: "Operation completed successfully",
	})
}
