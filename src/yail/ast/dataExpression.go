package ast

import (
	"yail/token"
)

var (
	TRUE  = &BooleanExpression{Token: token.New(token.TRUE), Value: true}
	FALSE = &BooleanExpression{Token: token.New(token.FALSE), Value: false}
	NULL  = &NullExpression{Token: token.NULL_TOKEN}
)

type IntegerLiteralExpression struct {
	Token token.Token
	Value int64
}

func NewIntegerLiteral(tok token.Token, value int64) *IntegerLiteralExpression {
	return &IntegerLiteralExpression{
		Token: tok,
		Value: value,
	}
}

func (il *IntegerLiteralExpression) expressionNode() {}
func (il *IntegerLiteralExpression) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteralExpression) String() string {
	return il.Token.Literal
}

type BooleanExpression struct {
	Token token.Token
	Value bool
}

func GetPooledBoolean(value bool) *BooleanExpression {
	if value {
		return TRUE
	}
	return FALSE
}

func (b *BooleanExpression) expressionNode() {}
func (b *BooleanExpression) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BooleanExpression) String() string {
	return b.Token.Literal
}

type NullExpression struct {
	Token token.Token
}

func (n *NullExpression) expressionNode() {}
func (n *NullExpression) TokenLiteral() string {
	return n.Token.Literal
}
func (n *NullExpression) String() string {
	return n.Token.Literal
}

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
