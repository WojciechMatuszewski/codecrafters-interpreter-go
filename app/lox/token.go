package lox

import (
	"fmt"
	"strconv"
)

type TokenType string

const (
	// Single-character tokens
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	COMMA       TokenType = "COMMA"
	DOT         TokenType = "DOT"
	MINUS       TokenType = "MINUS"
	PLUS        TokenType = "PLUS"
	SEMICOLON   TokenType = "SEMICOLON"
	SLASH       TokenType = "SLASH"
	STAR        TokenType = "STAR"

	// One or two character tokens
	BANG          TokenType = "BANG"
	BANG_EQUAL    TokenType = "BANG_EQUAL"
	EQUAL         TokenType = "EQUAL"
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"
	GREATER       TokenType = "GREATER"
	GREATER_EQUAL TokenType = "GREATER_EQUAL"
	LESS          TokenType = "LESS"
	LESS_EQUAL    TokenType = "LESS_EQUAL"

	// Literals
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	// Keywords
	AND    TokenType = "AND"
	CLASS  TokenType = "CLASS"
	ELSE   TokenType = "ELSE"
	FALSE  TokenType = "FALSE"
	FUN    TokenType = "FUN"
	FOR    TokenType = "FOR"
	IF     TokenType = "IF"
	NIL    TokenType = "NIL"
	OR     TokenType = "OR"
	PRINT  TokenType = "PRINT"
	RETURN TokenType = "RETURN"
	SUPER  TokenType = "SUPER"
	THIS   TokenType = "THIS"
	TRUE   TokenType = "TRUE"
	VAR    TokenType = "VAR"
	WHILE  TokenType = "WHILE"

	EOF TokenType = "EOF"
)

var TokenLexemes = map[TokenType]string{
	// Single-character tokens
	LEFT_PAREN:  "(",
	RIGHT_PAREN: ")",
	LEFT_BRACE:  "{",
	RIGHT_BRACE: "}",
	COMMA:       ",",
	DOT:         ".",
	MINUS:       "-",
	PLUS:        "+",
	SEMICOLON:   ";",
	SLASH:       "/",
	STAR:        "*",

	// One or two character tokens
	BANG:          "!",
	BANG_EQUAL:    "!=",
	EQUAL:         "=",
	EQUAL_EQUAL:   "==",
	GREATER:       ">",
	GREATER_EQUAL: ">=",
	LESS:          "<",
	LESS_EQUAL:    "<=",

	// Keywords
	AND:    "and",
	CLASS:  "class",
	ELSE:   "else",
	FALSE:  "false",
	FUN:    "fun",
	FOR:    "for",
	IF:     "if",
	NIL:    "nil",
	OR:     "or",
	PRINT:  "print",
	RETURN: "return",
	SUPER:  "super",
	THIS:   "this",
	TRUE:   "true",
	VAR:    "var",
	WHILE:  "while",

	// Special tokens
	EOF: "",
}

// Map of Keywords where key is the keyword string and value is the TokenType
var Keywords = map[string]TokenType{
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

type Token struct {
	Type    TokenType
	Lexme   *string
	Literal any
}

func NewToken(tokenType TokenType) Token {
	lexme, found := TokenLexemes[tokenType]
	if !found {
		panic(fmt.Sprintf("could not find lexme for tokenType: %v", tokenType))
	}

	return Token{
		Type:    tokenType,
		Lexme:   &lexme,
		Literal: nil,
	}
}

func NewStringToken(value string) Token {
	lexme := fmt.Sprintf("\"%v\"", value)

	return Token{
		Type:    STRING,
		Lexme:   &lexme,
		Literal: value,
	}
}

func NewIdentifierToken(value string) Token {
	return Token{
		Type:    IDENTIFIER,
		Lexme:   &value,
		Literal: nil,
	}
}

func NewNumberToken(value string) Token {
	literal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return Token{
		Type:    NUMBER,
		Lexme:   &value,
		Literal: literal,
	}
}

func (t Token) String() string {
	lexme := "null"
	if t.Lexme != nil {
		lexme = *t.Lexme
	}

	literal := "null"
	if t.Literal == nil {
		return fmt.Sprintf("%v %v %v\n", t.Type, lexme, literal)
	}

	literal = fmt.Sprintf("%v", t.Literal)
	if t.Type != NUMBER {
		return fmt.Sprintf("%v %v %v\n", t.Type, lexme, literal)
	}

	formatted, err := formatToDecimalString(lexme)
	if err != nil {
		panic(fmt.Errorf("could not parse number to literal: %w", err))
	}

	return fmt.Sprintf("%v %v %v\n", t.Type, lexme, formatted)
}

func formatToDecimalString(value string) (string, error) {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse number from string: %w", err)
	}

	if num == float64(int(num)) {
		return fmt.Sprintf("%.1f", num), nil
	}

	return fmt.Sprintf("%v", num), nil
}
