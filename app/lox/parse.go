package lox

import (
	"io"
)

func (l *Lox) Parse(r io.Reader) (Expr, error) {
	result, err := l.Tokenize(r)
	if err != nil {
		return nil, err
	}

	parser := newParser(result.Tokens)
	expr := parser.parse()

	return expr, nil
}

type parser struct {
	tokens  []Token
	current int
}

func newParser(tokens []Token) *parser {
	return &parser{tokens: tokens, current: 0}
}

func (p *parser) parse() Expr {
	return p.expression()
}

func (p *parser) expression() Expr {
	return p.equality()
}

func (p *parser) equality() Expr {
	expr := p.comparison()

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &binaryExpr{Left: expr, Operator: *operator.Lexme, Right: right}
	}

	return expr
}

func (p *parser) comparison() Expr {
	expr := p.term()

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &binaryExpr{Left: expr, Operator: *operator.Lexme, Right: right}
	}

	return expr
}

func (p *parser) term() Expr {
	expr := p.factor()

	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = &binaryExpr{Left: expr, Operator: *operator.Lexme, Right: right}
	}

	return expr
}

func (p *parser) factor() Expr {
	expr := p.unary()

	for p.match(SLASH, STAR) {
		operator := p.previous()
		right := p.unary()
		expr = &binaryExpr{Left: expr, Operator: *operator.Lexme, Right: right}
	}

	return expr
}

func (p *parser) unary() Expr {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		return &unaryExpr{Operator: *operator.Lexme, Right: right}
	}

	return p.primary()
}

func (p *parser) primary() Expr {
	if p.match(FALSE) {
		return &literalExpr{Value: false}
	}

	if p.match(TRUE) {
		return &literalExpr{Value: true}
	}

	if p.match(NIL) {
		return &literalExpr{Value: "nil"}
	}

	if p.match(NUMBER, STRING) {
		return &literalExpr{Value: *p.previous().Literal}
	}

	panic("Unhandled primary expression case")
}

func (p *parser) match(tokenTypes ...TokenType) bool {
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

func (p *parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == tokenType
}

func (p *parser) peek() Token {
	return p.tokens[p.current]
}

func (p *parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *parser) advance() Token {
	if !p.isAtEnd() {
		p.current += 1
	}

	return p.previous()
}
