package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
	util "github.com/jmamadeu/episode-inviter.com/internal/utils"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

func (userService *User) Token(ctx context.Context, email string) (*model.User, error) {
	user, err := userService.FindOrCreateUser(ctx, email)
	if err != nil {
		return nil, err
	}

	emailTokenDigits := util.GenerateDigits()

	tokenMessage := fmt.Sprintf("Your authentication token is: %d", emailTokenDigits)

	tokenEmail := NewEmail(mail.Email{
		Name:    "João Amadeu",
		Address: "joaomateusamadeu@gmail.com",
	}, mail.Email{
		Name:    "Episode Inviter",
		Address: "episode_inviter@geral.com",
	}, "Your ", tokenMessage)

	tokenEmail.SendEmail()

	return user, nil
}
