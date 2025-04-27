package lox

import (
	"fmt"
)

type Expression interface {
	accept(visitor expressionVisitor) (any, error)
}

func FormatExpression(expr Expression) string {
	result, err := expr.accept(&printer{})
	if err != nil {
		panic(fmt.Sprintf("failed to format the expression: %v", err))
	}

	return fmt.Sprintf("%v", result)
}

type expressionVisitor interface {
	visitBinaryExpression(expr *binaryExpression) (any, error)
	visitGroupingExpression(expr *groupingExpression) (any, error)
	visitLiteralExpression(expr *literalExpression) (any, error)
	visitUnaryExpression(expr *unaryExpression) (any, error)
}

// Example: 2+3
type binaryExpression struct {
	Left     Expression
	Right    Expression
	Operator Token
}

func (b *binaryExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitBinaryExpression(b)
}

// Example: (<EXPR>)
type groupingExpression struct {
	Expression Expression
}

func (g *groupingExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitGroupingExpression(g)
}

// Example: 3
type literalExpression struct {
	Value any
}

func (l *literalExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitLiteralExpression(l)
}

// Example: -x
type unaryExpression struct {
	Right    Expression
	Operator Token
}

func (u *unaryExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitUnaryExpression(u)
}

type printer struct{}

func (p *printer) visitBinaryExpression(expr *binaryExpression) (any, error) {
	left, err := expr.Left.accept(p)
	if err != nil {
		return nil, err
	}
	right, err := expr.Right.accept(p)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("(%v %s %v)", expr.Operator, left, right), nil
}

func (p *printer) visitGroupingExpression(expr *groupingExpression) (any, error) {
	return parenthesize("group", p, expr.Expression)
}

func (p *printer) visitLiteralExpression(expr *literalExpression) (any, error) {
	switch n := expr.Value.(type) {
	case float64:
		{
			if n == float64(int(n)) {
				return fmt.Sprintf("%.1f", n), nil
			}
			return fmt.Sprintf("%v", n), nil
		}
	default:
		{
			if expr.Value == nil {
				return "nil", nil
			}
			return fmt.Sprintf("%v", expr.Value), nil
		}
	}
}

func (p *printer) visitUnaryExpression(expr *unaryExpression) (any, error) {
	right, err := expr.Right.accept(p)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("(%s %v)", expr.Operator, right), nil
}

func parenthesize(name string, visitor expressionVisitor, exprs ...Expression) (string, error) {
	output := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		result, err := expr.accept(visitor)
		if err != nil {
			panic(fmt.Sprintf("failed to parenthesize the expression: %v", err))
		}

		output += fmt.Sprintf(" %v", result)
	}
	output += ")"
	return output, nil
}
