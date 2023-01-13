package parser

import (
	"fmt"
	"testing"
	"yail/ast"
	"yail/lexer"
	"yail/token"
	"yail/utils"
)

func TestVariableBindingStatement(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"var x = 5;", "x", 5},
		{"val y = true;", "y", true},
		{"val a = b;", "a", "b"},
		{"val _ = null;", "_", nil},
	}

	for _, tt := range tests {
		program := parseAndValidate(t, tt.input)
		utils.ValidateValue(len(program.Statements), 1, t)
		stmt := program.Statements[0]
		testVariableBindingStatement(t, stmt, tt.expectedIdentifier)
		actualValue := stmt.(*ast.VariableBindingStatement).Value
		testLiteralExpression(t, actualValue, tt.expectedValue)
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return;", nil},
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
	}

	for _, tt := range tests {
		program := parseAndValidate(t, tt.input)
		utils.ValidateValue(len(program.Statements), 1, t)
		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		utils.ValidateValue(ok, true, t)
		utils.ValidateValue(returnStmt.TokenLiteral(), token.RETURN, t)
		testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue)
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		program := parseAndValidate(t, tt.input)
		utils.ValidateValue(len(program.Statements), 1, t)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		utils.ValidateValue(ok, true, t)
		boolean, ok := stmt.Expression.(*ast.BooleanExpression)
		utils.ValidateValue(ok, true, t)
		utils.ValidateValue(boolean.Value, tt.expectedBoolean, t)
	}
}

func TestStringExpression(t *testing.T) {
	input := `"hello world";`
	program := parseAndValidate(t, input)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(len(program.Statements), 1, t)
	literal, ok := stmt.Expression.(*ast.StringLiteralExpression)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(literal.Value, "hello world", t)
}

func TestPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
		{"!foobar", "!", "foobar"},
		{"-foobar;", "-", "foobar"},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		program := parseAndValidate(t, tt.input)
		utils.ValidateValue(len(program.Statements), 1, t)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		utils.ValidateValue(ok, true, t)
		expr, ok := stmt.Expression.(*ast.PrefixExpression)
		utils.ValidateValue(ok, true, t)
		utils.ValidateValue(expr.Operator, tt.operator, t)
		testLiteralExpression(t, expr.RightNode, tt.value)
	}
}

func TestInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + x;", 5, "+", "x"},
		{"5 - x;", 5, "-", "x"},
		{"5 * x;", 5, "*", "x"},
		{"5 / x;", 5, "/", "x"},
		{"5 > x;", 5, ">", "x"},
		{"5 < x;", 5, "<", "x"},
		{"5 == x;", 5, "==", "x"},
		{"5 != x;", 5, "!=", "x"},
		{"5 <= x;", 5, "<=", "x"},
		{"5 >= x;", 5, ">=", "x"},
	}

	for _, tt := range infixTests {
		program := parseAndValidate(t, tt.input)
		utils.ValidateValue(len(program.Statements), 1, t)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		utils.ValidateValue(ok, true, t)
		testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`
	program := parseAndValidate(t, input)

	utils.ValidateValue(len(program.Statements), 1, t)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)

	expr, ok := stmt.Expression.(*ast.IfExpression)
	utils.ValidateValue(ok, true, t)
	testInfixExpression(t, expr.Condition, "x", "<", "y")

	utils.ValidateValue(len(expr.Consequence.Statements), 1, t)
	consequence, ok := expr.Consequence.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)
	testLiteralExpression(t, consequence.Expression, "x")

	utils.ValidateValue(expr.Alternative, nil, t)
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { return x; } else { y }`
	program := parseAndValidate(t, input)

	utils.ValidateValue(len(program.Statements), 1, t)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)

	expr, ok := stmt.Expression.(*ast.IfExpression)
	utils.ValidateValue(ok, true, t)
	testInfixExpression(t, expr.Condition, "x", "<", "y")

	utils.ValidateValue(len(expr.Consequence.Statements), 1, t)
	consequence, ok := expr.Consequence.Statements[0].(*ast.ReturnStatement)
	utils.ValidateValue(ok, true, t)
	testLiteralExpression(t, consequence.ReturnValue, "x")

	utils.ValidateValue(len(expr.Alternative.Statements), 1, t)
	alternative, ok := expr.Alternative.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)
	testLiteralExpression(t, alternative.Expression, "y")
}

