package main

import (
	"github.com/aerosystems/mail-service/internal/config"
	HttpServer "github.com/aerosystems/mail-service/internal/http"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	httpServer *HttpServer.Server
	rpcServer  *RpcServer.Server
}
