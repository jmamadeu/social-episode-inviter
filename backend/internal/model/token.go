package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenType string

const (
	Email TokenType = "email"
	Api   TokenType = "api"
)

type Token struct {
	Id         uuid.UUID
	Created_at time.Time
	Updated_at time.Time
	Type       TokenType
	EmailToken string
	Valid      bool
	Expiration time.Time
	UserId     uuid.UUID
	User       *User
}
