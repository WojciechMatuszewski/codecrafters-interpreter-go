package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode"
)

type ErrUnexpectedToken struct {
	Message string
	Line    int
}

func (e ErrUnexpectedToken) Error() string {
	return fmt.Sprintf("[line %v] Error: %s\n", e.Line, e.Message)
}

type TokenizeResult struct {
	Tokens []Token
	Errors []ErrUnexpectedToken
}

func (l *Lox) Tokenize(r io.Reader) (TokenizeResult, error) {
	reader := bufio.NewReader(r)

	line := 1

	var tokenErrors []ErrUnexpectedToken
	var tokens []Token

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				tokens = append(tokens, NewToken(EOF))
				break
			}

			return TokenizeResult{}, fmt.Errorf("failed to read token: %w", err)
		}

		sb := string(b)
		switch sb {
		case TokenLexemes[LEFT_BRACE]:
			{
				tokens = append(tokens, NewToken(LEFT_BRACE))
			}
		case TokenLexemes[RIGHT_BRACE]:
			{
				tokens = append(tokens, NewToken(RIGHT_BRACE))
			}
		case TokenLexemes[LEFT_PAREN]:
			{
				tokens = append(tokens, NewToken(LEFT_PAREN))
			}
		case TokenLexemes[RIGHT_PAREN]:
			{
				tokens = append(tokens, NewToken(RIGHT_PAREN))
			}
		case TokenLexemes[COMMA]:
			{
				tokens = append(tokens, NewToken(COMMA))
			}
		case TokenLexemes[DOT]:
			{
				tokens = append(tokens, NewToken(DOT))
			}
		case TokenLexemes[MINUS]:
			{
				tokens = append(tokens, NewToken(MINUS))
			}
		case TokenLexemes[PLUS]:
			{
				tokens = append(tokens, NewToken(PLUS))
			}
		case TokenLexemes[SEMICOLON]:
			{
				tokens = append(tokens, NewToken(SEMICOLON))
			}
		case TokenLexemes[STAR]:
			{
				tokens = append(tokens, NewToken(STAR))
			}
		case TokenLexemes[BANG]:
			{
				matches, err := matchNextToken(reader, NewToken(EQUAL))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, NewToken(BANG_EQUAL))
					reader.ReadByte()
				} else {
					tokens = append(tokens, NewToken(BANG))
				}
			}
		case TokenLexemes[EQUAL]:
			{
				matches, err := matchNextToken(reader, NewToken(EQUAL))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, NewToken(EQUAL_EQUAL))
					reader.ReadByte()
				} else {
					tokens = append(tokens, NewToken(EQUAL))
				}
			}
		case TokenLexemes[LESS]:
			{
				matches, err := matchNextToken(reader, NewToken(EQUAL))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, NewToken(LESS_EQUAL))
					reader.ReadByte()
				} else {
					tokens = append(tokens, NewToken(LESS))
				}
			}
		case TokenLexemes[GREATER]:
			{
				matches, err := matchNextToken(reader, NewToken(EQUAL))
				if err != nil {
					return TokenizeResult{}, fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					tokens = append(tokens, NewToken(GREATER_EQUAL))
					reader.ReadByte()
				} else {
					tokens = append(tokens, NewToken(GREATER))
				}
			}
		case TokenLexemes[SLASH]:
			{
				matches, err := matchNextToken(reader, NewToken(SLASH))
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
					tokens = append(tokens, NewToken(SLASH))
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

						tokenErrors = append(tokenErrors, ErrUnexpectedToken{Line: line, Message: "Unterminated string."})
						break
					}

					st := string(bt)
					if st == "\n" {
						line += 1
						continue
					}

					if st == "\"" {
						tokens = append(tokens, NewStringToken(contents))
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

					tokens = append(tokens, NewNumberToken(number))
				} else if isAlphaNumeric(sb) {
					content := sb

					for isAlphaNumeric(peekNext(reader)) {
						b, _ := reader.ReadByte()
						content += string(b)
					}

					keyword, found := Keywords[content]
					if found {
						tokens = append(tokens, NewToken(keyword))
					} else {
						tokens = append(tokens, NewIdentifierToken(content))
					}

				} else {
					tokenErrors = append(tokenErrors, ErrUnexpectedToken{Line: line, Message: fmt.Sprintf("Unexpected character: %v", sb)})
				}
			}
		}
	}

	return TokenizeResult{
		Tokens: tokens,
		Errors: tokenErrors,
	}, nil

	// for _, token := range tokens {
	// 	_, err := outW.Write([]byte(token.String()))
	// 	if err != nil {
	// 		return fmt.Errorf("failed to write: %w", err)
	// 	}
	// }

	// if len(tokenizeErrors) > 0 {
	// 	errors := strings.Join(tokenizeErrors, "")
	// 	_, err := errW.Write([]byte(errors))
	// 	if err != nil {
	// 		return fmt.Errorf("failed to write: %w", err)
	// 	}

	// 	return ErrUnexpectedTokens
	// }

	// return nil
}

func matchNextToken(r *bufio.Reader, matchToken Token) (bool, error) {
	nextB, err := r.Peek(1)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false, nil
		}

		return false, fmt.Errorf("failed to peek: %w", err)
	}

	nextLexme := string(nextB)
	return nextLexme == *matchToken.Lexme, nil
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
