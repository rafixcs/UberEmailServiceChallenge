package email_service

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type AwsSimpleEmailService struct{}

func (AwsSimpleEmailService) SendEmail() error {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		err := fmt.Errorf("[LoadDefaultConfig]: %w", err)
		return err
	}

	client := ses.NewFromConfig(cfg)

	input := &ses.SendEmailInput{
		Source: aws.String("rafael.camargo.rs+dev01@gmail.com"),
		Destination: &types.Destination{
			ToAddresses: []string{"rafaelcs.dev.br@gmail.com"},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data:    aws.String("Hello from Go + AWS SES"),
				Charset: aws.String("UTF-8"),
			},
			Body: &types.Body{
				Text: &types.Content{
					Data:    aws.String("This is a test email sent with aws ses and go!"),
					Charset: aws.String("UTF-8"),
				},
			},
		},
	}
	res, err := client.SendEmail(ctx, input)
	if err != nil {
		err := fmt.Errorf("failed to send email: %w", err)
		return err
	}

	log.Printf("Email sent! Message ID: %s\n", *res.MessageId)
	return nil
}
