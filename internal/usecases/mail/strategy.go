package mail

import (
	"github.com/aerosystems/mail-service/internal/models"
)

type EmailProvider interface {
	SendEmail(msg models.Message) error
}

type EmailService struct {
	Provider EmailProvider
}

func NewEmailService(provider EmailProvider) *EmailService {
	var emailService EmailService
	emailService.SetProvider(provider)
	return &emailService
}

func (e *EmailService) SetProvider(provider EmailProvider) {
	e.Provider = provider
}

func (e *EmailService) Send(msg models.Message) error {
	return e.Provider.SendEmail(msg)
}
