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
		{
			input:       "(\"hello world\")",
			expectedOut: "hello world",
			expectedErr: "",
		},
		{
			input:       "((true))",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "-73",
			expectedOut: "-73",
			expectedErr: "",
		},
		{
			input:       "!true",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "!10.40",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "!nil",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "nil",
			expectedOut: "<nil>",
			expectedErr: "",
		},
		{
			input:       "\"42\"+\"42\"",
			expectedOut: "4242",
			expectedErr: "",
		},
		{
			input:       "42 + 42",
			expectedOut: "84",
			expectedErr: "",
		},
		{
			input:       "57 > -65",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "(54 - 67) >= -(114 / 57 + 11)",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "(54 - 67) < -(114 / 57 + 11)",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "\"hello\" == \"hello\"",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "\"hello1\" == \"hello\"",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "\"hello\" != \"hello\"",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "\"hello1\" != \"hello\"",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "42 == 42",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "43 == 42",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "42 != 42",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "43 != 42",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "42 == \"42\"",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "nil == nil",
			expectedOut: "true",
			expectedErr: "",
		},
		{
			input:       "nil == 42",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "42 == nil",
			expectedOut: "false",
			expectedErr: "",
		},
		{
			input:       "-\"foo\"",
			expectedOut: "",
			expectedErr: "Operand must be a number.\n[line 1]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := bytes.NewReader([]byte(tt.input))

			l := lox.NewLox()
			result, err := l.Evaluate(r)

			if err != nil && tt.expectedErr == "" {
				t.Fatalf("did not expect error, but got: %v", err)
			}

			if err != nil && tt.expectedErr != err.Error() {
				t.Errorf("\nexpected error:\n%q\ngot:\n%q\n", tt.expectedErr, err.Error())
			}

			out := fmt.Sprint(result)
			if result == nil && tt.expectedOut != "" && tt.expectedOut != "<nil>" {
				t.Fatalf("did not expect output, but got: %v", out)
			}

			if result != nil && out != tt.expectedOut {
				t.Errorf("\nexpected output:\n%q\ngot:\n%q\n", tt.expectedOut, out)
			}
		})
	}

}

func TestRun(t *testing.T) {
	tests := []struct {
		input       string
		expectedOut string
		expectedErr string
	}{
		{
			input:       "print false != false;",
			expectedOut: "false\n",
			expectedErr: "",
		},
		{
			input:       "print \"36\n10\n78\n\";print\"foo\";",
			expectedOut: "36\n10\n78\n\nfoo\n",
			expectedErr: "",
		},
		{
			input:       "27 - 60 >= -99 * 2 / 99 + 76;\ntrue == true;\n(\"world\" == \"bar\") == (\"baz\" != \"hello\");\nprint true;",
			expectedOut: "true\n",
			expectedErr: "",
		},
		{
			input:       "print \"the expression below is invalid\";\n49 + \"baz\";\nprint \"this should not be printed\";\n",
			expectedOut: "the expression below is invalid\n",
			expectedErr: "Operands must be two numbers or two strings.\n[line 2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := bytes.NewReader([]byte(tt.input))
			output := bytes.NewBuffer([]byte(""))

			l := lox.NewLox()
			err := l.Run(input, output)

			if err != nil && tt.expectedErr == "" {
				t.Fatalf("did not expect error, but got: %v", err)
			}

			if err != nil && tt.expectedErr != err.Error() {
				t.Errorf("\nexpected error:\n%q\ngot:\n%q\n", tt.expectedErr, err.Error())
			}

			result := output.String()
			out := fmt.Sprint(result)
			if result == "" && tt.expectedOut != "" && tt.expectedOut != "<nil>" {
				t.Fatalf("did not expect output, but got: %v", out)
			}

			if result != "" && out != tt.expectedOut {
				t.Errorf("\nexpected output:\n%q\ngot:\n%q\n", tt.expectedOut, out)
			}
		})
	}

}
