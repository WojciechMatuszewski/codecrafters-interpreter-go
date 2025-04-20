package lox

import (
	"fmt"
	"strconv"
)

type Expression interface {
	accept(visitor expressionVisitor) any
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

func FormatExpression(expr Expression) string {
	return fmt.Sprintf("%v", expr.accept(&printer{}))
}

type interpreter struct{}

func (p *interpreter) visitBinaryExpression(expr *binaryExpression) any {
	left := expr.Left.accept(p)
	right := expr.Right.accept(p)

	switch expr.Operator {
	case TokenLexemes[MINUS]:
		{
			return mustFloat64(left) - mustFloat64(right)
		}
	case TokenLexemes[PLUS]:
		{
			sum, err := add(mustAdditive(left), mustAdditive(right))
			if err != nil {
				panic(err)
			}

			return sum
		}
	case TokenLexemes[STAR]:
		{
			return mustFloat64(left) * mustFloat64(right)
		}
	case TokenLexemes[SLASH]:
		{
			return mustFloat64(left) / mustFloat64(right)
		}
	default:
		{
			return nil
		}
	}
}

func (p *interpreter) visitGroupingExpression(expr *groupingExpression) any {
	return expr.accept(p)
}

func (p *interpreter) visitLiteralExpression(expr *literalExpression) any {
	return expr.Value
}

func (p *interpreter) visitUnaryExpression(expr *unaryExpression) any {
	value := expr.Right.accept(p)
	switch expr.Operator {
	case TokenLexemes[MINUS]:
		{
			return -mustFloat64(value)
		}
	case TokenLexemes[PLUS]:
		{
			return mustAdditive(value)
		}
	default:
		{
			return nil
		}
	}
}

func mustFloat64(v any) float64 {
	switch n := v.(type) {
	case float64:
		{
			return n
		}
	case string:
		{
			num, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				panic(err)
			}

			return num
		}
	default:
		{
			panic("value is neither a float64 nor a string")
		}
	}

}

func mustAdditive(v any) any {
	switch value := v.(type) {
	case string:
		{
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return value
			}

			return num
		}
	case float64:
		{
			return value
		}
	default:
		{
			panic("value is not additive")
		}
	}
}

func add(left, right any) (any, error) {
	switch l := left.(type) {
	case float64:
		if r, ok := right.(float64); ok {
			return l + r, nil
		}
	case string:
		if r, ok := right.(string); ok {
			return l + r, nil
		}
	}

	return nil, fmt.Errorf("unsupported operand types for +: %T and %T", left, right)
}
