package ast

import (
	"bytes"
	"yail/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func NewIf(condition Expression, consequence *BlockStatement) *IfExpression {
	return &IfExpression{
		Token:       token.IF_TOKEN,
		Condition:   condition,
		Consequence: consequence,
		Alternative: nil,
	}
}

func NewIfElse(condition Expression, consequence, alternative *BlockStatement) *IfExpression {
	return &IfExpression{
		Token:       token.IF_TOKEN,
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

func (i *IfExpression) expressionNode() {}
func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())
	if i.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(i.Alternative.String())
	}
	return out.String()
}
