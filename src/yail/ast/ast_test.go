package ast

import (
	"testing"
	"strconv"
	"yail/ast/expression"
	"yail/ast/statement"
	"yail/token"
)

func TestAstIntegration(t *testing.T) {
	program := &Program{
		Statements: []statement.Statement{
			statement.NewVariableAssignement(expression.NewIdentifierFrom("x"), newIntegerLiteral("10")),
			statement.NewValueAssignement(expression.NewIdentifierFrom("y"), newIntegerLiteral("20")),
			newExpressionStatement("z"),
		},
	}

	if program.String() != "var x = 10; val y = 20; z;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func newIntegerLiteral(literal string) *expression.IntegerLiteral {
	value, err := strconv.ParseInt(literal, 0, 64)
	if err != nil {
		return nil
	}
	return &expression.IntegerLiteral{
		Token: token.NewInteger(literal),
		Value: value,
	}
}

func newExpressionStatement(identifier string) *statement.ExpressionStatement {
	expr := expression.NewIdentifierFrom(identifier)
	return statement.NewExpressionStatement(expr.Token, expr)
}
