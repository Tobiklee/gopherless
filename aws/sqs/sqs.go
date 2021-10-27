package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Service struct {
	Client *sqs.Client
	URL    string
}

// New creates a new service instance
func New(region, url string) *Service {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = region
		return nil
	})
	if err != nil {
		panic(err)
	}

	client := sqs.NewFromConfig(cfg)

	return &Service{
		Client: client,
		URL:    url,
	}
}

// SendMessage sends a message to the initialized URL
func (s Service) SendMessage(body string) error {
	input := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageBody:  aws.String(body),
		QueueUrl:     aws.String(s.URL),
	}
	if _, err := s.Client.SendMessage(context.TODO(), input); err != nil {
		return err
	}
	return nil
}
