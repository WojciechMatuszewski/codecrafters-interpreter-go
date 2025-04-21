package lox

import (
	"fmt"
)

type Expression interface {
	accept(visitor expressionVisitor) any
}

func FormatExpression(expr Expression) string {
	return fmt.Sprintf("%v", expr.accept(&printer{}))
}

type expressionVisitor interface {
	visitBinaryExpression(expr *binaryExpression) any
	visitGroupingExpression(expr *groupingExpression) any
	visitLiteralExpression(expr *literalExpression) any
	visitUnaryExpression(expr *unaryExpression) any
}

// Example: 2+3
type binaryExpression struct {
	Left     Expression
	Right    Expression
	Operator string
}

func (b *binaryExpression) accept(visitor expressionVisitor) any {
	return visitor.visitBinaryExpression(b)
}

// Example: (<EXPR>)
type groupingExpression struct {
	Expression Expression
}

func (g *groupingExpression) accept(visitor expressionVisitor) any {
	return visitor.visitGroupingExpression(g)
}

// Example: 3
type literalExpression struct {
	Value any
}

func (l *literalExpression) accept(visitor expressionVisitor) any {
	return visitor.visitLiteralExpression(l)
}

// Example: -x
type unaryExpression struct {
	Right    Expression
	Operator string
}

func (u *unaryExpression) accept(visitor expressionVisitor) any {
	return visitor.visitUnaryExpression(u)
}

type printer struct{}

func (p *printer) visitBinaryExpression(expr *binaryExpression) any {
	return fmt.Sprintf("(%v %s %v)", expr.Operator, expr.Left.accept(p), expr.Right.accept(p))
}

func (p *printer) visitGroupingExpression(expr *groupingExpression) any {
	return parenthesize("group", p, expr.Expression)
}

func (p *printer) visitLiteralExpression(expr *literalExpression) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p *printer) visitUnaryExpression(expr *unaryExpression) any {
	return fmt.Sprintf("(%s %v)", expr.Operator, expr.Right.accept(p))
}

func parenthesize(name string, visitor expressionVisitor, exprs ...Expression) string {
	output := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		output += fmt.Sprintf(" %v", expr.accept(visitor))
	}
	output += ")"
	return output
}
