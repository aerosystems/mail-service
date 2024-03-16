//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aerosystems/mail-service/internal/config"
	HttpServer "github.com/aerosystems/mail-service/internal/infrastructure/http"
	"github.com/aerosystems/mail-service/internal/infrastructure/http/handlers"
	RpcServer "github.com/aerosystems/mail-service/internal/infrastructure/rpc"
	"github.com/aerosystems/mail-service/internal/usecases/mail"
	"github.com/aerosystems/mail-service/internal/usecases/mail/provider"
	"github.com/aerosystems/mail-service/pkg/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(handlers.MailService), new(*mail.EmailService)),
		wire.Bind(new(RpcServer.MailService), new(*mail.EmailService)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHttpServer,
		ProvideRpcServer,
		ProvideLogrusLogger,
		ProvideBaseHandler,
		ProvideFeedbackHandler,
		ProvideMailhogProvider,
		ProvideBrevoProvider,
		ProvideMailService,
	))
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server, rpcServer *RpcServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHttpServer(log *logrus.Logger, cfg *config.Config, feedbackHandler *handlers.FeedbackHandler) *HttpServer.Server {
	panic(wire.Build(HttpServer.NewServer))
}

func ProvideRpcServer(log *logrus.Logger, mailService RpcServer.MailService) *RpcServer.Server {
	panic(wire.Build(RpcServer.NewServer))
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *handlers.BaseHandler {
	return handlers.NewBaseHandler(log, cfg.Mode)
}

func ProvideFeedbackHandler(baseHandler *handlers.BaseHandler, mailService handlers.MailService) *handlers.FeedbackHandler {
	return handlers.NewFeedbackHandler(baseHandler, mailService)
}

func ProvideMailhogProvider(cfg *config.Config) *provider.Smtp {
	return provider.NewSmtp(
		cfg.MailhogDomain,
		cfg.MailhogHost,
		cfg.MailhogPort,
		cfg.MailhogUsername,
		cfg.MailhogPassword,
		cfg.MailhogEncryption,
	)
}

func ProvideBrevoProvider(cfg *config.Config) *provider.Brevo {
	return provider.NewBrevo(cfg.BrevoApiKey)
}

func ProvideMailService(cfg *config.Config, brevo *provider.Brevo, mailhog *provider.Smtp) *mail.EmailService {
	provider, err := mail.FromString(cfg.EmailProvider)
	if err != nil {
		panic(err)
	}
	switch provider {
	case mail.Mailhog:
		return mail.NewEmailService(mailhog)
	case mail.Brevo:
		return mail.NewEmailService(brevo)
	}
	panic("no email provider set")
}
