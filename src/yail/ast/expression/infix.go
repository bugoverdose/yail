package expression

import (
	"bytes"
	"yail/token"
)

type Infix struct {
	Token     token.Token
	LeftNode  Expression
	Operator  string
	RightNode Expression
}

func NewInfix(leftNode Expression, operatorToken token.Token, rightNode Expression) *Infix {
	return &Infix{
		Token:     operatorToken,
		LeftNode:  leftNode,
		Operator:  operatorToken.Literal,
		RightNode: rightNode,
	}
}

func (i *Infix) expressionNode() {}

func (i *Infix) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Infix) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.LeftNode.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.RightNode.String())
	out.WriteString(")")
	return out.String() // (10 + 20)
}
