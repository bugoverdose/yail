package expression

import (
	"bytes"
	"yail/token"
)

type Prefix struct {
	Token     token.Token
	Operator  string
	RightNode Expression
}

func NewPrefix(operatorToken token.Token, rightNode Expression) *Prefix {
	return &Prefix{
		Token:     operatorToken,
		Operator:  operatorToken.Literal,
		RightNode: rightNode,
	}
}

func (p *Prefix) expressionNode() {}

func (p *Prefix) TokenLiteral() string {
	return p.Token.Literal
}

func (p *Prefix) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.RightNode.String())
	out.WriteString(")")
	return out.String() // (!true)
}
