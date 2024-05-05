package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/service"
)

type MediaChannel struct {
	mediaChannelService *service.MediaChannel
}

func NewMediaChannelHandler(mediaChannel *service.MediaChannel) *MediaChannel {
	return &MediaChannel{
		mediaChannel,
	}
}

type createChannelRequest struct {
	Name        string    `json:"name" binding:"required"`
	BannerUrl   string    `json:"bannerUrl"`
	Description string    `json:"description" binding:"required"`
	Username    string    `json:"username" binding:"required"`
	OwnerId     uuid.UUID `json:"ownerId" binding:"required"`
}

func (mcs *MediaChannel) CreateNewMediaChannel(ctx *gin.Context) {
	var requestBody createChannelRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	mediaChannel, err := mcs.mediaChannelService.CreateMediaChannel(ctx, requestBody.Name, requestBody.BannerUrl, requestBody.Description, requestBody.Username, requestBody.OwnerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse(mediaChannel))
}
