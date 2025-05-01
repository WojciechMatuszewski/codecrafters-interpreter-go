package lox

type Statement interface {
	accept(visitor statementVisitor) (any, error)
}

type statementVisitor interface {
	visitPrintStatement(ps *printStatement) (any, error)
	visitExprStatement(ps *exprStatement) (any, error)
}

type printStatement struct {
	expr Expression
}

type exprStatement struct {
	expr Expression
}

func (ps *printStatement) accept(visitor statementVisitor) (any, error) {
	return visitor.visitPrintStatement(ps)
}

func (ps *exprStatement) accept(visitor statementVisitor) (any, error) {
	return visitor.visitExprStatement(ps)
}
