package ast

import (
	"bytes"
	"strings"
	"yail/token"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func NewArrayLiteral(elements []Expression) *ArrayLiteral {
	return &ArrayLiteral{
		Token:    token.LEFT_BRACKET_TOKEN,
		Elements: elements,
	}
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	var elements []string
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type IndexAccessExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func NewIndexAccess(left, index Expression) *IndexAccessExpression {
	return &IndexAccessExpression{
		Token: token.LEFT_BRACKET_TOKEN,
		Left:  left,
		Index: index,
	}
}

func (ie *IndexAccessExpression) expressionNode() {}
func (ie *IndexAccessExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IndexAccessExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
