package model

import "github.com/google/uuid"

type Event struct {
	Id             uuid.UUID     `json:"id"`
	Name           string        `json:"name"`
	Location       string        `json:"location"`
	Description    string        `json:"description"`
	Audience       string        `json:"audience"`
	MediaChannelId uuid.UUID     `json:"media_channel_id"`
	MediaChannel   *MediaChannel `json:"media_channel"`
}
