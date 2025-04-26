package lox

import (
	"fmt"
	"io"
	"strconv"
)

func (l *Lox) Evaluate(r io.Reader) any {
	expr, err := l.Parse(r)
	if err != nil {
		panic(err)
	}

	return expr.accept(&evaluator{})
}

type evaluator struct{}

func (e *evaluator) visitBinaryExpression(expr *binaryExpression) any {
	left := expr.Left.accept(e)
	right := expr.Right.accept(e)

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

func (e *evaluator) visitGroupingExpression(expr *groupingExpression) any {
	return expr.Expression.accept(e)
}

func (e *evaluator) visitLiteralExpression(expr *literalExpression) any {
	return expr.Value
}

func (e *evaluator) visitUnaryExpression(expr *unaryExpression) any {
	value := expr.Right.accept(e)
	switch expr.Operator {
	case TokenLexemes[MINUS]:
		{
			return -mustFloat64(value)
		}
	case TokenLexemes[PLUS]:
		{
			return mustAdditive(value)
		}
	case TokenLexemes[BANG]:
		{
			return !isTruthy(value)
		}
	default:
		{
			return nil
		}
	}
}

func mustFloat64(v any) float64 {
	switch v := v.(type) {
	case float64:
		{
			return v
		}
	case string:
		{
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}

			return num
		}
	default:
		{
			panic(fmt.Sprintf("Value %v is not string or float64", v))
		}
	}

}

func mustAdditive(v any) any {
	switch v := v.(type) {
	case string:
		{
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return v
			}

			return num
		}
	case float64:
		{
			return v
		}
	default:
		{
			panic(fmt.Sprintf("Value %v is not string or float64", v))
		}
	}
}

func add(left, right any) (any, error) {
	switch lv := left.(type) {
	case float64:
		if r, ok := right.(float64); ok {
			return lv + r, nil
		}
	case string:
		if r, ok := right.(string); ok {
			return lv + r, nil
		}
	}

	return nil, fmt.Errorf("unsupported operand types for +: %T and %T", left, right)
}

func isTruthy(v any) bool {
	if v == nil {
		return false
	}

	switch v := v.(type) {
	case bool:
		{
			return v
		}
	default:
		{
			return true
		}
	}
}
