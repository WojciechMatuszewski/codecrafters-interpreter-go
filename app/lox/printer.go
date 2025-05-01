package lox

import "fmt"

type printer struct{}

func Format(statements []Statement) string {
	result := ""
	printer := &printer{}

	for _, statement := range statements {
		out, err := statement.accept(printer)
		if err != nil {
			panic(err)
		}

		result += fmt.Sprintf("%v", out)
	}

	return result
}

func (p *printer) visitExprStatement(statement *exprStatement) (any, error) {
	return parenthesize(";", p, statement.expr)
}

func (p *printer) visitPrintStatement(statement *printStatement) (any, error) {
	return parenthesize(";", p, statement.expr)
}

func (p *printer) visitBinaryExpression(expr *binaryExpression) (any, error) {
	left, err := expr.Left.accept(p)
	if err != nil {
		return nil, err
	}
	right, err := expr.Right.accept(p)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("(%v %s %v)", *expr.Operator.Lexeme, left, right), nil
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

	return fmt.Sprintf("(%s %v)", *expr.Operator.Lexeme, right), nil
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
