package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	*pgxpool.Pool
	url string
}

func NewDatabase(ctx context.Context, url string) (*Database, error) {
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Database{
		db,
		url,
	}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
}
