package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id             uuid.UUID     `json:"id"`
	Name           string        `json:"name"`
	Location       string        `json:"location"`
	Date           time.Time     `json:"date"`
	Description    string        `json:"description"`
	MediaChannelId uuid.UUID     `json:"media_channel_id"`
	MediaChannel   *MediaChannel `json:"media_channel"`
}
