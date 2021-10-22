package dynamodb

import (
	"time"
)

type EventStoreService struct {
	DynamoService *DynamoService
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

	err := store.DynamoService.Put(eventModel)

	return err
}
