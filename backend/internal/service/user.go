package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
)

type User struct {
	db *data.Database
}

func NewUser(database *data.Database) *User {
	return &User{
		database,
	}
}

func (userService *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT * FROM USERS WHERE email = $1`
	var user model.User
	err := userService.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, errors.New("no user was found with this email")
	}

	return &user, nil
}

func (userService *User) CreateUser(ctx context.Context, email string) (*model.User, error) {
	query := `INSERT INTO users (id,email) VALUES($1,$2) RETURNING *`
	var user model.User
	err := userService.db.QueryRow(ctx, query, uuid.New(), email).Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userService *User) FindOrCreateUser(ctx context.Context, email string) (*model.User, error) {
	user, err := userService.GetUserByEmail(ctx, email)
	if err != nil {
		user, err = userService.CreateUser(ctx, email)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (userService *User) GetUserById(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var user model.User
	err := userService.db.QueryRow(ctx, query, userId).Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
