package lox

import (
	"errors"
)

type Token string

const (
	// Single-character tokens
	LEFT_PAREN  Token = "("
	RIGHT_PAREN Token = ")"
	LEFT_BRACE  Token = "{"
	RIGHT_BRACE Token = "}"
	COMMA       Token = ","
	DOT         Token = "."
	MINUS       Token = "-"
	PLUS        Token = "+"
	SEMICOLON   Token = ";"
	SLASH       Token = "/"
	STAR        Token = "*"

	// One or two character tokens
	BANG          Token = "!"
	BANG_EQUAL    Token = "!="
	EQUAL         Token = "="
	EQUAL_EQUAL   Token = "=="
	GREATER       Token = ">"
	GREATER_EQUAL Token = ">="
	LESS          Token = "<"
	LESS_EQUAL    Token = "<="

	// Literals (keeping descriptive names as they don't have single symbols)
	IDENTIFIER Token = "IDENTIFIER"
	STRING     Token = "STRING"
	NUMBER     Token = "NUMBER"

	// Keywords (using actual keywords)
	AND    Token = "and"
	CLASS  Token = "class"
	ELSE   Token = "else"
	FALSE  Token = "false"
	FUN    Token = "fun"
	FOR    Token = "for"
	IF     Token = "if"
	NIL    Token = "nil"
	OR     Token = "or"
	PRINT  Token = "print"
	RETURN Token = "return"
	SUPER  Token = "super"
	THIS   Token = "this"
	TRUE   Token = "true"
	VAR    Token = "var"
	WHILE  Token = "while"

	EOF Token = "EOF"
)

// Map of keywords where key is the keyword string and value is the TokenType
var keywords = map[string]Token{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

var ErrUnexpectedTokens = errors.New("unexpected tokens found")

type Lox struct{}

func NewLox() *Lox {
	return &Lox{}
}
