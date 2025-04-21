package lox_test

import (
	"strings"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input       string
		expectedOut string
		expectedErr string
		only        bool
	}{
		{
			input:       "({*.,+*})",
			expectedOut: "LEFT_PAREN ( null\nLEFT_BRACE { null\nSTAR * null\nDOT . null\nCOMMA , null\nPLUS + null\nSTAR * null\nRIGHT_BRACE } null\nRIGHT_PAREN ) null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "4043.6490",
			expectedOut: "NUMBER 4043.6490 4043.649\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "10.40",
			expectedOut: "NUMBER 10.40 10.4\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "and",
			expectedOut: "AND and null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "foo bar _hello\nbar baz",
			expectedOut: "IDENTIFIER foo null\nIDENTIFIER bar null\nIDENTIFIER _hello null\nIDENTIFIER bar null\nIDENTIFIER baz null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "42\n42.42\n.42\n.42.42\n42.\n123.123\n67.0000",
			expectedOut: "NUMBER 42 42.0\nNUMBER 42.42 42.42\nDOT . null\nNUMBER 42 42.0\nDOT . null\nNUMBER 42.42 42.42\nNUMBER 42. 42.0\nNUMBER 123.123 123.123\nNUMBER 67.0000 67.0\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "\"bar\n\"foo\n\"\n\"bar",
			expectedOut: "STRING \"bar\" bar\nIDENTIFIER foo null\nSTRING \"\" \nIDENTIFIER bar null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "() #    {}\n@\n$\n+++\n// Let's Go!\n+++\n#",
			expectedOut: "LEFT_PAREN ( null\nRIGHT_PAREN ) null\nLEFT_BRACE { null\nRIGHT_BRACE } null\nPLUS + null\nPLUS + null\nPLUS + null\nPLUS + null\nPLUS + null\nPLUS + null\nEOF  null\n",
			expectedErr: "[line 1] Error: Unexpected character: #\n[line 2] Error: Unexpected character: @\n[line 3] Error: Unexpected character: $\n[line 7] Error: Unexpected character: #\n",
		},
		{
			input:       "(   \n )",
			expectedOut: "LEFT_PAREN ( null\nRIGHT_PAREN ) null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "()// Comment\n// foo ()\n/{}\n// foo",
			expectedOut: "LEFT_PAREN ( null\nRIGHT_PAREN ) null\nSLASH / null\nLEFT_BRACE { null\nRIGHT_BRACE } null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       "={===}",
			expectedOut: "EQUAL = null\nLEFT_BRACE { null\nEQUAL_EQUAL == null\nEQUAL = null\nRIGHT_BRACE } null\nEOF  null\n",
			expectedErr: "",
		},
		{
			input:       ",.$(#",
			expectedOut: "COMMA , null\nDOT . null\nLEFT_PAREN ( null\nEOF  null\n",
			expectedErr: "[line 1] Error: Unexpected character: $\n[line 1] Error: Unexpected character: #\n",
		},
		{
			input:       "(()",
			expectedOut: "LEFT_PAREN ( null\nLEFT_PAREN ( null\nRIGHT_PAREN ) null\nEOF  null\n",
			expectedErr: "",
		},
	}

	focusedTest := false
	for _, tt := range tests {
		if tt.only {
			focusedTest = true
		}
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if focusedTest && !tt.only {
				t.Skip()
			}

			l := lox.NewLox()

			result, err := l.Tokenize(strings.NewReader(tt.input))
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(result.Errors) > 0 && tt.expectedErr == "" {
				t.Errorf("unexpected tokenize errors: %v", result.Errors)
			}

			out := ""
			for _, token := range result.Tokens {
				out += token.String()
			}
			if out != tt.expectedOut {
				t.Errorf("\nexpected output:\n%q\ngot:\n%q\n", tt.expectedOut, out)
			}

			outErr := ""
			for _, err := range result.Errors {
				outErr += err.Error()
			}
			if outErr != tt.expectedErr {
				t.Errorf("\nexpected error:\n%q\ngot:\n%q\n", tt.expectedErr, outErr)
			}
		})
	}
}
