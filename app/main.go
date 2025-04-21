package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

var logger = log.Default()

const (
	CMD_TOKENIZE = "tokenize"
	CMD_PARSE    = "parse"
	CMD_EVALUATE = "evaluate"
)

func main() {
	if len(os.Args) < 3 {
		logger.Fatal("Missing arguments")
	}

	cmd := os.Args[1]
	switch cmd {
	case CMD_TOKENIZE:
		{
			filePath := os.Args[2]
			tokenize(filePath)
		}
	case CMD_PARSE:
		{
			filePath := os.Args[2]
			parse(filePath)
		}
	case CMD_EVALUATE:
		{
			filePath := os.Args[2]
			evaluate(filePath)
		}
	default:
		{
			logger.Fatalf("Unknown command %s\n", os.Args[1])
		}
	}
}

func evaluate(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	l := lox.NewLox()
	fmt.Fprint(os.Stdout, l.Evaluate(file))
}

func parse(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	l := lox.NewLox()
	expr, err := l.Parse(file)
	if err != nil {
		if errors.As(err, &lox.SyntaxError{}) {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(65)
		}

		logger.Fatalf("Failed to parse the file: %v", err)
	}

	fmt.Fprint(os.Stdout, lox.FormatExpression(expr))
}

func tokenize(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	l := lox.NewLox()
	result, err := l.Tokenize(file)
	if err != nil {
		logger.Fatalf("Failed to execute command: %v", err)
	}

	for _, token := range result.Tokens {
		os.Stdout.Write([]byte(token.String()))
	}

	for _, tokenError := range result.Errors {
		os.Stderr.Write([]byte(tokenError.Error()))
	}

	if len(result.Errors) > 0 {
		os.Exit(65)
	}
}
