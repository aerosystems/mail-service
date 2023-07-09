package main

import (
	"log"
)

// RPCServer is the type for our RPC Server. Methods that take this as a receiver are available
// over RPC, as long as they are exported.
type RPCServer struct{}

type RPCPayload struct {
	To      string
	Subject string
	Body    string
}

func (r *RPCServer) SendEmail(payload RPCPayload, resp *string) error {
	log.Println("sending email to", payload.To)
	*resp = "Sent email to " + payload.To
	return nil
}
