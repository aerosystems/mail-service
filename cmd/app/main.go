package main

import (
	"fmt"
	"github.com/aerosystems/mail-service/internal/provider"
	RPCServer "github.com/aerosystems/mail-service/internal/rpc_server"
	WebServer "github.com/aerosystems/mail-service/internal/web_server"
	"github.com/aerosystems/mail-service/pkg/logger"
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/rpc"
	"os"
	"strconv"
)

const webPort = 80
const rpcPort = 5001

// @title Mail Service
// @version 1.0.1
// @description A part of microservice infrastructure, who responsible for sending emails

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host gw.verifire.com/mail
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	mailService := &MailService.MailService{}
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
		log.Fatal("no email provider set")
	}

	mailServer := RPCServer.New(rpcPort, log.Logger, mailService)
	webServer := WebServer.New(log.Logger, mailService)

	e := webServer.NewRouter()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	e.Use(middleware.Recover())

	errChan := make(chan error)

	go func() {
		log.Infof("starting mail-service RPC server on port %d\n", rpcPort)
		errChan <- rpc.Register(mailServer)
		errChan <- mailServer.Listen(rpcPort)
	}()

	go func() {
		log.Infof("starting mail-service HTTP server on port %d\n", webPort)
		errChan <- e.Start(fmt.Sprintf(":%d", webPort))
	}()

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
