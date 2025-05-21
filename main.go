package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func SendEmail() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		err := fmt.Errorf("[LoadDefaultConfig]: %w", err)
		log.Fatal(err)
	}

	client := ses.NewFromConfig(cfg)

	input := &ses.SendEmailInput{
		Source: aws.String("rafaelcs.dev.br@gmail.com"),
		Destination: &types.Destination{
			ToAddresses: []string{"rafael.camargo.rs+dev01@gmail.com"},
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
		log.Fatalf("failed to send email: %w", err.Error())
	}

	log.Printf("Email sent! Message ID: %s\n", *res.MessageId)
}

func main() {
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello\n"))
	//})
	//log.Fatal(http.ListenAndServe(":3000", r))
	SendEmail()
}
