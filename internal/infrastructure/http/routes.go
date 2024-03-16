package HttpServer

func (s *Server) setupRoutes() {
	s.echo.POST("/v1/feedback", s.feedbackHandler.SendFeedback)
}
