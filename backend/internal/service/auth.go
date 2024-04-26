package service

import (
	"context"
	"fmt"

	"github.com/jmamadeu/episode-inviter.com/internal/model"
	util "github.com/jmamadeu/episode-inviter.com/internal/utils"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Auth struct {
	userService *User
}

func NewAuth(userService *User) *Auth {
	return &Auth{
		userService,
	}
}

func (authService *Auth) Token(ctx context.Context, email string) (*model.User, error) {
	user, err := authService.userService.FindOrCreateUser(ctx, email)
	if err != nil {
		return nil, err
	}

	emailToken := util.GenerateDigits()

	tokenMessage := fmt.Sprintf("Your authentication token is: %d", emailToken)

	tokenEmail := NewEmail(mail.Email{
		Name:    "João Amadeu",
		Address: "joao.amadeu.coding@gmail.com",
	}, mail.Email{
		Name:    "Episode Inviter",
		Address: "joao.amadeu.coding@gmail.com",
	}, tokenMessage, "Retivini Authentication code")

	err = tokenEmail.SendEmail()
	if err != nil {
		return nil, err
	}

	return user, nil
}
