package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	email_service "github.com/rafixcs/uberemailservicechallenge/src/services/email"
)

type EmailController struct {
	EmailService email_service.AwsSimpleEmailService
}

type SendEmailRequest struct {
	Source      string
	Destination string
	Content     string
	Subject     string
}

func (e *EmailController) Handle(w http.ResponseWriter, r *http.Request) {
	var reqBody SendEmailRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		err := fmt.Errorf("[EmailController.Handle]: failed to send email: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	emailContent := email_service.SendEmailContent{
		Source:      reqBody.Source,
		Destination: reqBody.Destination,
		Content:     reqBody.Content,
		Subject:     reqBody.Subject,
	}

	err = e.EmailService.SendEmail(emailContent)
	if err != nil {
		err := fmt.Errorf("[EmailController.Handle]: failed to send email: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
