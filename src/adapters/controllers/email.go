package controllers

import (
	"fmt"
	"net/http"

	email_service "github.com/rafixcs/uberemailservicechallenge/src/services/email"
)

type EmailController struct {
	EmailService email_service.IEmailService
}

func (e *EmailController) Handle(w http.ResponseWriter, r *http.Request) {
	err := e.EmailService.SendEmail()
	if err != nil {
		err := fmt.Errorf("[EmailController.Handle]: failed to send email: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
