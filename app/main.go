package main

import (
	"errors"
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

			err = l.Tokenize(file, os.Stdout, os.Stderr)
			if err != nil {
				if errors.Is(err, lox.ErrUnexpectedTokens) {
					os.Exit(65)
				}

				logger.Fatalf("Failed to execute command: %v", err)
			}
		}
	default:
		{
			logger.Fatalf("Unknown command %s\n", os.Args[1])
		}
	}
}
