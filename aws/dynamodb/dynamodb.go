package dynamodb

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/config"
)

type Key struct {
	PK string
	SK string
}

type DynamoService struct {
	Client *dynamodb.Client
	Table  string
}

// New creates a new service instance
func New(region, table string) *DynamoService {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = region
		return nil
	})
	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &DynamoService{
		Client: client,
		Table:  table,
	}
}

// Put inserts incoming event into event store.
func (store DynamoService) Put(object interface{}) error {
	marshalMap, err := attributevalue.MarshalMap(object)

	if err != nil {
		return err
	}

	_, err = store.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(store.Table),
		Item:      marshalMap,
	})

	return err
}

// SimpleUpdate updates an item with give values. It only uses SET expressions.
func (store DynamoService) SimpleUpdate(item interface{}) error {
	marshalMap, err := attributevalue.MarshalMap(item)

	if err != nil {
		return err
	}

	updateExpressions := []string{}

	expressionValues := make(map[string]types.AttributeValue)

	marshalKey := make(map[string]types.AttributeValue)

	for k, v := range marshalMap {
		if k == "PK" || k == "SK" {
			marshalKey[k] = v
			continue
		}
		expression := "set " + k + " = :" + k
		updateExpressions = append(updateExpressions, expression)
		expressionValues[":"+k] = v
	}

	_, err = store.Client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(store.Table),
		Key:                       marshalKey,
		UpdateExpression:          aws.String(strings.Join(updateExpressions, ",")),
		ExpressionAttributeValues: expressionValues,
	})

	return err
}
