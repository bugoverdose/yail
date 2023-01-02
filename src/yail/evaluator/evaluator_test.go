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
		{"-10", -10},
		{"-0", 0},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 + 2 * 3", 7},
		{"1 + 3 / 2", 2},
		{"1 + 10 % 4", 3},
		{"(1 + 2) * (5 - 2)", 9},
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
		{"!true", false},
		{"!false", true},
		{"10 > 5", true},
		{"10 < 5", false},
		{"5 < 5", false},
		{"5 < 5", false},
		{"5 == 5", true},
		{"5 == true", false},
		{"true == true", true},
		{"true == false", false},
		{"5 != 5", false},
		{"5 != true", true},
		{"true != true", false},
		{"true != false", true},
		{"5 <= 5", true},
		{"5 <= 6", true},
		{"5 >= 5", true},
		{"5 >= 6", false},
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

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if tt.expected != nil {
			testIntegerObject(t, evaluated, int64(tt.expected.(int)))
		} else {
			utils.ValidateObject(evaluated, object.NULL, t)
		}
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
			"val a = 5; a = 10;",
			"can not reassign variables declared with 'val'",
		},
		{
			"5 + true; 10;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5; -true; 10;",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false; 5;",
			"unknown operator: BOOLEAN + BOOLEAN",
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
	utils.ValidateObject(obj, object.NewInteger(expected), t)
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	utils.ValidateObject(obj, &object.Boolean{Value: expected}, t)
}
