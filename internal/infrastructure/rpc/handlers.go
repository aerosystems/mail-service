package RpcServer

import "github.com/aerosystems/mail-service/internal/models"

type MailRPCPayload struct {
	To      string
	Subject string
	Body    string
}

func (s Server) SendEmail(payload MailRPCPayload, resp *string) error {
	s.log.Infof("sending email to %s", payload.To)

	msg := models.Message{
		To:       payload.To,
		ToName:   "Customer",
		FromName: "Verifire💎",
		From:     "no-reply@verifire.dev",
		Subject:  payload.Subject,
		Body:     payload.Body,
	}

	if err := s.mailService.Send(msg); err != nil {
		s.log.Errorf("error sending email: %v", err)
		return err
	}

	*resp = "Sent email to " + payload.To
	return nil
}
