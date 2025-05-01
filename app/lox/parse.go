package lox

import (
	"fmt"
	"io"
)

type SyntaxError struct {
	line    int
	message string
}

func (se SyntaxError) Error() string {
	return fmt.Sprintf("[line %v] %v", se.line, se.message)
}

func (l *Lox) Parse(r io.Reader) ([]Statement, error) {
	result, err := l.Tokenize(r)
	if err != nil {
		return nil, err
	}
	parser := newParser(result.Tokens)
	statements, err := parser.parse()

	return statements, err
}

type parser struct {
	tokens  []token
	current int
}

func newParser(tokens []token) *parser {
	return &parser{tokens: tokens, current: 0}
}

func (p *parser) parse() ([]Statement, error) {
	p.current = 0
	var statements []Statement

	for !p.isAtEnd() {
		statement, err := p.statement()
		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)
	}

	return statements, nil
}

func (p *parser) statement() (Statement, error) {
	if p.match(PRINT) {
		expr, err := p.expression()
		if err != nil {
			panic(err)
		}

		if p.match(SEMICOLON) {
			return &printStatement{
				expr: expr,
			}, nil
		}

		return nil, SyntaxError{line: 1, message: "Expect ';' after value."}
	}

	return p.expressionStatement()
}

func (p *parser) expressionStatement() (Statement, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	if p.match(SEMICOLON) {
		return &exprStatement{expr: expr}, nil
	}

	return nil, SyntaxError{line: 1, message: "Expect ';' after expression."}
}

func (p *parser) expression() (Expression, error) {
	return p.equality()
}

func (p *parser) equality() (Expression, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}

		expr = &binaryExpression{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *parser) comparison() (Expression, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = &binaryExpression{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *parser) term() (Expression, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = &binaryExpression{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *parser) factor() (Expression, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(SLASH, STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &binaryExpression{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *parser) unary() (Expression, error) {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &unaryExpression{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

func (p *parser) primary() (Expression, error) {
	if p.match(FALSE) {
		return &literalExpression{Value: false}, nil
	}

	if p.match(TRUE) {
		return &literalExpression{Value: true}, nil
	}

	if p.match(NIL) {
		return &literalExpression{Value: nil}, nil
	}

	if p.match(NUMBER, STRING) {
		return &literalExpression{Value: p.previous().Literal}, nil
	}

	if p.match(LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, SyntaxError{line: 1, message: "Error at ')': Expect expression."}
		}

		if p.match(RIGHT_PAREN) {
			return &groupingExpression{Expression: expr}, nil
		}

		return nil, SyntaxError{line: 1, message: "Error at ')': Expect expression."}
	}

	return nil, SyntaxError{line: 1, message: "Expect expression."}
}

func (p *parser) match(tokenTypes ...tokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *parser) check(tokenType tokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == tokenType
}

func (p *parser) peek() token {
	return p.tokens[p.current]
}

func (p *parser) previous() token {
	return p.tokens[p.current-1]
}

func (p *parser) advance() token {
	if !p.isAtEnd() {
		p.current += 1
	}

	return p.previous()
}
