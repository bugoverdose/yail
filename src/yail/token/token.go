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
	SEMICOLON         = ";"
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACKET      = "{"
	RIGHT_BRACKET     = "}"

	// Keywords
	VAR    = "var"
	VAL    = "val"
	TRUE   = "true"
	FALSE  = "false"
	IF     = "if"
	ELSE   = "else"
	RETURN = "return"
)

type Token struct {
	Type    TokenType
	Literal string
}

var (
	EOF_TOKEN          = Token{Type: EOF, Literal: ""}
	UNUSED_TOKEN       = New(ILLEGAL)
	IF_TOKEN           = New(IF)
	LEFT_BRACKET_TOKEN = New(LEFT_BRACKET)
)

var keywords = map[string]Token{
	VAR:    New(VAR),
	VAL:    New(VAL),
	TRUE:   New(TRUE),
	FALSE:  New(FALSE),
	IF:     IF_TOKEN,
	ELSE:   New(ELSE),
	RETURN: New(RETURN),
}

var SingleCharacterTokens = map[string]Token{
	ASSIGN:            New(ASSIGN),
	NOT:               New(NOT),
	PLUS:              New(PLUS),
	MINUS:             New(MINUS),
	MULTIPLY:          New(MULTIPLY),
	DIVIDE:            New(DIVIDE),
	MODULO:            New(MODULO),
	LESS_THAN:         New(LESS_THAN),
	GREATER_THAN:      New(GREATER_THAN),
	SEMICOLON:         New(SEMICOLON),
	LEFT_PARENTHESIS:  New(LEFT_PARENTHESIS),
	RIGHT_PARENTHESIS: New(RIGHT_PARENTHESIS),
	LEFT_BRACKET:      LEFT_BRACKET_TOKEN,
	RIGHT_BRACKET:     New(RIGHT_BRACKET),
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

func NewIllegal(literal string) Token {
	return Token{Type: ILLEGAL, Literal: literal}
}
