package rest

import (
	"errors"
	"github.com/aerosystems/mail-service/internal/models"
	"github.com/aerosystems/mail-service/internal/validators"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/rpc"
	"os"
)

type FeedbackHandler struct {
	*BaseHandler
	mailService MailService
}

func NewFeedbackHandler(baseHandler *BaseHandler, mailService MailService) *FeedbackHandler {
	return &FeedbackHandler{
		BaseHandler: baseHandler,
		mailService: mailService,
	}

}

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
// @Failure 400 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/feedback [post]
func (fh FeedbackHandler) SendFeedback(c echo.Context) error {
	var feedbackRequest FeedbackRequest
	if err := c.Bind(&feedbackRequest); err != nil {
		return fh.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}

	if feedbackRequest.Name == "" {
		return fh.ErrorResponse(c, http.StatusBadRequest, "name is required", nil)
	}

	if feedbackRequest.Email == "" {
		return fh.ErrorResponse(c, http.StatusBadRequest, "email is required", nil)
	}

	if feedbackRequest.Message == "" {
		return fh.ErrorResponse(c, http.StatusBadRequest, "message is required", nil)
	}

	// checking email in blacklist via RPC
	if checkmailClientRPC, err := rpc.Dial("tcp", "checkmail-service:5001"); err == nil {
		var result string
		if err := checkmailClientRPC.Call(
			"Server.Inspect",
			InspectRPCPayload{
				Domain:   feedbackRequest.Email,
				ClientIp: c.RealIP(),
			},
			&result); err != nil {
			return fh.ErrorResponse(c, http.StatusBadRequest, "email address does not valid", err)
		}

		if result == "blacklist" {
			err := errors.New("email address is blacklisted")
			return fh.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		}
	} else {
		fh.log.Error(err)
	}

	feedbackEmail, err := validators.ValidateEmail(os.Getenv("FEEDBACK_EMAIL"))
	if err != nil {
		panic(err)
	}

	msg := models.Message{
		To:       feedbackEmail,
		ToName:   "Verifire💎",
		FromName: feedbackRequest.Name,
		From:     "no-reply@verifire.dev",
		Cc:       feedbackRequest.Email,
		Subject:  "Feedback from " + feedbackRequest.Email,
		Body:     feedbackRequest.Message,
	}

	if err := fh.mailService.Send(msg); err != nil {
		return fh.ErrorResponse(c, http.StatusInternalServerError, "error sending feedback", err)
	}

	return fh.SuccessResponse(c, http.StatusOK, "feedback sent successfully", nil)
}
