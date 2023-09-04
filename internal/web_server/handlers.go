package WebServer

import (
	"github.com/aerosystems/auth-service/pkg/validators"
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/rpc"
	"os"
)

type FeedbackRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type RPCInspectPayload struct {
	Domain   string
	ClientIp string
}

func (app *Config) SendFeedback(c echo.Context) error {
	var feedbackRequest FeedbackRequest
	if err := c.Bind(&feedbackRequest); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid request body")
	}

	if feedbackRequest.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required")
	}

	if feedbackRequest.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email is required")
	}

	if feedbackRequest.Message == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "message is required")
	}

	// checking email in blacklist via RPC
	if checkmailClientRPC, err := rpc.Dial("tcp", "checkmail-service:5001"); err == nil {
		var result string
		if err := checkmailClientRPC.Call(
			"CheckmailServer.Inspect",
			RPCInspectPayload{
				Domain:   feedbackRequest.Email,
				ClientIp: c.RealIP(),
			},
			&result); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "email address does not valid")
		}

		if result == "blacklist" {
			return echo.NewHTTPError(http.StatusBadRequest, "email address contains in blacklist")
		}
	} else {
		app.log.Error(err)
	}

	feedbackEmail, err := validators.ValidateEmail(os.Getenv("FEEDBACK_EMAIL"))
	if err != nil {
		panic(err)
	}

	msg := MailService.Message{
		To:       feedbackEmail,
		ToName:   "Testmail💎",
		FromName: feedbackRequest.Name,
		From:     "no-reply@testmail.top",
		Cc:       feedbackRequest.Email,
		Subject:  "Feedback from " + feedbackRequest.Email,
		Body:     feedbackRequest.Message,
	}

	if err := app.mailService.SendEmail(msg); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error sending email")
	}

	return c.JSON(http.StatusOK, "feedback sent successfully")
}
