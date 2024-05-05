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

func (mediaChannelService *MediaChannel) CreateMediaChannel(ctx context.Context, name, bannerUrl, description, username string, ownerId uuid.UUID) (*model.MediaChannel, error) {
	query := `INSERT INTO media_channels
		(id,name,banner_url,description,username,owner_id)
		VALUES (
			$1,$2,$3,$4,$5,$6
		)
	`
	mediaChannel := &model.MediaChannel{
		Id:          uuid.New(),
		Name:        name,
		BannerUrl:   bannerUrl,
		Description: description,
		Username:    username,
		OwnerId:     ownerId,
	}

	_, err := mediaChannelService.db.Exec(ctx, query,
		mediaChannel.Id,
		mediaChannel.Name,
		mediaChannel.BannerUrl,
		mediaChannel.Description,
		mediaChannel.Username,
		mediaChannel.OwnerId,
	)
	if err != nil {
		return nil, err
	}

	user, err := mediaChannelService.userService.GetUserById(ctx, ownerId)
	if err != nil {
		return nil, err
	}
	mediaChannel.Owner = user

	return mediaChannel, nil
}
