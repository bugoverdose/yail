package expression

import (
	"strconv"
	"yail/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func NewIntegerLiteral(literal string) *IntegerLiteral {
	value, err := strconv.ParseInt(literal, 0, 64)
	if err != nil {
		return nil
	}
	return &IntegerLiteral{
		Token: token.NewInteger(literal),
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
