package parser

import (
	"fmt"
	"testing"
	"yail/ast/expression"
	"yail/ast/statement"
	"yail/lexer"
	"yail/token"
	"yail/utils"
)

func TestVariableBindingStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"var x = 5;", "x", 5},
		{"val a = b;", "a", "b"},
		{"val _ = 10", "_", 10},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		utils.ValidateValue(len(program.Statements), 1, t)
		stmt := program.Statements[0]
		testVariableBindingStatement(t, stmt, tt.expectedIdentifier)
		actualValue := stmt.(*statement.VariableBinding).Value
		testLiteralExpression(t, actualValue, tt.expectedValue)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testVariableBindingStatement(t *testing.T, s statement.Statement, identifier string) {
	utils.ValidateMatchAnyValue(s.TokenLiteral(), []string{token.VAR, token.VAL}, t)
	stmt, ok := s.(*statement.VariableBinding)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(stmt.Name.Value, identifier, t)
	utils.ValidateValue(stmt.Name.TokenLiteral(), identifier, t)
}

func testLiteralExpression(
	t *testing.T,
	exp expression.Expression,
	expected interface{},
) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
	case int64:
		testIntegerLiteral(t, exp, v)
	case string:
		testIdentifier(t, exp, v)
	default:
		t.Errorf("Failed to handle %T.", exp)
	}
}

func testIntegerLiteral(t *testing.T, il expression.Expression, value int64) {
	integer, ok := il.(*expression.IntegerLiteral)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(integer.Value, value, t)
	utils.ValidateValue(integer.TokenLiteral(), fmt.Sprintf("%d", value), t)
}

func testIdentifier(t *testing.T, exp expression.Expression, value string) {
	ident, ok := exp.(*expression.Identifier)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(ident.Value, value, t)
	utils.ValidateValue(ident.TokenLiteral(), value, t)
}
