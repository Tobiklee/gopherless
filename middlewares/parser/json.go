package parser

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// Parse changes a json to a struct of type Request
func Parse(input string, request interface{}) *events.APIGatewayV2HTTPResponse {
	err := json.Unmarshal([]byte(input), &request)

	if err != nil {
		log.Println("Unmarshal on body failed")
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
		}
	}

	return nil
}

// Stringify changes from struct to json
func Stringify(input interface{}) (string, error) {
	output, err := json.Marshal(input)
	return string(output), err
}
