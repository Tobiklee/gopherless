package validation

import (
	"reflect"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	type args struct {
		input interface{}
	}
	var tests []struct {
		name string
		args args
		want []*ErrorResponse
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateStruct(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
