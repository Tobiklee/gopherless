package parser

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

const (
	TestJSON = "{\"string1\":\"string1\",\"number1\":123456788,\"number2\":123456789,\"bool1\":true," +
		"\"bool2\":false,\"string2\":\"string2\",\"number3\":1000}"
	TestJSON2 = "{\"string1\":\"string1\",\"number1\":123456788,\"number2\":123456789,\"bool1\":true," +
		"\"bool2\":false}"
)

type TestJSONStruct struct {
	String1 string `json:"string1"`
	Number  int    `json:"number1"`
	Number2 int    `json:"number2"`
	Bool1   bool   `json:"bool1"`
	Bool2   bool   `json:"bool2"`
	String2 string `json:"string2"`
	Number3 int    `json:"number3"`
}

func TestJsonParse(t *testing.T) {
	var result TestJSONStruct

	response := parse(TestJSON, &result)

	expected := TestJSONStruct{
		String1: "string1",
		Number:  123456788,
		Number2: 123456789,
		Bool1:   true,
		Bool2:   false,
		String2: "string2",
		Number3: 1000,
	}

	if assert.NotNil(t, result) {
		assert.Equal(t, result, expected, "parsed json should match given struct")
		assert.Equal(t, response, (*events.APIGatewayV2HTTPResponse)(nil), "error response should be nil")
	}
}

func TestJsonWithIncompleteData(t *testing.T) {
	var result TestJSONStruct

	response := parse(TestJSON2, &result)

	expected := TestJSONStruct{
		String1: "string1",
		Number:  123456788,
		Number2: 123456789,
		Bool1:   true,
		Bool2:   false,
		String2: "",
		Number3: 0,
	}

	if assert.NotNil(t, result) {
		assert.Equal(t, result, expected, "parsed json should match given (empty) struct")
		assert.Equal(t, response, (*events.APIGatewayV2HTTPResponse)(nil), "error response should be nil")
	}
}

func TestStringify(t *testing.T) {
	input := TestJSONStruct{
		String1: "string1",
		Number:  123456788,
		Number2: 123456789,
		Bool1:   true,
		Bool2:   false,
		String2: "string2",
		Number3: 1000,
	}

	jsonString, err := stringify(input)

	if assert.NotNil(t, jsonString) {
		assert.Equal(t, jsonString, TestJSON, "json string should match")
		assert.Equal(t, err, nil, "err should be nil")
	}
}
