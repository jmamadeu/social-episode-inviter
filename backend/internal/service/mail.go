package service

import (
	"fmt"

	"github.com/jmamadeu/episode-inviter.com/internal/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	to      mail.Email
	from    mail.Email
	message string
	subject string
}

type Mailer mail.Email

func NewEmail(to, from mail.Email, message, subject string) *Email {
	return &Email{
		to,
		from,
		message,
		subject,
	}
}

func (emailService *Email) SendEmail() error {
	appConfig := config.New()

	message := mail.NewSingleEmail(&emailService.from, emailService.subject, &emailService.to, emailService.message, "")
	client := sendgrid.NewSendClient(appConfig.Email.SendGridApiKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	fmt.Println(response.StatusCode)

	return nil
}
