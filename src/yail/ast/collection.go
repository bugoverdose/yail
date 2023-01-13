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

type HashMapLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func NewHashMapLiteral(pairs map[Expression]Expression) *HashMapLiteral {
	return &HashMapLiteral{
		Token: token.LEFT_BRACE_TOKEN,
		Pairs: pairs,
	}
}

func (hl *HashMapLiteral) expressionNode() {}
func (hl *HashMapLiteral) TokenLiteral() string {
	return hl.Token.Literal
}
func (hl *HashMapLiteral) String() string {
	var out bytes.Buffer
	var pairs []string
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type CollectionAccessExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func NewCollectionAccess(left, index Expression) *CollectionAccessExpression {
	return &CollectionAccessExpression{
		Token: token.LEFT_BRACKET_TOKEN,
		Left:  left,
		Index: index,
	}
}

func (c *CollectionAccessExpression) expressionNode() {}
func (c *CollectionAccessExpression) TokenLiteral() string {
	return c.Token.Literal
}
func (c *CollectionAccessExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(c.Left.String())
	out.WriteString("[")
	out.WriteString(c.Index.String())
	out.WriteString("])")
	return out.String()
}
