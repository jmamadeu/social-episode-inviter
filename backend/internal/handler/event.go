package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/service"
)

type Event struct {
	eventService *service.Event
}

func NewEventHandler(eventService *service.Event) *Event {
	return &Event{
		eventService: eventService,
	}
}

type createEventRequest struct {
	Name           string    `json:"name" binding:"required"`
	Location       string    `json:"location" binding:"required"`
	Date           time.Time `json:"date" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	MediaChannelId uuid.UUID `json:"mediaChannelId" binding:"required"`
}

func (es *Event) CreateEvent(ctx *gin.Context) {
	var requestBody createEventRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	event, err := es.eventService.CreateNewEvent(ctx, requestBody.Name, requestBody.Location, requestBody.Description, requestBody.Date, requestBody.MediaChannelId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse(event))
}
