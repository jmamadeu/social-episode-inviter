package model

import (
	"time"

	"github.com/google/uuid"
)

const EMAIL_TOKEN_EXPIRATION_MINUTES = 60 * 10

type TokenType string

const (
	TokenTypeEmail TokenType = "email"
	TokenTypeApi   TokenType = "api"
)

type Token struct {
	Id         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Type       TokenType
	EmailToken string
	Valid      bool
	Expiration time.Time
	UserId     uuid.UUID
	User       *User
}
