package lox

import (
	"fmt"
	"io"
)

func (l *Lox) Evaluate(r io.Reader) {
	expr, err := l.Parse(r)
	if err != nil {
		panic(err)
	}

	out := expr.accept(&interpreter{})
	fmt.Println(out)
}
