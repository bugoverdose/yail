package expression

import (
	"bytes"
	"yail/ast/statement"
	"yail/token"
)

type If struct {
	Token       token.Token
	Condition   Expression
	Consequence *statement.Block
	Alternative *statement.Block
}

func NewIf(condition Expression, consequence *statement.Block) *If {
	return &If{
		Token:       token.IF_TOKEN,
		Condition:   condition,
		Consequence: consequence,
		Alternative: nil,
	}
}

func NewIfElse(condition Expression, consequence, alternative *statement.Block) *If {
	return &If{
		Token:       token.IF_TOKEN,
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

func (i *If) expressionNode() {}

func (i *If) TokenLiteral() string {
	return i.Token.Literal
}

func (i *If) String() string {
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
