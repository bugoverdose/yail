package lexer

import (
	"testing"
	"yail/token"
	"yail/utils"
)

func TestVariableBinding(t *testing.T) {
	input := `var five = 5;
    	      val a = b;
              val x = false;`
	lexer := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},

		{token.VAL, "val"},
		{token.IDENTIFIER, "a"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "b"},
		{token.SEMICOLON, ";"},

		{token.VAL, "val"},
		{token.IDENTIFIER, "x"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	for _, tt := range tests {
		tok := lexer.NextToken()
		utils.ValidateValue(tok.Type, tt.expectedType, t)
		utils.ValidateValue(tok.Literal, tt.expectedLiteral, t)
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x > y) { x } else { y };`
	lexer := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.GREATER_THAN, ">"},
		{token.IDENTIFIER, "y"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACKET, "{"},
		{token.IDENTIFIER, "x"},
		{token.RIGHT_BRACKET, "}"},
		{token.ELSE, "else"},
		{token.LEFT_BRACKET, "{"},
		{token.IDENTIFIER, "y"},
		{token.RIGHT_BRACKET, "}"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	for _, tt := range tests {
		tok := lexer.NextToken()
		utils.ValidateValue(tok.Type, tt.expectedType, t)
		utils.ValidateValue(tok.Literal, tt.expectedLiteral, t)
	}
}

func TestSingleCharacterToken(t *testing.T) {
	input := `!true;
              1 + (2 - 3) * 10 / 2 % 3;
			  -10 < 5;
			  5 > 10;`
	lexer := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NOT, "!"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},

		{token.INTEGER, "1"},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESIS, "("},
		{token.INTEGER, "2"},
		{token.MINUS, "-"},
		{token.INTEGER, "3"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.MULTIPLY, "*"},
		{token.INTEGER, "10"},
		{token.DIVIDE, "/"},
		{token.INTEGER, "2"},
		{token.MODULO, "%"},
		{token.INTEGER, "3"},
		{token.SEMICOLON, ";"},

		{token.MINUS, "-"},
		{token.INTEGER, "10"},
		{token.LESS_THAN, "<"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},

		{token.INTEGER, "5"},
		{token.GREATER_THAN, ">"},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	for _, tt := range tests {
		tok := lexer.NextToken()
		utils.ValidateValue(tok.Type, tt.expectedType, t)
		utils.ValidateValue(tok.Literal, tt.expectedLiteral, t)
	}
}

func TestTwoCharacterToken(t *testing.T) {
	input := `5 == 3;
              5 != 3;
			  5 <= 3;
			  5 >= 3;`
	lexer := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INTEGER, "5"},
		{token.EQUAL, "=="},
		{token.INTEGER, "3"},
		{token.SEMICOLON, ";"},

		{token.INTEGER, "5"},
		{token.NOT_EQUAL, "!="},
		{token.INTEGER, "3"},
		{token.SEMICOLON, ";"},

		{token.INTEGER, "5"},
		{token.LESS_OR_EQUAL, "<="},
		{token.INTEGER, "3"},
		{token.SEMICOLON, ";"},

		{token.INTEGER, "5"},
		{token.GREATER_OR_EQUAL, ">="},
		{token.INTEGER, "3"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	for _, tt := range tests {
		tok := lexer.NextToken()
		utils.ValidateValue(tok.Type, tt.expectedType, t)
		utils.ValidateValue(tok.Literal, tt.expectedLiteral, t)
	}
}

func TestIllegalToken(t *testing.T) {
	input := `&;
    	      var x = a^b;`
	lexer := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ILLEGAL, "&"},
		{token.SEMICOLON, ";"},

		{token.VAR, "var"},
		{token.IDENTIFIER, "x"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "a"},
		{token.ILLEGAL, "^"},
		{token.IDENTIFIER, "b"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	for _, tt := range tests {
		tok := lexer.NextToken()
		utils.ValidateValue(tok.Type, tt.expectedType, t)
		utils.ValidateValue(tok.Literal, tt.expectedLiteral, t)
	}
}
