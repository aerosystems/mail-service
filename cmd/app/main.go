package main

import (
	RPCServer "github.com/aerosystems/mail-service/internal/rpc_server"
	"github.com/aerosystems/mail-service/pkg/logger"
	"net/rpc"
	"os"
)

func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	if err := rpc.Register(new(RPCServer.MailServer)); err != nil {
		log.Fatal(err)
	}
	if err := RPCServer.Listen(); err != nil {
		log.Fatal(err)
	}
}
