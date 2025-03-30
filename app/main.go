package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logger := log.Default()

	if len(os.Args) < 3 {
		logger.Fatal("Missing arguments")
	}

	lox := NewLox()

	switch os.Args[1] {
	case "tokenize":
		{
			filePath := os.Args[2]
			file, err := os.Open(filePath)
			if err != nil {
				logger.Fatalf("Failed to read file: %v", err)
			}

			err = lox.Run(file, os.Stdout)
			if err != nil {
				logger.Fatalf("Failed to execute command: %v", err)
			}
		}
	default:
		{
			logger.Fatalf("Unknown command %s\n", os.Args[1])
		}
	}
}

type TokenType string

const (
	// Single-character tokens
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	COMMA       TokenType = ","
	DOT         TokenType = "."
	MINUS       TokenType = "-"
	PLUS        TokenType = "+"
	SEMICOLON   TokenType = ";"
	SLASH       TokenType = "/"
	STAR        TokenType = "*"

	// One or two character tokens
	BANG          TokenType = "!"
	BANG_EQUAL    TokenType = "!="
	EQUAL         TokenType = "="
	EQUAL_EQUAL   TokenType = "=="
	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="
	LESS          TokenType = "<"
	LESS_EQUAL    TokenType = "<="

	// Literals (keeping descriptive names as they don't have single symbols)
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	// Keywords (using actual keywords)
	AND    TokenType = "and"
	CLASS  TokenType = "class"
	ELSE   TokenType = "else"
	FALSE  TokenType = "false"
	FUN    TokenType = "fun"
	FOR    TokenType = "for"
	IF     TokenType = "if"
	NIL    TokenType = "nil"
	OR     TokenType = "or"
	PRINT  TokenType = "print"
	RETURN TokenType = "return"
	SUPER  TokenType = "super"
	THIS   TokenType = "this"
	TRUE   TokenType = "true"
	VAR    TokenType = "var"
	WHILE  TokenType = "while"

	EOF TokenType = "EOF"
)

type Lox struct{}

func NewLox() *Lox {
	return &Lox{}
}

func (l *Lox) Run(r io.Reader, w io.Writer) error {
	reader := bufio.NewReader(r)

	output := ""

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				output += "EOF  null\n"
				break
			}

			return fmt.Errorf("failed to read token: %w", err)
		}

		token := string(b)
		switch token {
		case string(LEFT_BRACE):
			{
				output += fmt.Sprintf("LEFT_BRACE %v null\n", LEFT_BRACE)
			}
		case string(RIGHT_BRACE):
			{
				output += fmt.Sprintf("RIGHT_BRACE %v null\n", RIGHT_BRACE)
			}
		case string(LEFT_PAREN):
			{
				output += fmt.Sprintf("LEFT_PAREN %v null\n", LEFT_PAREN)
			}
		case string(RIGHT_PAREN):
			{
				output += fmt.Sprintf("RIGHT_PAREN %v null\n", RIGHT_PAREN)
			}
		case string(COMMA):
			{
				output += fmt.Sprintf("COMMA %v null\n", COMMA)
			}
		case string(DOT):
			{
				output += fmt.Sprintf("DOT %v null\n", DOT)
			}
		case string(MINUS):
			{
				output += fmt.Sprintf("MINUS %v null\n", MINUS)
			}
		case string(PLUS):
			{
				output += fmt.Sprintf("PLUS %v null\n", PLUS)
			}
		case string(SEMICOLON):
			{
				output += fmt.Sprintf("SEMICOLON %v null\n", SEMICOLON)
			}
		case string(SLASH):
			{
				output += fmt.Sprintf("SLASH %v null\n", SLASH)
			}
		case string(STAR):
			{
				output += fmt.Sprintf("STAR %v null\n", STAR)
			}
		default:
			return fmt.Errorf("token %v not implemented", token)
		}
	}

	_, err := w.Write([]byte(output))
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	return nil
}
