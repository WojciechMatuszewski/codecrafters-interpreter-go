package lox_test

import (
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func TestParse(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		l := lox.NewLox()
		l.Parse()
	})
}
