package lox

import (
	"fmt"
	"io"
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
			sum, err := add(left, right)
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
	case TokenLexemes[GREATER]:
		{
			return mustFloat64(left) > mustFloat64(right)
		}
	case TokenLexemes[GREATER_EQUAL]:
		{

			return mustFloat64(left) >= mustFloat64(right)
		}
	case TokenLexemes[LESS]:
		{
			return mustFloat64(left) < mustFloat64(right)
		}
	case TokenLexemes[LESS_EQUAL]:
		{
			return mustFloat64(left) <= mustFloat64(right)
		}
	case TokenLexemes[EQUAL_EQUAL]:
		{
			result, err := isEqual(left, right)
			if err != nil {
				panic(err)
			}
			return result

		}
	case TokenLexemes[BANG_EQUAL]:
		{

			result, err := isEqual(left, right)
			if err != nil {
				panic(err)
			}

			return !result
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
			return value
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
	default:
		{
			panic(fmt.Sprintf("Value %v is not float64", v))
		}
	}

}

func add(left, right any) (any, error) {
	switch lv := left.(type) {
	case float64:
		if rv, ok := right.(float64); ok {
			return lv + rv, nil
		}
	case string:
		if rv, ok := right.(string); ok {
			return lv + rv, nil
		}
	}

	return nil, fmt.Errorf("unsupported operand types for +: %T and %T", left, right)
}

func isEqual(left, right any) (bool, error) {
	if left == nil && right == nil {
		return true, nil
	}

	if left == nil {
		return false, nil
	}

	switch lv := left.(type) {
	case float64:
		if rv, ok := right.(float64); ok {
			return lv == rv, nil
		}
	case string:
		if rv, ok := right.(string); ok {
			return lv == rv, nil
		}
	}

	return false, nil
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
