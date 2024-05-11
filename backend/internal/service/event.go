package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
)

type Event struct {
	db           *data.Database
	mediaChannel *MediaChannel
}

func NewEventService(db *data.Database, mediaChannel *MediaChannel) *Event {
	return &Event{
		db,
		mediaChannel,
	}
}

func (es *Event) CreateNewEvent(ctx context.Context, name, location, description string, date time.Time, mediaChannelId uuid.UUID) (*model.Event, error) {
	query := `INSERT INTO events (id, name, location, date, description, media_channel_id) 
		VALUES($1,$2,$3,$4,$5,$6)
	`
	event := &model.Event{
		Id:             uuid.New(),
		Name:           name,
		Location:       location,
		Date:           date,
		Description:    description,
		MediaChannelId: mediaChannelId,
	}

	_, err := es.db.Exec(ctx, query,
		event.Id,
		event.Name,
		event.Location,
		event.Date,
		event.Description,
		event.MediaChannelId,
	)
	if err != nil {
		return nil, err
	}

	return event, nil
}
