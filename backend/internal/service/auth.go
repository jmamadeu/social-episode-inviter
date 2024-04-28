package service

import (
	"context"
	"fmt"

	"github.com/jmamadeu/episode-inviter.com/internal/model"
	util "github.com/jmamadeu/episode-inviter.com/internal/utils"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Auth struct {
	userService  *User
	tokenService *Token
}

func NewAuth(userService *User, tokenService *Token) *Auth {
	return &Auth{
		userService,
		tokenService,
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
		Name:    "Reitvin",
		Address: "joao.amadeu.coding@gmail.com",
	}, tokenMessage, "Authentication code")

	_, err = authService.tokenService.CreateToken(ctx,
		model.TokenTypeEmail,
		emailToken,
		user.Id,
	)
	if err != nil {
		return nil, err
	}

	err = tokenEmail.SendEmail()
	if err != nil {
		return nil, err
	}

	return user, nil
}
