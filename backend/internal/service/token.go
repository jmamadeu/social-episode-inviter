package service

import (
	"context"

	"github.com/jmamadeu/episode-inviter.com/internal/data"
)

type Token struct {
	db          *data.Database
	userService *User
}

func NewTokenService(db *data.Database, user *User) *Token {
	return &Token{
		db,
		user,
	}
}

func (this *Token) CreateToken(ctx context.Context) {
	// insertQuery := `INSERT INTO tokens (id,) RETURNING *`
}
