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
	expr, err := l.ParseExpression(r)
	if err != nil {
		return nil, err
	}

	return expr.accept(&evaluator{})
}

// TODO: We do not want to return anything here.
// We should write to STDIO, but how to make it robust?
func (l *Lox) Run(r io.Reader) (any, error) {
	statements, err := l.Parse(r)
	if err != nil {
		return nil, err
	}

	var result string
	for _, statement := range statements {
		out, err := statement.accept(&evaluator{})
		if err != nil {
			return nil, err
		}

		if out != nil {
			result += fmt.Sprintf("%v\n", out)
		}

	}

	if result == "" {
		return nil, nil
	}

	return result, nil
}

type evaluator struct{}

func (e *evaluator) visitPrintStatement(statement *printStatement) (any, error) {
	out, err := statement.expr.accept(e)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("%v", out), nil
}

func (e *evaluator) visitExprStatement(statement *exprStatement) (any, error) {
	_, err := statement.expr.accept(e)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

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
	case tokenLexemes[MINUS]:
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
	case tokenLexemes[PLUS]:
		{
			sum, err := add(left, right)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operands must be two numbers or two strings."}
			}

			return sum, nil
		}
	case tokenLexemes[STAR]:
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
	case tokenLexemes[SLASH]:
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
	case tokenLexemes[GREATER]:
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
	case tokenLexemes[GREATER_EQUAL]:
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
	case tokenLexemes[LESS]:
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
	case tokenLexemes[LESS_EQUAL]:
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
	case tokenLexemes[EQUAL_EQUAL]:
		{
			result, err := isEqual(left, right)
			if err != nil {
				panic(err)
			}

			return result, nil

		}
	case tokenLexemes[BANG_EQUAL]:
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
	case tokenLexemes[MINUS]:
		{
			v, err := toF64(value)
			if err != nil {
				return nil, RuntimeError{line: expr.Operator.Line, message: "Operand must be a number."}
			}

			return -v, nil
		}
	case tokenLexemes[PLUS]:
		{
			return value, nil
		}
	case tokenLexemes[BANG]:
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
	case bool:
		if rv, ok := right.(bool); ok {
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
