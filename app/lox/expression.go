package lox

type Expr interface {
	Accept(visitor ExprVisitor) any
}

type ExprVisitor interface {
	visitBinaryExpr(expr *binaryExpr) any
	visitGroupingExpr(expr *groupingExpr) any
	visitLiteralExpr(expr *literalExpr) any
	visitUnaryExpr(expr *unaryExpr) any
}

// Example: 2+3
type binaryExpr struct {
	Left     Expr
	Right    Expr
	Operator string
}

func (b *binaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.visitBinaryExpr(b)
}

// Example: (<EXPR>)
type groupingExpr struct {
	Expression Expr
}

func (g *groupingExpr) Accept(visitor ExprVisitor) any {
	return visitor.visitGroupingExpr(g)
}

// Example: 3
type literalExpr struct {
	Value any
}

func (l *literalExpr) Accept(visitor ExprVisitor) any {
	return visitor.visitLiteralExpr(l)
}

// Example: -x
type unaryExpr struct {
	Right    Expr
	Operator string
}

func (u *unaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.visitUnaryExpr(u)
}
