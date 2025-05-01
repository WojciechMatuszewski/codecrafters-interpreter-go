package lox

import (
	"fmt"
	"strconv"
)

type tokenType string

const (
	// Single-character tokens
	LEFT_PAREN  tokenType = "LEFT_PAREN"
	RIGHT_PAREN tokenType = "RIGHT_PAREN"
	LEFT_BRACE  tokenType = "LEFT_BRACE"
	RIGHT_BRACE tokenType = "RIGHT_BRACE"
	COMMA       tokenType = "COMMA"
	DOT         tokenType = "DOT"
	MINUS       tokenType = "MINUS"
	PLUS        tokenType = "PLUS"
	SEMICOLON   tokenType = "SEMICOLON"
	SLASH       tokenType = "SLASH"
	STAR        tokenType = "STAR"

	// One or two character tokens
	BANG          tokenType = "BANG"
	BANG_EQUAL    tokenType = "BANG_EQUAL"
	EQUAL         tokenType = "EQUAL"
	EQUAL_EQUAL   tokenType = "EQUAL_EQUAL"
	GREATER       tokenType = "GREATER"
	GREATER_EQUAL tokenType = "GREATER_EQUAL"
	LESS          tokenType = "LESS"
	LESS_EQUAL    tokenType = "LESS_EQUAL"

	// Literals
	IDENTIFIER tokenType = "IDENTIFIER"
	STRING     tokenType = "STRING"
	NUMBER     tokenType = "NUMBER"

	// Keywords
	AND    tokenType = "AND"
	CLASS  tokenType = "CLASS"
	ELSE   tokenType = "ELSE"
	FALSE  tokenType = "FALSE"
	FUN    tokenType = "FUN"
	FOR    tokenType = "FOR"
	IF     tokenType = "IF"
	NIL    tokenType = "NIL"
	OR     tokenType = "OR"
	PRINT  tokenType = "PRINT"
	RETURN tokenType = "RETURN"
	SUPER  tokenType = "SUPER"
	THIS   tokenType = "THIS"
	TRUE   tokenType = "TRUE"
	VAR    tokenType = "VAR"
	WHILE  tokenType = "WHILE"

	EOF tokenType = "EOF"
)

var tokenLexemes = map[tokenType]string{
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

// Map of keywords where key is the keyword string and value is the TokenType
var keywords = map[string]tokenType{
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

type token struct {
	Type    tokenType
	Lexeme  *string
	Literal any
	Line    int
}

func newToken(tokenType tokenType, line int) token {
	lexme, found := tokenLexemes[tokenType]
	if !found {
		panic(fmt.Sprintf("could not find lexme for tokenType: %v", tokenType))
	}

	return token{
		Type:    tokenType,
		Lexeme:  &lexme,
		Literal: nil,
		Line:    line,
	}
}

func newStringToken(value string, line int) token {
	lexme := fmt.Sprintf("\"%v\"", value)

	return token{
		Type:    STRING,
		Lexeme:  &lexme,
		Literal: value,
		Line:    line,
	}
}

func newIdentifierToken(value string, line int) token {
	return token{
		Type:    IDENTIFIER,
		Lexeme:  &value,
		Literal: nil,
		Line:    line,
	}
}

func newNumberToken(value string, line int) token {
	literal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return token{
		Type:    NUMBER,
		Lexeme:  &value,
		Literal: literal,
		Line:    line,
	}
}

func (t token) String() string {
	lexme := "null"
	if t.Lexeme != nil {
		lexme = *t.Lexeme
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
