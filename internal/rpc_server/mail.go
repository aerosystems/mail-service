package RPCServer

import (
	"github.com/aerosystems/mail-service/internal/provider"
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"log"
	"os"
	"strconv"
)

type MailServer struct{}

type MailPayload struct {
	To      string
	Subject string
	Body    string
}

func (r *MailServer) SendEmail(payload MailPayload, resp *string) error {
	log.Println("sending email to", payload.To)
	mailService := &MailService.MailService{}
	msg := MailService.Message{
		To:       payload.To,
		ToName:   "Customer",
		FromName: "Testmail💎",
		From:     "no-reply@testmail.top",
		Subject:  payload.Subject,
		Body:     payload.Body,
	}
	switch os.Getenv("EMAIL_PROVIDER") {
	case "mailhog":
		port, err := strconv.Atoi(os.Getenv("MAILHOG_PORT"))
		if err != nil {
			log.Fatal(err)
		}
		mailService.SetProvider(&provider.SMTP{
			Domain:     os.Getenv("MAILHOG_DOMAIN"),
			Host:       os.Getenv("MAILHOG_HOST"),
			Port:       port,
			Username:   os.Getenv("MAILHOG_USERNAME"),
			Password:   os.Getenv("MAILHOG_PASSWORD"),
			Encryption: os.Getenv("MAILHOG_ENCRYPTION"),
		})
	case "brevo":
		mailService.SetProvider(&provider.Brevo{})
	case "postfix":
		mailService.SetProvider(&provider.SMTP{})
	default:
		log.Fatal("No email provider set")
	}
	if err := mailService.SendEmail(msg); err != nil {
		return err
	}

	*resp = "Sent email to " + payload.To
	return nil
}
