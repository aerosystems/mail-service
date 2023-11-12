package WebServer

import (
	"errors"
	MailService "github.com/aerosystems/mail-service/pkg/mail_service"
	"github.com/aerosystems/mail-service/pkg/validators"
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

type InspectRPCPayload struct {
	Domain   string
	ClientIp string
}

// SendFeedback godoc
// @Summary Send feedback
// @Description Send feedback
// @Tags feedback
// @Accept json
// @Produce json
// @Param feedbackRequest body FeedbackRequest true "feedback request"
// @Success 200 {object} Response
// @Failure 400 {object} ErrResponse
// @Failure 422 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /v1/feedback [post]
func (app *Config) SendFeedback(c echo.Context) error {
	var feedbackRequest FeedbackRequest
	if err := c.Bind(&feedbackRequest); err != nil {
		return ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}

	if feedbackRequest.Name == "" {
		return ErrorResponse(c, http.StatusBadRequest, "name is required", nil)
	}

	if feedbackRequest.Email == "" {
		return ErrorResponse(c, http.StatusBadRequest, "email is required", nil)
	}

	if feedbackRequest.Message == "" {
		return ErrorResponse(c, http.StatusBadRequest, "message is required", nil)
	}

	// checking email in blacklist via RPC
	if checkmailClientRPC, err := rpc.Dial("tcp", "checkmail-service:5001"); err == nil {
		var result string
		if err := checkmailClientRPC.Call(
			"CheckmailServer.Inspect",
			InspectRPCPayload{
				Domain:   feedbackRequest.Email,
				ClientIp: c.RealIP(),
			},
			&result); err != nil {
			return ErrorResponse(c, http.StatusBadRequest, "email address does not valid", err)
		}

		if result == "blacklist" {
			err := errors.New("email address is blacklisted")
			return ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
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
		return ErrorResponse(c, http.StatusInternalServerError, "error sending feedback", err)
	}

	return SuccessResponse(c, http.StatusOK, "feedback sent successfully", nil)
}
