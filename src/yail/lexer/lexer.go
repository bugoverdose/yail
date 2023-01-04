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
	if lexer.curChar == EOF_CHAR {
		return token.EOF_TOKEN
	}
	if IsLetter(lexer.curChar) {
		return token.NewKeywordOrIdentifier(lexer.readConsecutiveLetters())
	}
	if IsDigit(lexer.curChar) {
		return token.NewInteger(lexer.readNumber())
	}
	if lexer.curChar == '"' {
		return lexer.readString()
	}
	return lexer.toSpecialCharacterToken()
}

func (lexer *Lexer) toSpecialCharacterToken() token.Token {
	if tok, ok := lexer.getTwoCharacterToken(); ok {
		lexer.readNextChar()
		lexer.readNextChar()
		return tok
	}
	tok, ok := token.SingleCharacterTokens[string(lexer.curChar)]
	if !ok {
		tok = token.NewIllegal(string(lexer.curChar))
	}
	lexer.readNextChar()
	return tok
}

func (lexer *Lexer) getTwoCharacterToken() (token.Token, bool) {
	if lexer.nextPosition >= len(lexer.sourceCode) {
		return token.UNUSED_TOKEN, false
	}
	chars := string(lexer.curChar) + string(lexer.sourceCode[lexer.nextPosition])
	tok, ok := token.TwoCharacterTokens[chars]
	if !ok {
		return token.UNUSED_TOKEN, false
	}
	return tok, true
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

func (lexer *Lexer) readString() token.Token {
	startPosition := lexer.curPosition + 1
	for {
		lexer.readNextChar()
		if lexer.curChar == '"' || lexer.curChar == EOF_CHAR {
			lexer.readNextChar()
			break
		}
	}
	return token.NewString(lexer.sourceCode[startPosition : lexer.curPosition-1])
}
