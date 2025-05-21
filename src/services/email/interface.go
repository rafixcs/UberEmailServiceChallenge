package email_service

type SendEmailContent struct {
	Source      string
	Destination string
	Content     string
	Subject     string
}
type IEmailService interface {
	SendEmail(emailContent SendEmailContent) error
}
