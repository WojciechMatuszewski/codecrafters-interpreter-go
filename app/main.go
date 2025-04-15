package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/app/lox"
)

func main() {
	logger := log.Default()

	if len(os.Args) < 3 {
		logger.Fatal("Missing arguments")
	}

	l := lox.NewLox()

	switch os.Args[1] {
	case "tokenize":
		{
			filePath := os.Args[2]
			file, err := os.Open(filePath)
			if err != nil {
				logger.Fatalf("Failed to read file: %v", err)
			}
			defer file.Close()

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
	case "parse":
		{
			filePath := os.Args[2]
			file, err := os.Open(filePath)
			if err != nil {
				logger.Fatalf("Failed to read file: %v", err)
			}
			defer file.Close()

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
	default:
		{
			logger.Fatalf("Unknown command %s\n", os.Args[1])
		}
	}
}
