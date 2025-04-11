package lox

import (
	"fmt"
	"io"
)

type Expr interface {
	Accept(visitor ExprVisitor) any
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *BinaryExpr) any
	VisitGroupingExpr(expr *GroupingExpr) any
	VisitLiteralExpr(expr *LiteralExpr) any
	VisitUnaryExpr(expr *UnaryExpr) any
}

// Example: 2+3
type BinaryExpr struct {
	Left     Expr
	Right    Expr
	Operator string
}

func (b *BinaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(b)
}

// Example: (<EXPR>)
type GroupingExpr struct {
	Expression Expr
}

func (g *GroupingExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitGroupingExpr(g)
}

// Example: 3
type LiteralExpr struct {
	Value any
}

func (l *LiteralExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteralExpr(l)
}

// Example: -x
type UnaryExpr struct {
	Right    Expr
	Operator string
}

func (u *UnaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(u)
}

func (l *Lox) Parse(r io.Reader) {

	expr := BinaryExpr{
		Left: &GroupingExpr{
			Expression: &BinaryExpr{
				Left: &LiteralExpr{
					Value: 1,
				},
				Operator: "+",
				Right: &LiteralExpr{
					Value: 2,
				},
			},
		},
		Operator: "*",
		Right: &GroupingExpr{
			Expression: &BinaryExpr{
				Left: &LiteralExpr{
					Value: 4,
				},
				Operator: "-",
				Right: &LiteralExpr{
					Value: 3,
				},
			},
		},
	}

	outRegular := expr.Accept(&PrinterVisitor{})
	outRPN := expr.Accept(&RPNPrinterVisitor{})

	fmt.Println(outRegular)
	fmt.Println(outRPN)

}

func parenthesize(name string, visitor ExprVisitor, exprs ...Expr) string {
	output := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		output += fmt.Sprintf(" %v", expr.Accept(visitor))
	}
	output += ")"
	return output
}

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
