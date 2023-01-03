package ast

import (
	"bytes"
	"strings"
	"yail/token"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*IdentifierExpression
	Body       *BlockStatement
}

func NewFunctionLiteral(parameters []*IdentifierExpression, body *BlockStatement) *FunctionLiteral {
	return &FunctionLiteral{
		Token:      token.FUNCTION_TOKEN,
		Parameters: parameters,
		Body:       body,
	}
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString("{ ")
	out.WriteString(fl.Body.String())
	out.WriteString(" }")
	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func NewFunctionCall(functionIdentifier Expression, arguments []Expression) *CallExpression {
	return &CallExpression{
		Token:     token.LEFT_PARENTHESIS_TOKEN,
		Function:  functionIdentifier,
		Arguments: arguments,
	}
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
