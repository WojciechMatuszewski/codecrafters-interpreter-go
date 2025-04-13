package lox_test

import (
	"bytes"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestParse(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		r := bytes.NewReader([]byte("3==3==3==4"))

		l := lox.NewLox()
		l.Parse(r)
	})
}
