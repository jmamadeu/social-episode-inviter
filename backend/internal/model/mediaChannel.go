package model

import (
	"errors"

	"github.com/google/uuid"
)

type MediaChannel struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	BannerUrl   string    `json:"bannerUrl"`
	Description string    `json:"description"`
	Username    string    `json:"username"`
	OwnerId     uuid.UUID `json:"ownerId"`
	Owner       *User     `json:"owner"`
}

var (
	ErrMediaChannelUsernameTaken = errors.New("media channel username is already taken")
	ErrMediaChannelNotFound      = errors.New("media channel not found")
)
