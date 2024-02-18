package RpcServer

import (
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"os"
)

type MailRPCPayload struct {
	To      string
	Subject string
	Body    string
}

func (ms *MailServer) SendEmail(payload MailRPCPayload, resp *string) error {
	ms.log.Infof("sending email to %s", payload.To)

	msg := MailService.Message{
		To:       payload.To,
		ToName:   "Customer",
		FromName: "Verifire💎",
		From:     "no-reply@verifire.dev",
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
