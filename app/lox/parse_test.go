package lox_test

import (
	"bytes"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input       string
		expectedOut string
	}{
		{
			input:       "(\"foo\")",
			expectedOut: "(group foo)",
		},
		{
			input:       "!true",
			expectedOut: "(! true)",
		},
		{
			input:       "16 * 38 / 58",
			expectedOut: "(/ (* 16.0 38.0) 58.0)",
		},
	}

	visitor := lox.PrinterVisitor{}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := bytes.NewReader([]byte(tt.input))

			l := lox.NewLox()
			expr, err := l.Parse(r)
			if err != nil {
				t.Fatalf("failed to parse\ninput: %v\nerror: %v", tt.input, err)
			}

			out := expr.Accept(&visitor)
			if out != tt.expectedOut {
				t.Errorf("\nexpected output:\n%q\ngot:\n%q\n", tt.expectedOut, out)
			}

		})
	}

}
