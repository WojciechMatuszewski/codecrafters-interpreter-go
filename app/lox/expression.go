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

func PrintExpression(expr Expr) {
	visitor := PrinterVisitor{}
	out := expr.Accept(&visitor)
	fmt.Println(out)
}
