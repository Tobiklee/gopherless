package dynamodb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type EventStoreService struct {
	Client *dynamodb.Client
	Table  string
}

type EventModel struct {
	PK         string
	SK         string
	Type       string
	ExternalID int32
	Request    interface{}
}

type PutParams struct {
	Name       string
	EventType  string
	ExternalID int32
	Request    interface{}
}

// Put inserts incoming event into event store.
//
// Format:
// PK: EVENT#CreateUser
// SK: <ISO-DATE>
// Type: HttpGatewayProxyEvent
// ExternalID: 500111
// Request: <Request Object>
func (store EventStoreService) Put(event PutParams) error {
	eventModel := EventModel{
		PK:         "EVENT#" + event.Name,
		SK:         time.Now().Format(time.RFC3339),
		Type:       event.EventType,
		ExternalID: event.ExternalID,
		Request:    event.Request,
	}

	marshalMap, err := attributevalue.MarshalMap(eventModel)

	if err != nil {
		return err
	}

	_, err = store.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(store.Table),
		Item:      marshalMap,
	})

	return err
}
