package lox

import "fmt"

type PrinterVisitor struct{}

func (p *PrinterVisitor) VisitBinaryExpr(expr *BinaryExpr) any {
	return fmt.Sprintf("%v %s %v", expr.Left.Accept(p), expr.Operator, expr.Right.Accept(p))
}

func (p *PrinterVisitor) VisitGroupingExpr(expr *GroupingExpr) any {
	return parenthesize("group", p, expr.Expression)
}

func (p *PrinterVisitor) VisitLiteralExpr(expr *LiteralExpr) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p *PrinterVisitor) VisitUnaryExpr(expr *UnaryExpr) any {
	return fmt.Sprintf("%s %v", expr.Operator, expr.Right.Accept(p))
}

type RPNPrinterVisitor struct{}

func (p *RPNPrinterVisitor) VisitBinaryExpr(expr *BinaryExpr) any {
	left := expr.Left.Accept(p)
	right := expr.Right.Accept(p)

	return fmt.Sprintf("%v %v %s", left, right, expr.Operator)
}

func (p *RPNPrinterVisitor) VisitGroupingExpr(expr *GroupingExpr) any {
	return fmt.Sprintf("%v", expr.Expression.Accept(p))
}

func (p *RPNPrinterVisitor) VisitLiteralExpr(expr *LiteralExpr) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p *RPNPrinterVisitor) VisitUnaryExpr(expr *UnaryExpr) any {
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
