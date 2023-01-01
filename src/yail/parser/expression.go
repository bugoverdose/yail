package parser

import (
	"fmt"
	"strconv"
	"yail/ast/expression"
	"yail/ast/statement"
	"yail/token"
)

const (
	_ int = iota
	NO_PREFERENCE
)

type prefixParseFn func(p *Parser) expression.Expression

var prefixParseFns = map[token.TokenType]prefixParseFn{
	token.IDENTIFIER: parseIdentifier,
	token.INTEGER:    parseIntegerLiteral,
	token.TRUE:       parseBoolean,
	token.FALSE:      parseBoolean,
}

func parseExpressionStatement(p *Parser) *statement.ExpressionStatement {
	stmt := statement.NewExpressionStatement(p.curToken, p.parseExpression(NO_PREFERENCE))
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) expression.Expression {
	prefix := prefixParseFns[p.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	return prefix(p)
}

func parseIdentifier(p *Parser) expression.Expression {
	return expression.NewIdentifier(p.curToken)
}

func parseIntegerLiteral(p *Parser) expression.Expression {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	return expression.NewIntegerLiteral(p.curToken, value)
}

func parseBoolean(p *Parser) expression.Expression {
	return expression.GetPooledBoolean(p.curTokenIs(token.TRUE))
}
