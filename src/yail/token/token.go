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
	VAR = "VARIABLE_ASSIGNMENT"
	VAL = "VALUE_ASSIGNMENT"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"var": VAR,
	"val": VAL,
}

func New(tokenType TokenType, curChar byte) Token {
	return Token{Type: tokenType, Literal: string(curChar)}
}

func NewInteger(literal string) Token {
	return Token{Type: INTEGER, Literal: literal}
}

func NewKeywordOrIdentifier(literal string) Token {
	if keyword_type, ok := keywords[literal]; ok {
		return Token{Type: keyword_type, Literal: literal}
	}
	return Token{Type: IDENTIFIER, Literal: literal}
}
