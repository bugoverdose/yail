package lexer

func IsLetter(curChar byte) bool {
	return ('a' <= curChar && curChar <= 'z') || ('A' <= curChar && curChar <= 'Z') || curChar == '_'
}

func IsDigit(curChar byte) bool {
	return '0' <= curChar && curChar <= '9'
}
