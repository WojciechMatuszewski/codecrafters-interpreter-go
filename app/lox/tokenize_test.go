package lox_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input       string
		expectedOut string
		expectedErr string
	}{
		{
			input:       "({*.,+*})",
			expectedOut: "LEFT_PAREN ( null\nLEFT_BRACE { null\nSTAR * null\nDOT . null\nCOMMA , null\nPLUS + null\nSTAR * null\nRIGHT_BRACE } null\nRIGHT_PAREN ) null\nEOF  null\n",
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

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := lox.NewLox()

			var outBuf, errBuf bytes.Buffer
			err := l.Tokenize(strings.NewReader(tt.input), &outBuf, &errBuf)
			if err != nil {
				if !errors.Is(err, lox.ErrUnexpectedTokens) || errors.Is(err, lox.ErrUnexpectedTokens) && tt.expectedErr == "" {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if errBuf.String() != tt.expectedErr {
				t.Errorf("expected error %q, got %q", tt.expectedErr, errBuf.String())
			}

			if outBuf.String() != tt.expectedOut {
				t.Errorf("expected output %q, got %q", tt.expectedOut, outBuf.String())
			}
		})
	}
}
