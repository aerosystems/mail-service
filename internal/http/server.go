package HttpServer

import (
	"fmt"
	"github.com/aerosystems/mail-service/internal/infrastructure/rest"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log             *logrus.Logger
	echo            *echo.Echo
	feedbackHandler *rest.FeedbackHandler
}

func NewServer(
	log *logrus.Logger,
	feedbackHandler *rest.FeedbackHandler,
) *Server {
	return &Server{
		log:             log,
		echo:            echo.New(),
		feedbackHandler: feedbackHandler,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	s.log.Infof("starting HTTP server mail-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}
