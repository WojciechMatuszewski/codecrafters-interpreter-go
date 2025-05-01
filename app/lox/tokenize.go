package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode"
)

type UnexpectedTokenError struct {
	Message string
	Line    int
}

func (e UnexpectedTokenError) Error() string {
	return fmt.Sprintf("[line %v] Error: %s\n", e.Line, e.Message)
}

type TokenizeResult struct {
	Tokens []token
	Errors []UnexpectedTokenError
}

func (l *Lox) Tokenize(r io.Reader) (TokenizeResult, error) {
	reader := bufio.NewReader(r)
	line := 1

	var tokenErrors []UnexpectedTokenError
	var tokens []token

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				tokens = append(tokens, newToken(EOF, line))
				break
			}

			return TokenizeResult{}, fmt.Errorf("failed to read token: %w", err)
		}

		sb := string(b)
		switch sb {
		case tokenLexemes[LEFT_BRACE]:
			{
				tokens = append(tokens, newToken(LEFT_BRACE, line))
			}
		case tokenLexemes[RIGHT_BRACE]:
			{
				tokens = append(tokens, newToken(RIGHT_BRACE, line))
			}
		case tokenLexemes[LEFT_PAREN]:
			{
				tokens = append(tokens, newToken(LEFT_PAREN, line))
			}
		case tokenLexemes[RIGHT_PAREN]:
			{
				tokens = append(tokens, newToken(RIGHT_PAREN, line))
			}
		case tokenLexemes[COMMA]:
			{
				tokens = append(tokens, newToken(COMMA, line))
			}
		case tokenLexemes[DOT]:
			{
				tokens = append(tokens, newToken(DOT, line))
			}
		case tokenLexemes[MINUS]:
			{
				tokens = append(tokens, newToken(MINUS, line))
			}
		case tokenLexemes[PLUS]:
			{
				tokens = append(tokens, newToken(PLUS, line))
			}
		case tokenLexemes[SEMICOLON]:
			{
				tokens = append(tokens, newToken(SEMICOLON, line))
			}
		case tokenLexemes[STAR]:
			{
				tokens = append(tokens, newToken(STAR, line))
			}
		case tokenLexemes[BANG]:
			{
				matches, err := matchNextToken(reader, newToken(EQUAL, line))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, newToken(BANG_EQUAL, line))
					reader.ReadByte()
				} else {
					tokens = append(tokens, newToken(BANG, line))
				}
			}
		case tokenLexemes[EQUAL]:
			{
				matches, err := matchNextToken(reader, newToken(EQUAL, line))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, newToken(EQUAL_EQUAL, line))
					reader.ReadByte()
				} else {
					tokens = append(tokens, newToken(EQUAL, line))
				}
			}
		case tokenLexemes[LESS]:
			{
				matches, err := matchNextToken(reader, newToken(EQUAL, line))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, newToken(LESS_EQUAL, line))
					reader.ReadByte()
				} else {
					tokens = append(tokens, newToken(LESS, line))
				}
			}
		case tokenLexemes[GREATER]:
			{
				matches, err := matchNextToken(reader, newToken(EQUAL, line))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, newToken(GREATER_EQUAL, line))
					reader.ReadByte()
				} else {
					tokens = append(tokens, newToken(GREATER, line))
				}
			}
		case tokenLexemes[SLASH]:
			{
				matches, err := matchNextToken(reader, newToken(SLASH, line))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					_, err = reader.ReadString('\n')
					if err != nil {
						if !errors.Is(err, io.EOF) {
							return TokenizeResult{}, fmt.Errorf("failed to consume rest of the comment: %w", err)
						}
					}
					line += 1
				} else {
					tokens = append(tokens, newToken(SLASH, line))
				}
			}
		case "\"":
			{
				contents := ""
				for {
					bt, err := reader.ReadByte()
					if err != nil {
						if !errors.Is(err, io.EOF) {
							return TokenizeResult{}, fmt.Errorf("failed to consume the string: %w", err)
						}

						tokenErrors = append(tokenErrors, UnexpectedTokenError{Line: line, Message: "Unterminated string."})
						break
					}

					st := string(bt)
					if st == "\n" {
						line += 1
						contents += st

						continue
					}

					if st == "\"" {
						tokens = append(tokens, newStringToken(contents, line))
						break
					}

					contents += st
				}
			}
		case " ":
		case "\r":
		case "\t":
			{
				continue
			}
		case "\n":
			{
				line += 1
			}
		default:
			{
				if isDigit(sb) {
					number := sb

					for isDigit(peekNext(reader)) {
						b, _ := reader.ReadByte()
						number += string(b)
					}

					if peekNext(reader) == "." {
						b, _ := reader.ReadByte()
						number += string(b)

						for isDigit(peekNext(reader)) {
							b, _ := reader.ReadByte()
							number += string(b)
						}
					}

					tokens = append(tokens, newNumberToken(number, line))
				} else if isAlphaNumeric(sb) {
					content := sb

					for isAlphaNumeric(peekNext(reader)) {
						b, _ := reader.ReadByte()
						content += string(b)
					}

					keyword, found := keywords[content]
					if found {
						tokens = append(tokens, newToken(keyword, line))
					} else {
						tokens = append(tokens, newIdentifierToken(content, line))
					}

				} else {
					tokenErrors = append(tokenErrors, UnexpectedTokenError{Line: line, Message: fmt.Sprintf("Unexpected character: %v", sb)})
				}
			}
		}
	}

	return TokenizeResult{
		Tokens: tokens,
		Errors: tokenErrors,
	}, nil

}

func matchNextToken(r *bufio.Reader, matchToken token) (bool, error) {
	nextB, err := r.Peek(1)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false, nil
		}

		return false, fmt.Errorf("failed to peek: %w", err)
	}

	nextLexme := string(nextB)
	return nextLexme == *matchToken.Lexeme, nil
}

func isDigit(s string) bool {
	if len(s) != 1 {
		return false
	}

	r := rune(s[0])
	return unicode.IsDigit(r)
}

func isAlpha(s string) bool {
	if len(s) != 1 {
		return false
	}

	r := rune(s[0])
	return unicode.IsLetter(r) || r == '_'
}

func isAlphaNumeric(s string) bool {
	return isDigit(s) || isAlpha(s)
}

func peekNext(r *bufio.Reader) string {
	next, err := r.Peek(1)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return ""
		}

		return ""
	}

	return string(next)
}
