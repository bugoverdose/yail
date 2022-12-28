package expression

import (
	"yail/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func NewIdentifier(value string) *Identifier {
	return &Identifier{
		Token: token.NewIdentifier(value),
		Value: value,
	}
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
