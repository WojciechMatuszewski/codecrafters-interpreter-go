package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.Default()

	tokenizeCmd := flag.NewFlagSet("tokenize", flag.ExitOnError)
	filename := tokenizeCmd.String("file", "", "Input file to tokenize")

	if len(os.Args) < 2 {
		logger.Fatal("Missing arguments")
	}

	switch os.Args[1] {
	case "tokenize":
		{
			tokenizeCmd.Parse(os.Args[2:])
			if *filename == "" {
				tokenizeCmd.PrintDefaults()
				logger.Fatal("Error: -file flag is required")
			}

			_, err := readFile(*filename)
			if err != nil {
				logger.Fatalf("Failed to read file: %e", err)
			}

			logger.Fatal("Scanner not implemented")
		}
	default:
		{
			logger.Fatalf("Unknown command %s\n", os.Args[1])
		}
	}
}

func readFile(path string) ([]byte, error) {
	fileBuf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(fileBuf) == 0 {
		return nil, fmt.Errorf("file %s must not be empty", path)
	}

	return fileBuf, nil
}
