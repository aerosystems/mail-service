package RpcServer

import (
	"fmt"
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"github.com/sirupsen/logrus"
	"net"
	"net/rpc"
)

type MailServer struct {
	rpcPort     int
	log         *logrus.Logger
	mailService *MailService.MailService
}

func New(
	rpcPort int,
	log *logrus.Logger,
	mailService *MailService.MailService,
) *MailServer {
	return &MailServer{
		rpcPort:     rpcPort,
		log:         log,
		mailService: mailService,
	}
}

func (ms *MailServer) Listen(rpcPort int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
