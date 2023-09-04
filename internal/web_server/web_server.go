package WebServer

import (
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"github.com/sirupsen/logrus"
)

type Config struct {
	log         *logrus.Logger
	mailService *MailService.MailService
}

func New(
	log *logrus.Logger,
	mailService *MailService.MailService,
) *Config {
	return &Config{
		log:         log,
		mailService: mailService,
	}
}
