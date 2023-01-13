package evaluator

import (
	"fmt"
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
		testObject(t, actual, tt.expected)
	}
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"Hello World!"`, "Hello World!"},
		{`"Hello" + "World"`, "HelloWorld"},
		{`""`, ""},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		str, _ := evaluated.(*object.String)
		utils.ValidateValue(str.Value, tt.expected, t)
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
		{`"Hello" == "Hello"`, false},
		{`"Hello" != "Hello"`, true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObject(t, evaluated, tt.expected)
	}
}

func TestEvalNullExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"null;"},
		{"val x = null; x;"},
		{"if (false) { 10 }"},
		{"if (false) { 10 }"},
		{"var y = 10; val z = if (true) { y = 15; }; z;"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		utils.ValidateObject(evaluated, object.NULL, t)
	}
}

func TestVariableBindingStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var a = 5; a;", 5},
		{"var a = 5; val b = a; b;", 5},
	}

	for _, tt := range tests {
		actual := testEval(tt.input)
		testObject(t, actual, tt.expected)
	}
}

func TestReassignmentStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var a = 5; a = 10; a;", 10},
		{"var a = 5; val b = 15; a = b; a", 15},
	}

	for _, tt := range tests {
		actual := testEval(tt.input)
		testObject(t, actual, tt.expected)
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{"if (10 > 1) { return 10; }", 10},
		{"if (10 > 1) { if (10 > 1) { return 10; } return 1;}", 10},
		{"val f = func(x) { return x; x + 10; }; f(10);", 10},
		{"val f = func(x) { val result = x + 10; return result; return 10;}; f(10);", 20},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
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
			testObject(t, evaluated, int64(tt.expected.(int)))
		} else {
			utils.ValidateObject(evaluated, object.NULL, t)
		}
	}
}

func TestFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"val identity = func(x) { x; }; identity(5);", 5},
		{"val identity = func(x) { return x; }; identity(5);", 5},
		{"val double = func(x) { x * 2; }; double(5);", 10},
		{"val add = func(x, y) { x + y; }; add(5, 5);", 10},
		{"val add = func(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"val callTwoTimes = func(x, f) { f(f(x)); }; callTwoTimes(2, func(x) { x * x; });", 16},
		{"val callTwoTimes = func(x, f) { f(f(x)); }; callTwoTimes(3, func(x) { x * x; });", 81},
		{"val callTwoTimes = func(x, f) { f(f(x)); }; callTwoTimes(1, func(x) { x + 10; });", 21},
	}
	for _, tt := range tests {
		testObject(t, testEval(tt.input), tt.expected)
	}
}

func TestNestedScopes(t *testing.T) {
	bindings := `
			var i = 5; 
			val useLocalVariableI = func() { return i; }; 
			val useLocalVariableInsideFunction = func() { val i = 15; return i; };
			val returnParameterI = func(i) { return i; }; `

	tests := []struct {
		input    string
		expected int64
	}{
		{"useLocalVariableI();", 5},
		{"i = 30; useLocalVariableI();", 30},
		{"useLocalVariableInsideFunction();", 15},
		{"i = 30; useLocalVariableInsideFunction();", 15},
		{"returnParameterI(10); i;", 5},
		{"i = 30; returnParameterI(10); i;", 30},
		{"returnParameterI(10);", 10},
		{"i = 30; returnParameterI(10);", 10},
	}
	for _, tt := range tests {
		testObject(t, testEval(bindings+tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
			val newAdder = func(x) {
			  func(y) { x + y };
			};
			val addTwo = newAdder(2);
			addTwo(5);`
	testObject(t, testEval(input), 7)
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(result.Elements), 3, t)
	testObject(t, result.Elements[0], 1)
	testObject(t, result.Elements[1], 4)
	testObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[1, 2, 3][0]", 1},
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][2]", 3},
		{"val i = 0; [1][i];", 1},
		{"[1, 2, 3][1 + 1];", 3},
		{"val arr = [1, 2, 3]; arr[2];", 3},
		{"val arr = [1, 2, 3]; arr[0] + arr[1] + arr[2];", 6},
		{"val arr = [1, 2, 3]; val i = arr[0]; arr[i]", 2},
		{"[1, 2, 3][3]", nil},
		{"[1, 2, 3][-1]", nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObject(t, evaluated, tt.expected)
	}
}
func TestHashMapLiterals(t *testing.T) {
	input := `val two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`
	expected := map[object.HashKey]int64{
		object.NewString("one").HashKey():   1,
		object.NewString("two").HashKey():   2,
		object.NewString("three").HashKey(): 3,
		object.NewInteger(4).HashKey():      4,
		object.TRUE.HashKey():               5,
		object.FALSE.HashKey():              6,
	}

	evaluated := testEval(input)
	result, ok := evaluated.(*object.HashMap)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(len(result.Pairs), len(expected), t)

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		utils.ValidateValue(ok, true, t)
		testObject(t, pair.Value, expectedValue)
	}
}

func TestHashMapAccessExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`{"foo": 5}["bar"]`, nil},
		{`{}["foo"]`, nil},
		{`{"foo": 5}["foo"]`, 5},
		{`val key = "foo"; {"foo": 5}[key]`, 5},
		{`{100: { "a": 1, "b": 2 }}[100]["a"]`, 1},
		{`{true: [10, 20, 30]}[true][0]`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObject(t, evaluated, tt.expected)
	}
}

func TestEvalBuiltinFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("Hello World")`, 11},
		{`len([])`, 0},
		{`len(["a", "b"])`, 2},
		{`head([1, 2, 3])`, 1},
		{`head([])`, nil},
		{`tail([1, 2, 3])`, 3},
		{`tail([])`, nil},
		{`val arr = [1, 2, 3]; push(arr, 10); arr;`, []int64{1, 2, 3, 10}},
		{`val arr = [1, 2, 3]; pushleft(arr, 10); arr;`, []int64{10, 1, 2, 3}},
		{`val arr = [1, 2, 3]; pop(arr); arr;`, []int64{1, 2}},
		{`val arr = [1, 2, 3]; popleft(arr); arr;`, []int64{2, 3}},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObject(t, evaluated, tt.expected)
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
			"fn(10, 20)",
			"identifier not found: fn",
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
		{
			"!5; 10;",
			"unknown operator: !INTEGER",
		},
		{
			"var i = 5; val reassignFunc = func() { i = 10; }; reassignFunc();",
			"identifier not found: 'i'",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
		{
			`[1, 2, 3]["wrong_index_format"];`,
			"unsupported operation: ARRAY[STRING]",
		},
		{
			`{"key": "abc"}[func(x) {x}];`,
			"unusable as hash key: FUNCTION",
		},
		{
			`len(1)`,
			"len(INTEGER) not supported",
		},
		{
			`len("one", "two")`,
			"wrong number of arguments: expected 1, but received 2",
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

func testObject(
	t *testing.T,
	actual object.Object,
	expected interface{},
) {
	if v, ok := actual.(*object.Error); ok {
		fmt.Printf("%+v", v.Message)
	}
	switch v := expected.(type) {
	case int:
		utils.ValidateObject(actual, object.NewInteger(int64(v)), t)
	case int64:
		utils.ValidateObject(actual, object.NewInteger(v), t)
	case bool:
		utils.ValidateObject(actual, object.GetPooledBooleanObject(v), t)
	case []int64:
		actual := actual.(*object.Array).Elements
		utils.ValidateValue(len(actual), len(v), t)
		for idx, val := range actual {
			utils.ValidateValue(val.(*object.Integer).Value, v[idx], t)
		}
	case nil:
		utils.ValidateObject(actual, object.NULL, t)
	default:
		t.Errorf("Failed to handle %T.", v)
	}
}
