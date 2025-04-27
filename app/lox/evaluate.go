package lox

import (
	"fmt"
	"io"
)

type RuntimeError struct {
	line    int
	message string
}

func (re RuntimeError) Error() string {
	return fmt.Sprintf("%s\n[line %v]", re.message, re.line)
}

func (l *Lox) Evaluate(r io.Reader) (any, error) {
	expr, err := l.Parse(r)
	if err != nil {
		panic(err)
	}

	return expr.accept(&evaluator{})
}

type evaluator struct{}

func (e *evaluator) visitBinaryExpression(expr *binaryExpression) (any, error) {
	left, err := expr.Left.accept(e)
	if err != nil {
		return nil, err
	}

	right, err := expr.Right.accept(e)
	if err != nil {
		return nil, err
	}

	switch *expr.Operator.Lexeme {
	case TokenLexemes[MINUS]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return lv - rv, nil
		}
	case TokenLexemes[PLUS]:
		{
			sum, err := add(left, right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers or two strings."}
			}

			return sum, nil
		}
	case TokenLexemes[STAR]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return lv * rv, nil
		}
	case TokenLexemes[SLASH]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return lv / rv, nil
		}
	case TokenLexemes[GREATER]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return lv > rv, nil

		}
	case TokenLexemes[GREATER_EQUAL]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return lv >= rv, nil
		}
	case TokenLexemes[LESS]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return rv < lv, nil

		}
	case TokenLexemes[LESS_EQUAL]:
		{
			lv, err := toF64(left)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}
			rv, err := toF64(right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers."}
			}

			return lv <= rv, nil
		}
	case TokenLexemes[EQUAL_EQUAL]:
		{
			result, err := isEqual(left, right)
			if err != nil {
				panic(err)
			}

			return result, nil

		}
	case TokenLexemes[BANG_EQUAL]:
		{

			result, err := isEqual(left, right)
			if err != nil {
				panic(err)
			}

			return !result, nil
		}
	default:
		{
			panic(fmt.Errorf("unknown operator for binary operation"))
		}
	}
}

func (e *evaluator) visitGroupingExpression(expr *groupingExpression) (any, error) {
	return expr.Expression.accept(e)
}

func (e *evaluator) visitLiteralExpression(expr *literalExpression) (any, error) {
	return expr.Value, nil
}

func (e *evaluator) visitUnaryExpression(expr *unaryExpression) (any, error) {
	value, err := expr.Right.accept(e)
	if err != nil {
		return nil, err
	}

	switch *expr.Operator.Lexeme {
	case TokenLexemes[MINUS]:
		{
			v, err := toF64(value)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operand must be a number."}
			}

			return -v, nil
		}
	case TokenLexemes[PLUS]:
		{
			return value, nil
		}
	case TokenLexemes[BANG]:
		{
			return !isTruthy(value), nil
		}
	default:
		{
			panic(fmt.Errorf("unknown operator"))
		}
	}
}

func toF64(v any) (float64, error) {
	switch v := v.(type) {
	case float64:
		{
			return v, nil
		}
	default:
		{
			return 0, fmt.Errorf("value %v is not float64", v)
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
