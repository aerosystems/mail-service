package RPCServer

import (
	"fmt"
	"github.com/aerosystems/mail-service/pkg/logger"
	"net"
	"net/rpc"
	"os"
)

const rpcPort = 5001

func Listen() error {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))
	log.Info("starting RPC server on 0.0.0.0:%d", rpcPort)
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
