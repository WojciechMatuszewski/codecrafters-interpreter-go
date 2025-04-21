package lox_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		input       string
		expectedOut string
		expectedErr string
	}{
		{
			input:       "10",
			expectedOut: "10",
			expectedErr: "",
		},
		{
			input:       "\"hello world!\"",
			expectedOut: "hello world!",
			expectedErr: "",
		},
		{
			input:       "10.40",
			expectedOut: "10.4",
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := bytes.NewReader([]byte(tt.input))

			l := lox.NewLox()
			out := fmt.Sprintf("%v", l.Evaluate(r))

			if out != tt.expectedOut {
				t.Errorf("\nexpected output:\n%q\ngot:\n%q\n", tt.expectedOut, out)
			}
		})
	}

}
