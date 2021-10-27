package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Service struct {
	Client *sqs.Client
	URL    string
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
