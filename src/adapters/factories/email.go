package factories

import (
	"github.com/rafixcs/uberemailservicechallenge/src/adapters/controllers"
	email_service "github.com/rafixcs/uberemailservicechallenge/src/services/email"
)

func NewEmailController() *controllers.EmailController {
	controller := &controllers.EmailController{
		EmailService: email_service.AwsSimpleEmailService{},
	}
	return controller
}
