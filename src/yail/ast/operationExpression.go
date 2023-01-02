package ast

import (
	"bytes"
	"yail/token"
)

type PrefixExpression struct {
	Token     token.Token
	Operator  string
	RightNode Expression
}

func NewPrefix(operatorToken token.Token, rightNode Expression) *PrefixExpression {
	return &PrefixExpression{
		Token:     operatorToken,
		Operator:  operatorToken.Literal,
		RightNode: rightNode,
	}
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}
func (p *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.RightNode.String())
	out.WriteString(")")
	return out.String() // (!true)
}

type InfixExpression struct {
	Token     token.Token
	LeftNode  Expression
	Operator  string
	RightNode Expression
}

func NewInfix(leftNode Expression, operatorToken token.Token, rightNode Expression) *InfixExpression {
	return &InfixExpression{
		Token:     operatorToken,
		LeftNode:  leftNode,
		Operator:  operatorToken.Literal,
		RightNode: rightNode,
	}
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.LeftNode.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.RightNode.String())
	out.WriteString(")")
	return out.String() // (10 + 20)
}
