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

func (AwsSimpleEmailService) SendEmail(emailContent SendEmailContent) error {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		err := fmt.Errorf("[LoadDefaultConfig]: %w", err)
		return err
	}

	client := ses.NewFromConfig(cfg)

	input := &ses.SendEmailInput{
		Source: aws.String(emailContent.Source),
		Destination: &types.Destination{
			ToAddresses: []string{emailContent.Destination},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data:    aws.String(emailContent.Subject),
				Charset: aws.String("UTF-8"),
			},
			Body: &types.Body{
				Text: &types.Content{
					Data:    aws.String(emailContent.Content),
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
