package model

import "github.com/google/uuid"

type MediaChannelPlatform string

var MediaChannelPlatformYoutube MediaChannelPlatform = "youtube"

type MediaChannel struct {
	Id        uuid.UUID
	Name      string
	Platform  MediaChannelPlatform
	Url       string
	BannerUrl string
	UserId    uuid.UUID
	User      *User
}
