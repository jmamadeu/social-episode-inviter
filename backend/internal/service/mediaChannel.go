package service

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	usernameTaken := mediaChannelService.CheckMediaChannelExistsByUsername(ctx, username)
	if usernameTaken {
		return nil, model.ErrMediaChannelUsernameTaken
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

func (mcs *MediaChannel) CheckMediaChannelExistsByUsername(ctx context.Context, username string) bool {
	query := `SELECT COUNT(*) FROM media_channels WHERE username = $1`
	var mediaChannelCount string
	mcs.db.QueryRow(ctx, query, username).Scan(&mediaChannelCount)

	count, err := strconv.Atoi(mediaChannelCount)
	if err != nil {
		return false
	}

	return count > 0
}

func (mcs *MediaChannel) GetMediaChannelById(ctx context.Context, id uuid.UUID) (*model.MediaChannel, error) {
	query := `SELECT * FROM media_channel WHERE id = $1`
	var mediaChannel model.MediaChannel

	mcs.db.QueryRow(ctx, query, mediaChannel, id).Scan(
		&mediaChannel.Id,
		&mediaChannel.Name,
		&mediaChannel.Description,
		&mediaChannel.BannerUrl,
		&mediaChannel.OwnerId,
		&mediaChannel.Username,
	)

	return &mediaChannel, nil
}

func (mcs *MediaChannel) FetchMediaChannels(ctx context.Context) ([]model.MediaChannel, error) {
	query := `SELECT * FROM media_channels`
	var mediaChannels []model.MediaChannel

	var mediaChannel model.MediaChannel

	rows, _ := mcs.db.Query(ctx, query)
	pgx.ForEachRow(rows, []any{
		&mediaChannel.Id,
		&mediaChannel.Name,
		&mediaChannel.Description,
		&mediaChannel.BannerUrl,
		&mediaChannel.OwnerId,
		&mediaChannel.Username,
	}, func() error {
		mediaChannels = append(mediaChannels, mediaChannel)
		return nil
	})

	return mediaChannels, nil
}
