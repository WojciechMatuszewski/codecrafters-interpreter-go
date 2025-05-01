package lox

type Expression interface {
	accept(visitor expressionVisitor) (any, error)
}

type expressionVisitor interface {
	visitBinaryExpression(expr *binaryExpression) (any, error)
	visitGroupingExpression(expr *groupingExpression) (any, error)
	visitLiteralExpression(expr *literalExpression) (any, error)
	visitUnaryExpression(expr *unaryExpression) (any, error)
}

// Example: 2+3
type binaryExpression struct {
	Left     Expression
	Right    Expression
	Operator token
}

func (b *binaryExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitBinaryExpression(b)
}

// Example: (<EXPR>)
type groupingExpression struct {
	Expression Expression
}

func (g *groupingExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitGroupingExpression(g)
}

// Example: 3
type literalExpression struct {
	Value any
}

func (l *literalExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitLiteralExpression(l)
}

// Example: -x
type unaryExpression struct {
	Right    Expression
	Operator token
}

func (u *unaryExpression) accept(visitor expressionVisitor) (any, error) {
	return visitor.visitUnaryExpression(u)
}
