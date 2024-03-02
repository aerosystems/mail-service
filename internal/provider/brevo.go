package provider

import (
	"bytes"
	"encoding/json"
	MailService "github.com/aerosystems/mail-service/internal/usecases/mail"
	"net/http"
	"os"
)

type Brevo struct{}

type BrevoRequestPayload struct {
	Sender      BrevoMailPerson   `json:"sender"`
	To          []BrevoMailPerson `json:"to"`
	Subject     string            `json:"subject"`
	HTMLContent string            `json:"htmlContent"`
}

type BrevoMailPerson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (b *Brevo) SendEmail(msg MailService.Message) error {
	requestPayload := &BrevoRequestPayload{
		Sender: BrevoMailPerson{
			Name:  msg.FromName,
			Email: msg.From,
		},
		To: []BrevoMailPerson{
			{
				Name:  msg.ToName,
				Email: msg.To,
			},
		},
		Subject:     msg.Subject,
		HTMLContent: msg.Body,
	}

	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", os.Getenv("BREVO_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return err
	}

	// TODO: add logging

	return nil
}
