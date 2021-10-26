package parser

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// parse changes a json to a struct of type Request
func parse(input string, request interface{}) *events.APIGatewayV2HTTPResponse {
	err := json.Unmarshal([]byte(input), &request)

	if err != nil {
		log.Println("Unmarshal on body failed")
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
		}
	}

	return nil
}

// stringify changes from struct to json
func stringify(input interface{}) (string, error) {
	output, err := json.Marshal(input)
	return string(output), err
}
