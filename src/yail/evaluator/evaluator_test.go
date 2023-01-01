package evaluator

import (
	"testing"
	"yail/environment"
	"yail/lexer"
	"yail/object"
	"yail/parser"
	"yail/utils"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"15", 15},
	}

	for _, tt := range tests {
		actual := testEval(tt.input)
		testIntegerObject(t, actual, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestVariableBindingStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var a = 5; a;", 5},
		{"var a = 5; val b = a; b;", 5},
	}

	for _, tt := range tests {
		actual := testEval(tt.input)
		testIntegerObject(t, actual, tt.expected)
	}
}

func TestReassignmentStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var a = 5; a = 10; a;", 10},
		{"var a = 5; val b = 15; a = b; a", 15},
	}

	for _, tt := range tests {
		actual := testEval(tt.input)
		testIntegerObject(t, actual, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"x",
			"identifier not found: x",
		},
		{
			"val a = 5; a = b;",
			"can not reassign variables declared with 'val'",
		},
	}

	for _, tt := range tests {
		actual := testEval(tt.input)
		expected := &object.Error{Message: tt.expectedMessage}
		utils.ValidateObject(actual, expected, t)
	}
}

func testEval(input string) object.Object {
	p := parser.New(lexer.New(input))
	program := p.ParseProgram()
	env := environment.NewEnvironment()
	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	utils.ValidateObject(obj, &object.Integer{Value: expected}, t)
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	utils.ValidateObject(obj, &object.Boolean{Value: expected}, t)
}
