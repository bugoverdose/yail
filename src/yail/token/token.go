package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "END_OF_FILE"

	IDENTIFIER = "IDENTIFIER"      // x, y, ...
	INTEGER    = "INTEGER_LITERAL" // 1, 2, 10, ...

	// Operators
	ASSIGN           = "="
	NOT              = "!"
	PLUS             = "+"
	MINUS            = "-"
	MULTIPLY         = "*"
	DIVIDE           = "/"
	MODULO           = "%"
	LESS_THAN        = "<"
	GREATER_THAN     = ">"
	EQUAL            = "=="
	NOT_EQUAL        = "!="
	LESS_OR_EQUAL    = "<="
	GREATER_OR_EQUAL = ">="

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

var (
	ILLEGAL_TOKEN = New(ILLEGAL)
	EOF_TOKEN     = Token{Type: EOF, Literal: ""}
)

var keywords = map[string]Token{
	VAR:   New(VAR),
	VAL:   New(VAL),
	TRUE:  New(TRUE),
	FALSE: New(FALSE),
}

var SingleCharacterTokens = map[string]Token{
	ASSIGN:       New(ASSIGN),
	NOT:          New(NOT),
	PLUS:         New(PLUS),
	MINUS:        New(MINUS),
	MULTIPLY:     New(MULTIPLY),
	DIVIDE:       New(DIVIDE),
	MODULO:       New(MODULO),
	LESS_THAN:    New(LESS_THAN),
	GREATER_THAN: New(GREATER_THAN),
	SEMICOLON:    New(SEMICOLON),
}

var TwoCharacterTokens = map[string]Token{
	EQUAL:            New(EQUAL),
	NOT_EQUAL:        New(NOT_EQUAL),
	LESS_OR_EQUAL:    New(LESS_OR_EQUAL),
	GREATER_OR_EQUAL: New(GREATER_OR_EQUAL),
}

func New(tokenType TokenType) Token {
	return Token{Type: tokenType, Literal: string(tokenType)}
}

func NewInteger(literal string) Token {
	return Token{Type: INTEGER, Literal: literal}
}

func NewKeywordOrIdentifier(literal string) Token {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return NewIdentifier(literal)
}

func NewIdentifier(literal string) Token {
	return Token{Type: IDENTIFIER, Literal: literal}
}
