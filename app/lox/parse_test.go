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
		expectedErr string
	}{
		{
			input:       "(\"foo\")",
			expectedOut: "(group foo)",
			expectedErr: "",
		},
		{
			input:       "(12.999)",
			expectedOut: "(group 12.999)",
			expectedErr: "",
		},
		{
			input:       "(nil)",
			expectedOut: "(group nil)",
			expectedErr: "",
		},
		{
			input:       "(12.87)",
			expectedOut: "(group 12.87)",
			expectedErr: "",
		},
		{
			input:       "!true",
			expectedOut: "(! true)",
			expectedErr: "",
		},
		{
			input:       "16 * 38 / 58",
			expectedOut: "(/ (* 16.0 38.0) 58.0)",
			expectedErr: "",
		},
		{
			input:       "\"baz\" == \"baz\"",
			expectedOut: "(== baz baz)",
			expectedErr: "",
		},
		{
			input:       "(72+)",
			expectedOut: "",
			expectedErr: "[line 1] Error at ')': Expect expression.",
		},
		{
			input:       "\"baz\"!=\"world\"",
			expectedOut: "(!= baz world)",
			expectedErr: "",
		},
		{
			input:       "!true",
			expectedOut: "(! true)",
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := bytes.NewReader([]byte(tt.input))

			l := lox.NewLox()
			statements, err := l.Parse(r)
			if err != nil {
				if tt.expectedErr == "" {
					t.Fatalf("did not expect error, but got: %v", err)
				}

				if err.Error() != tt.expectedErr {
					t.Errorf("\nexpected error:\n%q\ngot:\n%q\n", tt.expectedOut, err.Error())
				}
			}

			if statements != nil {
				out := lox.Format(statements)

				if tt.expectedOut == "" {
					t.Fatalf("did not expect output, but got: %v", out)
				}

				if out != tt.expectedOut {
					t.Errorf("\nexpected output:\n%q\ngot:\n%q\n", tt.expectedOut, out)
				}
			}

			if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error:\n%q\nreceived: none\n", tt.expectedErr)
			}

			if statements == nil && tt.expectedOut != "" {
				t.Errorf("expected output:\n%q\nreceived: none\n", tt.expectedOut)
			}

		})
	}

}
