package WebServer

import (
	"github.com/labstack/echo/v4"
)

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	e.POST("/v1/feedback", app.SendFeedback)

	return e
}
