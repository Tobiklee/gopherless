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

type IService interface {
	Put(object interface{}) error
	SimpleUpdate(item interface{}) error
	SimpleGet(primaryKey, sortKey string) (*dynamodb.GetItemOutput, error)
	SimpleDelete(primaryKey, sortKey string) (*dynamodb.DeleteItemOutput, error)
	GetByPK(pk string) (*dynamodb.QueryOutput, error)
}

type Service struct {
	Client *dynamodb.Client
	Table  string
}

// New creates a new service instance
func New(region, table string, cfg *aws.Config) *Service {
	var c aws.Config
	if cfg == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
			opts.Region = region
			return nil
		})
		c = cfg
		if err != nil {
			panic(err)
		}
	} else {
		c = *cfg
	}
	client := dynamodb.NewFromConfig(c)

	return &Service{
		Client: client,
		Table:  table,
	}
}

// Put inserts incoming event into event store.
func (store Service) Put(object interface{}) error {
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
func (store Service) SimpleUpdate(item interface{}) error {
	marshalMap, err := attributevalue.MarshalMap(item)

	if err != nil {
		return err
	}

	updateExpressions := []string{}
	expressionValues := make(map[string]types.AttributeValue)
	expressionAttributeNames := make(map[string]string)
	marshalKey := make(map[string]types.AttributeValue)

	for k, v := range marshalMap {
		if k == "PK" || k == "SK" {
			marshalKey[k] = v
			continue
		}
		if k == "CreatedAt" {
			var pk, sk string
			err := attributevalue.Unmarshal(marshalMap["PK"], &pk)
			if err != nil {
				return err
			}
			err = attributevalue.Unmarshal(marshalMap["SK"], &sk)
			if err != nil {
				return err
			}
			get, err := store.SimpleGet(pk, sk)
			if err != nil {
				return err
			}
			// skip adding CreatedAt value if an item with the given PK and SK already exists
			if get != nil && get.Item != nil {
				continue
			}
		}

		expressionAttributeNames["#"+k] = k // used to avoid running in errors with reserved dynamodb-keywords
		expression := "#" + k + " = :" + k
		updateExpressions = append(updateExpressions, expression)
		expressionValues[":"+k] = v
	}

	_, err = store.Client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(store.Table),
		Key:                       marshalKey,
		UpdateExpression:          aws.String("set " + strings.Join(updateExpressions, ",")),
		ExpressionAttributeValues: expressionValues,
		ExpressionAttributeNames:  expressionAttributeNames,
	})

	return err
}

func (store Service) SimpleGet(primaryKey, sortKey string) (*dynamodb.GetItemOutput, error) {
	keyMap, err := attributevalue.MarshalMap(map[string]string{
		"PK": primaryKey,
		"SK": sortKey,
	})
	if err != nil {
		return nil, err
	}

	output, err := store.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(store.Table),
		Key:       keyMap,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (store Service) SimpleDelete(primaryKey, sortKey string) (*dynamodb.DeleteItemOutput, error) {
	keyMap, err := attributevalue.MarshalMap(map[string]string{
		"PK": primaryKey,
		"SK": sortKey,
	})
	if err != nil {
		return nil, err
	}

	output, err := store.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(store.Table),
		Key:       keyMap,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (store Service) GetByPK(pk string) (*dynamodb.QueryOutput, error) {
	expressionValues := make(map[string]types.AttributeValue)
	expressionValues[":PK"], _ = attributevalue.Marshal(pk)

	output, err := store.Client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(store.Table),
		KeyConditionExpression:    aws.String("PK = :PK"),
		ExpressionAttributeValues: expressionValues,
	})
	return output, err
}