func TestFunctionLiteral(t *testing.T) {
	input := `func(x, y) { x + y; }`
	program := parseAndValidate(t, input)

	utils.ValidateValue(len(program.Statements), 1, t)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	utils.ValidateValue(ok, true, t)

	utils.ValidateValue(len(function.Parameters), 2, t)
	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	utils.ValidateValue(len(function.Body.Statements), 1, t)
	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)
	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameters(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "func() {};", expectedParams: []string{}},
		{input: "func(x) {};", expectedParams: []string{"x"}},
		{input: "func(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		program := parseAndValidate(t, tt.input)
		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)
		utils.ValidateValue(len(function.Parameters), len(tt.expectedParams), t)
		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestEmptyFunctionCallExpression(t *testing.T) {
	input := "add();"

	program := parseAndValidate(t, input)
	utils.ValidateValue(len(program.Statements), 1, t)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(exp.Function.String(), "add", t)
	utils.ValidateValue(len(exp.Arguments), 0, t)
}

func TestFunctionCallExpression(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	program := parseAndValidate(t, input)
	utils.ValidateValue(len(program.Statements), 1, t)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	utils.ValidateValue(ok, true, t)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(exp.Function.String(), "add", t)
	utils.ValidateValue(len(exp.Arguments), 3, t)
	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestParsingEmptyArrayLiterals(t *testing.T) {
	input := "[];"
	program := parseAndValidate(t, input)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(array.Elements), 0, t)
}

func TestParsingArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3];"
	program := parseAndValidate(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(array.Elements), 3, t)
	testLiteralExpression(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestParsingAccessingArrayByIndexExpressions(t *testing.T) {
	input := "arr[1 + 1]"
	program := parseAndValidate(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.CollectionAccessExpression)
	utils.ValidateValue(ok, true, t)

	testLiteralExpression(t, indexExp.Left, "arr")
	testInfixExpression(t, indexExp.Index, 1, "+", 1)
}

func TestParsingEmptyHashLiteral(t *testing.T) {
	input := "{}"
	program := parseAndValidate(t, input)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.HashMapLiteral)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(hash.Pairs), 0, t)
}

func TestParsingHashLiteralKeys(t *testing.T) {
	input := `{"one": 1, "two": 2, true: 10, false: 20, 1: 100, 2: 200}`
	expected := map[interface{}]int64{
		"one": 1,
		"two": 2,
		true:  10,
		false: 20,
		1:     100,
		2:     200,
	}
	program := parseAndValidate(t, input)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.HashMapLiteral)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(hash.Pairs), len(expected), t)
	for key, value := range hash.Pairs {
		switch k := key.(type) {
		case *ast.StringLiteralExpression:
			testLiteralExpression(t, value, expected[k.Value])
		case *ast.BooleanExpression:
			testLiteralExpression(t, value, expected[k.Value])
		case *ast.IntegerLiteralExpression:
			testLiteralExpression(t, value, expected[int(k.Value)])
		}
	}
}

