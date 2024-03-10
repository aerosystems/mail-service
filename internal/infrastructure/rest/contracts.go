package rest

import "github.com/aerosystems/mail-service/internal/models"

type MailService interface {
	SendEmail(msg models.Message) error
}
