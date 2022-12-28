package lexer

import "yail/token"

const EOF_CHAR = 0

type Lexer struct {
	sourceCode   string
	curPosition  int
	nextPosition int
	curChar      byte
}

func New(sourceCode string) *Lexer {
	lexer := &Lexer{sourceCode: sourceCode}
	lexer.readNextChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	lexer.eatWhitespace()
	if IsLetter(lexer.curChar) {
		return token.NewKeywordOrIdentifier(lexer.readConsecutiveLetters())
	}
	if IsDigit(lexer.curChar) {
		return token.NewInteger(lexer.readNumber())
	}
	return lexer.toSingleCharacterToken()
}

func (lexer *Lexer) toSingleCharacterToken() token.Token {
	var curToken token.Token
	switch lexer.curChar {
	case '=':
		curToken = token.New(token.ASSIGN, lexer.curChar)
	case ';':
		curToken = token.New(token.SEMICOLON, lexer.curChar)
	case EOF_CHAR:
		curToken.Literal = ""
		curToken.Type = token.EOF
	default:
		curToken = token.New(token.ILLEGAL, lexer.curChar)
	}
	lexer.readNextChar()
	return curToken
}

func (lexer *Lexer) readNextChar() {
	if lexer.nextPosition >= len(lexer.sourceCode) {
		lexer.curChar = EOF_CHAR
	} else {
		lexer.curChar = lexer.sourceCode[lexer.nextPosition]
	}
	lexer.curPosition = lexer.nextPosition
	lexer.nextPosition += 1
}

func (lexer *Lexer) peekNextChar() byte {
	if lexer.nextPosition >= len(lexer.sourceCode) {
		return EOF_CHAR
	}
	return lexer.sourceCode[lexer.nextPosition]
}

func (lexer *Lexer) eatWhitespace() {
	for lexer.curChar == ' ' || lexer.curChar == '\t' || lexer.curChar == '\n' || lexer.curChar == '\r' {
		lexer.readNextChar()
	}
}

func (lexer *Lexer) readConsecutiveLetters() string {
	curPosition := lexer.curPosition
	for IsLetter(lexer.curChar) || IsDigit(lexer.curChar) {
		lexer.readNextChar()
	}
	return lexer.sourceCode[curPosition:lexer.curPosition]
}

func (lexer *Lexer) readNumber() string {
	curPosition := lexer.curPosition
	for IsDigit(lexer.curChar) {
		lexer.readNextChar()
	}
	return lexer.sourceCode[curPosition:lexer.curPosition]
}
