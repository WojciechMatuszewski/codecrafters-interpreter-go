package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func (l *Lox) Tokenize(r io.Reader, outW, errW io.Writer) error {
	reader := bufio.NewReader(r)

	successOutput := ""
	errOutput := ""
	line := 1

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				successOutput += "EOF  null\n"
				break
			}

			return fmt.Errorf("failed to read token: %w", err)
		}

		token := string(b)
		switch token {
		case string(LEFT_BRACE):
			{
				successOutput += fmt.Sprintf("LEFT_BRACE %v null\n", LEFT_BRACE)
			}
		case string(RIGHT_BRACE):
			{
				successOutput += fmt.Sprintf("RIGHT_BRACE %v null\n", RIGHT_BRACE)
			}
		case string(LEFT_PAREN):
			{
				successOutput += fmt.Sprintf("LEFT_PAREN %v null\n", LEFT_PAREN)
			}
		case string(RIGHT_PAREN):
			{
				successOutput += fmt.Sprintf("RIGHT_PAREN %v null\n", RIGHT_PAREN)
			}
		case string(COMMA):
			{
				successOutput += fmt.Sprintf("COMMA %v null\n", COMMA)
			}
		case string(DOT):
			{
				successOutput += fmt.Sprintf("DOT %v null\n", DOT)
			}
		case string(MINUS):
			{
				successOutput += fmt.Sprintf("MINUS %v null\n", MINUS)
			}
		case string(PLUS):
			{
				successOutput += fmt.Sprintf("PLUS %v null\n", PLUS)
			}
		case string(SEMICOLON):
			{
				successOutput += fmt.Sprintf("SEMICOLON %v null\n", SEMICOLON)
			}
		case string(STAR):
			{
				successOutput += fmt.Sprintf("STAR %v null\n", STAR)
			}
		case string(BANG):
			{

				matches, err := matchNextToken(reader, string(EQUAL))
				if err != nil {
					return fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					successOutput += fmt.Sprintf("BANG_EQUAL %v null\n", string(BANG_EQUAL))
					reader.ReadByte()
				} else {
					successOutput += fmt.Sprintf("BANG %v null\n", string(BANG))
				}
			}
		case string(EQUAL):
			{
				matches, err := matchNextToken(reader, string(EQUAL))
				if err != nil {
					return fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					successOutput += fmt.Sprintf("EQUAL_EQUAL %v null\n", string(EQUAL_EQUAL))
					reader.ReadByte()
				} else {
					successOutput += fmt.Sprintf("EQUAL %v null\n", string(EQUAL))
				}
			}
		case string(LESS):
			{
				matches, err := matchNextToken(reader, string(EQUAL))
				if err != nil {
					return fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					successOutput += fmt.Sprintf("LESS_EQUAL %v null\n", string(LESS_EQUAL))
					reader.ReadByte()
				} else {
					successOutput += fmt.Sprintf("LESS %v null\n", string(LESS))
				}
			}
		case string(GREATER):
			{
				matches, err := matchNextToken(reader, string(EQUAL))
				if err != nil {
					return fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					successOutput += fmt.Sprintf("GREATER_EQUAL %v null\n", string(GREATER_EQUAL))
					reader.ReadByte()
				} else {
					successOutput += fmt.Sprintf("GREATER %v null\n", string(GREATER))
				}
			}
		case string(SLASH):
			{
				matches, err := matchNextToken(reader, string(SLASH))
				if err != nil {
					return fmt.Errorf("failed to match next token: %w", err)
				}
				if matches {
					_, err = reader.ReadString('\n')
					if err != nil {
						if !errors.Is(err, io.EOF) {
							return fmt.Errorf("failed to consume rest of the comment: %w", err)
						}
					}
					line += 1
				} else {
					successOutput += fmt.Sprintf("SLASH %v null\n", string(SLASH))
				}
			}
		case "\"":
			{
				contents := ""
				for {
					bt, err := reader.ReadByte()
					if err != nil {
						if !errors.Is(err, io.EOF) {
							return fmt.Errorf("failed to consume the string: %w", err)
						}

						errOutput += fmt.Sprintf("[line %v] Error: Unterminated string.\n", line)
						break
					}

					st := string(bt)
					if st == "\n" {
						line += 1
						continue
					}

					if st == "\"" {
						successOutput += fmt.Sprintf("STRING \"%v\" %v\n", contents, contents)
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
				if isDigit(token) {
					contents := token

					for isDigit(peekNext(reader)) {
						b, _ := reader.ReadByte()
						contents += string(b)
					}

					if peekNext(reader) == "." {
						b, _ := reader.ReadByte()
						contents += string(b)

						for isDigit(peekNext(reader)) {
							b, _ := reader.ReadByte()
							contents += string(b)
						}
					}

					formatted, err := formatToDecimalString(contents)
					if err != nil {
						return err
					}

					successOutput += fmt.Sprintf("NUMBER %v %v\n", contents, formatted)
				} else if isAlphaNumeric(token) {
					contents := token

					for isAlphaNumeric(peekNext(reader)) {
						b, _ := reader.ReadByte()
						contents += string(b)
					}

					_, found := keywords[contents]
					if found {
						successOutput += fmt.Sprintf("%v %v null\n", strings.ToUpper(contents), contents)
					} else {
						successOutput += fmt.Sprintf("IDENTIFIER %v null\n", contents)
					}

				} else {
					errOutput += fmt.Sprintf("[line %v] Error: Unexpected character: %v\n", line, token)
				}
			}
		}
	}

	_, err := outW.Write([]byte(successOutput))
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	if errOutput != "" {
		_, err := errW.Write([]byte(errOutput))
		if err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}
	}

	if errOutput != "" {
		return ErrUnexpectedTokens
	}

	return nil
}

func matchNextToken(r *bufio.Reader, nextToken string) (bool, error) {
	nextB, err := r.Peek(1)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false, nil
		}

		return false, fmt.Errorf("failed to peek: %w", err)
	}

	nextT := string(nextB)
	return nextT == nextToken, nil
}

func isDigit(token string) bool {
	if len(token) == 0 {
		return false
	}

	r := rune(token[0])
	return unicode.IsDigit(r)
}

func isAlpha(token string) bool {
	if len(token) == 0 {
		return false
	}

	r := rune(token[0])
	return unicode.IsLetter(r) || r == '_'
}

func isAlphaNumeric(token string) bool {
	return isDigit(token) || isAlpha(token)
}

func peekNext(r *bufio.Reader) string {
	nextB, err := r.Peek(1)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return ""
		}

		return ""
	}

	return string(nextB)
}

func formatToDecimalString(value string) (string, error) {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse number from string: %w", err)
	}

	if num == float64(int64(num)) {
		return fmt.Sprintf("%.1f", num), nil
	}

	return value, nil
}
