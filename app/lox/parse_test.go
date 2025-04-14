package lox_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestParse(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		r := bytes.NewReader([]byte("true"))

		l := lox.NewLox()
		expr, err := l.Parse(r)
		if err != nil {
			t.Fatalf("Failed to parse: %v", err)
		}

		visitor := lox.PrinterVisitor{}
		out := expr.Accept(&visitor)
		fmt.Println(out)
	})
}
