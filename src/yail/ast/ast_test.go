package ast

import (
	"testing"
	"yail/ast/expression"
	"yail/ast/statement"
)

func TestAstIntegration(t *testing.T) {
	program := &Program{
		Statements: []statement.Statement{
			statement.NewVariableAssignement(
				expression.NewIdentifier("x"),
				expression.NewIntegerLiteral("10"),
			),
			statement.NewValueAssignement(
				expression.NewIdentifier("y"),
				expression.NewIntegerLiteral("20"),
			),
		},
	}

	if program.String() != "var x = 10; val y = 20;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
