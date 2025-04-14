package lox

import "fmt"

type PrinterVisitor struct{}

func (p *PrinterVisitor) visitBinaryExpr(expr *binaryExpr) any {
	return fmt.Sprintf("%v %s %v", expr.Left.Accept(p), expr.Operator, expr.Right.Accept(p))
}

func (p *PrinterVisitor) visitGroupingExpr(expr *groupingExpr) any {
	return parenthesize("group", p, expr.Expression)
}

func (p *PrinterVisitor) visitLiteralExpr(expr *literalExpr) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p *PrinterVisitor) visitUnaryExpr(expr *unaryExpr) any {
	return fmt.Sprintf("%s %v", expr.Operator, expr.Right.Accept(p))
}

type RPNPrinterVisitor struct{}

func (p *RPNPrinterVisitor) visitBinaryExpr(expr *binaryExpr) any {
	left := expr.Left.Accept(p)
	right := expr.Right.Accept(p)

	return fmt.Sprintf("%v %v %s", left, right, expr.Operator)
}

func (p *RPNPrinterVisitor) visitGroupingExpr(expr *groupingExpr) any {
	return fmt.Sprintf("%v", expr.Expression.Accept(p))
}

func (p *RPNPrinterVisitor) visitLiteralExpr(expr *literalExpr) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p *RPNPrinterVisitor) visitUnaryExpr(expr *unaryExpr) any {
	return fmt.Sprintf("%s%v", expr.Operator, expr.Right.Accept(p))
}

func parenthesize(name string, visitor ExprVisitor, exprs ...Expr) string {
	output := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		output += fmt.Sprintf(" %v", expr.Accept(visitor))
	}
	output += ")"
	return output
}
