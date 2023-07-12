package RPCServer

import (
	"github.com/aerosystems/mail-service/internal/provider"
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"log"
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
	mailService.SetProvider(&provider.Brevo{})
	if err := mailService.SendEmail(payload.To, payload.Subject, payload.Body); err != nil {
		return err
	}

	*resp = "Sent email to " + payload.To
	return nil
}
