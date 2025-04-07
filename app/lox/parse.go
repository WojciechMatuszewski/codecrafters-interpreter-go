package lox

import "fmt"

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

type PrinterVisitor struct{}

func (p *PrinterVisitor) VisitBinaryExpr(expr *BinaryExpr) any {
	return parenthesize(expr.Operator, p, expr.Left, expr.Right)
}

func (p *PrinterVisitor) VisitGroupingExpr(expr *GroupingExpr) any {
	return parenthesize("group", p, expr.Expression)
}

func (p *PrinterVisitor) VisitLiteralExpr(expr *LiteralExpr) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p *PrinterVisitor) VisitUnaryExpr(expr *UnaryExpr) any {
	return parenthesize(expr.Operator, p, expr.Right)
}

func (l *Lox) Parse() {
	expr := BinaryExpr{
		Left: &UnaryExpr{
			Operator: "-",
			Right: &LiteralExpr{
				Value: 123,
			},
		},
		Right: &GroupingExpr{
			Expression: &LiteralExpr{
				Value: 45.67,
			},
		},
		Operator: "*",
	}

	out := expr.Accept(&PrinterVisitor{})

	fmt.Println(out)

}

func parenthesize(name string, visitor ExprVisitor, exprs ...Expr) string {
	output := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		output += fmt.Sprintf(" %v", expr.Accept(visitor))
	}
	output += ")"
	return output
}

// TODO: RPN printer
