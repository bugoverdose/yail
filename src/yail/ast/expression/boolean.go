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

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}
