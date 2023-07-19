package main

import (
	RPCServer "github.com/aerosystems/mail-service/internal/rpc_server"
	"log"
	"net/rpc"
)

func main() {
	if err := rpc.Register(new(RPCServer.MailServer)); err != nil {
		log.Fatal(err)
	}
	if err := RPCServer.Listen(); err != nil {
		log.Fatal(err)
	}
}
