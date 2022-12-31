package lexer

import (
	"testing"
	"yail/token"
	"yail/utils"
)

func TestVariableBinding(t *testing.T) {
	input := `var five = 5;
    	      val ten = 10;`
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
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
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
