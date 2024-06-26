package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (es *Event) FetchEvents(ctx context.Context) ([]model.Event, error) {
	query := `SELECT * FROM events`
	var events []model.Event
	var event model.Event

	rows, _ := es.db.Query(ctx, query)
	pgx.ForEachRow(rows, []any{
		&event.Id,
		&event.Name,
		&event.Location,
		&event.Description,
		&event.MediaChannelId,
		&event.Date,
	}, func() error {
		events = append(events, event)
		return nil
	})

	fmt.Print(events)

	return events, nil
}
