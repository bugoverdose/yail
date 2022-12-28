package expression

import (
	"yail/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func NewIntegerLiteral(tok token.Token, value int64) *IntegerLiteral {
	return &IntegerLiteral{
		Token: tok,
		Value: value,
	}
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
