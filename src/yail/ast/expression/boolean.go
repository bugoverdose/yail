package expression

import (
	"yail/token"
)

type Boolean struct {
	Token token.Token
	Value bool
}

var (
	TRUE  = Boolean{Token: token.NewKeyword(token.TRUE), Value: true}
	FALSE = Boolean{Token: token.NewKeyword(token.FALSE), Value: false}
)

func GetPooledBoolean(value bool) *Boolean {
	if value {
		return &TRUE
	}
	return &FALSE
}

func (il *Boolean) expressionNode() {}

func (il *Boolean) TokenLiteral() string {
	return il.Token.Literal
}

func (il *Boolean) String() string {
	return il.Token.Literal
}
