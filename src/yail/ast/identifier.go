package ast

import (
	"yail/token"
)

type IdentifierExpression struct {
	Token token.Token
	Value string
}

func NewIdentifierFrom(value string) *IdentifierExpression {
	return &IdentifierExpression{
		Token: token.NewIdentifier(value),
		Value: value,
	}
}

func NewIdentifier(tok token.Token) *IdentifierExpression {
	return &IdentifierExpression{
		Token: tok,
		Value: tok.Literal,
	}
}

func (i *IdentifierExpression) expressionNode() {}
func (i *IdentifierExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i *IdentifierExpression) String() string {
	return i.Value
}
