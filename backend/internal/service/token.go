package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
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

func (tokenService *Token) CreateToken(ctx context.Context, tokenType model.TokenType, emailToken int, userId uuid.UUID) (*model.Token, error) {
	insertQuery := `INSERT INTO tokens  
		(id,created_at,updated_at,type,email_token, valid,expiration,user_id) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8);`

	token := &model.Token{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Type:       tokenType,
		EmailToken: fmt.Sprint(emailToken),
		Valid:      true,
		Expiration: time.Now().Add(time.Minute + 10),
		UserId:     userId,
	}

	_, err := tokenService.db.Exec(ctx, insertQuery,
		token.Id,
		token.CreatedAt,
		token.UpdatedAt,
		token.Type,
		token.EmailToken,
		token.Valid,
		token.Expiration,
		token.UserId,
	)
	if err != nil {
		return nil, err
	}

	user, err := tokenService.userService.GetUserById(ctx, userId)
	if err != nil {
		println("Error getitng the user back")
		return nil, err
	}

	token.User = user

	return token, nil
}
