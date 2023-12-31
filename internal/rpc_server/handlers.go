package RPCServer

import (
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"os"
)

type MailPayload struct {
	To      string
	Subject string
	Body    string
}

func (ms *MailServer) SendEmail(payload MailPayload, resp *string) error {
	ms.log.Infof("sending email to %s", payload.To)

	msg := MailService.Message{
		To:       payload.To,
		ToName:   "Customer",
		FromName: "Testmail💎",
		From:     "no-reply@testmail.top",
		Subject:  payload.Subject,
		Body:     payload.Body,
	}

	ms.log.Infof("sending email with %s provider", os.Getenv("EMAIL_PROVIDER"))
	if err := ms.mailService.SendEmail(msg); err != nil {
		ms.log.Errorf("error sending email: %v", err)
		return err
	}

	*resp = "Sent email to " + payload.To
	return nil
}
