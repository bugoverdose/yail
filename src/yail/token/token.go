package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "END_OF_FILE"

	IDENTIFIER = "IDENTIFIER"      // x, y, ...
	INTEGER    = "INTEGER_LITERAL" // 1, 2, 10, ...

	// Operators
	ASSIGN = "="

	// Delimiters
	SEMICOLON = ";"

	// Keywords
	VAR = "var"
	VAL = "val"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	VAR: VAR,
	VAL: VAL,
}

func New(tokenType TokenType, curChar byte) Token {
	return Token{Type: tokenType, Literal: string(curChar)}
}

func NewInteger(literal string) Token {
	return Token{Type: INTEGER, Literal: literal}
}

func NewKeywordOrIdentifier(literal string) Token {
	if _, ok := keywords[literal]; ok {
		return NewKeyword(literal)
	}
	return NewIdentifier(literal)
}

func NewKeyword(literal string) Token {
	return Token{Type: keywords[literal], Literal: literal}
}

func NewIdentifier(literal string) Token {
	return Token{Type: IDENTIFIER, Literal: literal}
}
