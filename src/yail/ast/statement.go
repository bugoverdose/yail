package ast

import (
	"bytes"
	"yail/token"
)

type VariableBindingStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func NewVariableBinding(keyword token.Token, name *IdentifierExpression, value Expression) *VariableBindingStatement {
	if keyword.Type != token.VAR && keyword.Type != token.VAL {
		panic("Invalid implementation: var or val token expected.")
	}
	return &VariableBindingStatement{
		Token: keyword,
		Name:  name,
		Value: value,
	}
}

func (statement *VariableBindingStatement) statementNode() {}
func (statement *VariableBindingStatement) TokenLiteral() string {
	return statement.Token.Literal
}
func (statement *VariableBindingStatement) String() string {
	var out bytes.Buffer
	out.WriteString(statement.TokenLiteral() + " ") // var
	out.WriteString(statement.Name.String())        // a
	out.WriteString(" = ")
	out.WriteString(statement.Value.String()) // 10
	out.WriteString(";")
	return out.String() // var a = 10;
}

type ReassignmentStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func NewReassignment(keyword token.Token, name *IdentifierExpression, value Expression) *ReassignmentStatement {
	if keyword.Type != token.IDENTIFIER {
		panic("Invalid implementation: identifier token expected.")
	}
	return &ReassignmentStatement{
		Token: keyword,
		Name:  name,
		Value: value,
	}
}

func (statement *ReassignmentStatement) statementNode() {}
func (statement *ReassignmentStatement) TokenLiteral() string {
	return statement.Token.Literal
}
func (statement *ReassignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(statement.Name.String())
	out.WriteString(" = ")
	out.WriteString(statement.Value.String())
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func NewReturn(value Expression) *ReturnStatement {
	return &ReturnStatement{
		Token:       token.RETURN_TOKEN,
		ReturnValue: value,
	}
}

func (statement *ReturnStatement) statementNode() {}
func (statement *ReturnStatement) TokenLiteral() string {
	return statement.Token.Literal
}
func (statement *ReturnStatement) String() string {
	return statement.ReturnValue.String() + ";"
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func NewBlock(statements []Statement) *BlockStatement {
	return &BlockStatement{
		Token:      token.LEFT_BRACE_TOKEN,
		Statements: statements,
	}
}

func (b *BlockStatement) statementNode() {}
func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func NewExpressionStatement(tok token.Token, expr Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Token:      tok,
		Expression: expr,
	}
}

func (statement *ExpressionStatement) statementNode() {}
func (statement *ExpressionStatement) TokenLiteral() string {
	return statement.Token.Literal
}
func (statement *ExpressionStatement) String() string {
	return statement.Expression.String() + ";"
}