func TestParsingHashLiteralValues(t *testing.T) {
	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`
	hashTests := map[string]struct {
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		"one":   {0, "+", 1},
		"two":   {10, "-", 8},
		"three": {15, "/", 5},
	}
	program := parseAndValidate(t, input)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.HashMapLiteral)

	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(hash.Pairs), len(hashTests), t)
	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.StringLiteralExpression)
		utils.ValidateValue(ok, true, t)
		tt := hashTests[literal.Value]
		testInfixExpression(t, value, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestOperationPriorities(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b);",
		},
		{
			"!-a",
			"(!(-a));",
		},
		{
			"a + b + c",
			"((a + b) + c);",
		},
		{
			"a + (b - c)",
			"(a + (b - c));",
		},
		{
			"a * b * c",
			"((a * b) * c);",
		},
		{
			"a * (b / c)",
			"(a * (b / c));",
		},
		{
			"a + b / c",
			"(a + (b / c));",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f);",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4));",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4));",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false);",
		},
		{
			"3 < 5 != true",
			"((3 < 5) != true);",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4);",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2);",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5));",
		},
		{
			"(5 + 5) * 2 * (5 + 5)",
			"(((5 + 5) * 2) * (5 + 5));",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5));",
		},
		{
			"!(true == true)",
			"(!(true == true));",
		},
		{
			"!true == true",
			"((!true) == true);",
		},
		{
			"a + add(b, c) * d",
			"(a + (add(b, c) * d));",
		},
		{
			"add(-func(x, y) { x - y; }, z)",
			"add((-func(x, y) { (x - y); }), z);",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)));",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g));",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d);",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])));",
		},
	}

	for _, tt := range tests {
		program := parseAndValidate(t, tt.input)
		actual := program.String()
		utils.ValidateValue(actual, tt.expected, t)
	}
}

func TestIllegalInput(t *testing.T) {
	tests := []struct {
		input   string
		illegal string
	}{
		{"&;", "&"},
		{"5^2", "^"},
		{"a + #", "#"},
		{"2@", "@"},
		{"$1", "$"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		p.ParseProgram()
		errors := p.Errors()
		for _, actual := range errors {
			expected := fmt.Sprintf("failed to understand: '%s'", tt.illegal)
			utils.ValidateValue(actual, expected, t)
		}
	}
}

func validateNoParserErrors(t *testing.T, p *Parser) {
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

func parseAndValidate(t *testing.T, input string) *ast.Program {
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	validateNoParserErrors(t, p)
	return program
}

func testVariableBindingStatement(t *testing.T, s ast.Statement, identifier string) {
	utils.ValidateMatchAnyValue(s.TokenLiteral(), []string{token.VAR, token.VAL}, t)
	stmt, ok := s.(*ast.VariableBindingStatement)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(stmt.Name.Value, identifier, t)
	utils.ValidateValue(stmt.Name.TokenLiteral(), identifier, t)
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
	case int64:
		testIntegerLiteral(t, exp, v)
	case bool:
		testBooleanLiteral(t, exp, v)
	case nil:
		utils.ValidateValue(exp.(*ast.NullExpression), ast.NULL, t)
	case string:
		testIdentifier(t, exp, v)
	default:
		t.Errorf("Failed to handle %T.", exp)
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) {
	integer, ok := il.(*ast.IntegerLiteralExpression)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(integer.Value, value, t)
	utils.ValidateValue(integer.TokenLiteral(), fmt.Sprintf("%d", value), t)
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) {
	boolean, ok := exp.(*ast.BooleanExpression)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(boolean.Value, value, t)
	utils.ValidateValue(boolean.TokenLiteral(), fmt.Sprintf("%t", value), t)
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) {
	ident, ok := exp.(*ast.IdentifierExpression)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(ident.Value, value, t)
	utils.ValidateValue(ident.TokenLiteral(), value, t)
}

func testInfixExpression(
	t *testing.T,
	expr ast.Expression,
	expectedLeftValue interface{},
	expectedOperator string,
	expectedRightValue interface{},
) {
	infixExpr, ok := expr.(*ast.InfixExpression)
	utils.ValidateValue(ok, true, t)
	testLiteralExpression(t, infixExpr.LeftNode, expectedLeftValue)
	utils.ValidateValue(infixExpr.Operator, expectedOperator, t)
	testLiteralExpression(t, infixExpr.RightNode, expectedRightValue)
}
