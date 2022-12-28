package statement

import (
	"yail/ast/expression"
	"yail/token"
)

type ExpressionStatement struct {
	Token      token.Token 
	Expression expression.Expression
}

func NewExpressionStatement(tok token.Token, expr expression.Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Token: tok,
		Expression: expr,
	}
}

func (statement *ExpressionStatement) statementNode() {}

func (statement *ExpressionStatement) TokenLiteral() string { 
	return statement.Token.Literal 
}

func (statement *ExpressionStatement) String() string {
	return statement.Expression.String() + ";"
}
