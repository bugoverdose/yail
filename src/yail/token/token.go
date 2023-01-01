package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "END_OF_FILE"

	IDENTIFIER = "IDENTIFIER"      // x, y, ...
	INTEGER    = "INTEGER_LITERAL" // 1, 2, 10, ...

	// Operators
	ASSIGN = "="
	MINUS  = "-"
	NOT    = "!"

	// Delimiters
	SEMICOLON = ";"

	// Keywords
	VAR   = "var"
	VAL   = "val"
	TRUE  = "true"
	FALSE = "false"
)

type Token struct {
	Type    TokenType
	Literal string
}

// TODO: implement hash set if hash map is not needed
var keywords = map[string]TokenType{
	VAR:   VAR,
	VAL:   VAL,
	TRUE:  TRUE,
	FALSE: FALSE,
	NOT:   NOT,
	MINUS: MINUS,
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
