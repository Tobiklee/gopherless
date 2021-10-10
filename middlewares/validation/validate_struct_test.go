package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestObject struct {
	FieldOne string `validate:"required,min=3,max=32"`
	FieldTwo string `validate:"required,min=3,max=32"`
}

func TestValidateStruct(t *testing.T) {
	testPerson := TestObject{FieldOne: "Field", FieldTwo: "Mu"}

	errors := ValidateStruct(testPerson)

	expected := ErrorResponse{
		FailedField: "TestObject.FieldTwo",
		Tag:         "min",
		Value:       "3",
	}

	if assert.NotNil(t, errors) {
		assert.Equal(t, 1, len(errors), "length should be 1")
		assert.Equal(t, expected, *errors[0], "error response should match")
	}
}
