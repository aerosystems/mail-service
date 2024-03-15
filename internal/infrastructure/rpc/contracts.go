package RpcServer

import "github.com/aerosystems/mail-service/internal/models"

type MailService interface {
	Send(msg models.Message) error
}
