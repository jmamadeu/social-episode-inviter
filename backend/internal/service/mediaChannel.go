package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
)

type MediaChannel struct {
	db          *data.Database
	userService *User
}

func NewMediaChannelService(db *data.Database, userService *User) *MediaChannel {
	return &MediaChannel{
		db,
		userService,
	}
}

func (mediaChannelService *MediaChannel) CreateMediaChannel(ctx context.Context, name, url, bannerUrl string, platform model.MediaChannelPlatform, userId uuid.UUID) (*model.MediaChannel, error) {
	query := `INSERT INTO media_channels
		(id,name,platform,url,banner_url,user_id)
		VALUES (
			$1,$2,$3,$4,$5,$6
		)
	`
	mediaChannel := &model.MediaChannel{
		Id:        uuid.New(),
		Name:      name,
		Platform:  platform,
		Url:       url,
		BannerUrl: bannerUrl,
		UserId:    userId,
	}

	_, err := mediaChannelService.db.Exec(ctx, query,
		mediaChannel.Id,
		mediaChannel.Name,
		mediaChannel.Platform,
		mediaChannel.Url,
		mediaChannel.BannerUrl,
		mediaChannel.UserId,
	)
	if err != nil {
		return nil, err
	}

	user, err := mediaChannelService.userService.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	mediaChannel.User = user

	return mediaChannel, nil
}
