package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.Default()

	if len(os.Args) < 3 {
		logger.Fatal("Missing arguments")
	}

	switch os.Args[1] {
	case "tokenize":
		{
			filePath := os.Args[2]
			fileBuf, err := os.ReadFile(filePath)
			if err != nil {
				logger.Fatalf("Failed to read file: %v", err)
			}

			if len(fileBuf) == 0 {
				fmt.Println("EOF  null")
				os.Exit(0)
			}

			logger.Fatalln("Not implemented yet")
		}
	default:
		{
			logger.Fatalf("Unknown command %s\n", os.Args[1])
		}
	}
}
