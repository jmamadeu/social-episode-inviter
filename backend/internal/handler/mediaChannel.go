package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
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

type mediaChannelResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	BannerUrl   string    `json:"bannerUrl"`
	Description string    `json:"description"`
	Username    string    `json:"username"`
	OwnerId     uuid.UUID `json:"ownerId"`
}

func newMediaChannelResponse(data []model.MediaChannel) (mediaChannels []mediaChannelResponse) {
	for m := range data {
		mediaChannels = append(mediaChannels, mediaChannelResponse{
			data[m].Id,
			data[m].Name,
			data[m].BannerUrl,
			data[m].Description,
			data[m].Username,
			data[m].OwnerId,
		})
	}

	return mediaChannels
}

func (mcs *MediaChannel) FetchMediaChannels(ctx *gin.Context) {
	mediaChannels, err := mcs.mediaChannelService.FetchMediaChannels(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse(newMediaChannelResponse(mediaChannels)))
}
