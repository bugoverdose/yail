package expression

import (
	"yail/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func NewIdentifierFrom(value string) *Identifier {
	return &Identifier{
		Token: token.NewIdentifier(value),
		Value: value,
	}
}

func NewIdentifier(tok token.Token) *Identifier {
	return &Identifier{
		Token: tok,
		Value: tok.Literal,
	}
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
