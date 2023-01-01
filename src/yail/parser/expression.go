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
	PREFIX_PREFERENCE
)

type nullDenotation func(p *Parser) expression.Expression

func (p *Parser) initNullDenotations() {
	p.nuds = map[token.TokenType]nullDenotation{
		token.IDENTIFIER: parseIdentifier,
		token.INTEGER:    parseIntegerLiteral,
		token.TRUE:       parseBoolean,
		token.FALSE:      parseBoolean,
		token.NOT:        parsePrefixExpression,
		token.MINUS:      parsePrefixExpression,
	}
}

func parseExpressionStatement(p *Parser) *statement.ExpressionStatement {
	stmt := statement.NewExpressionStatement(p.curToken, p.parseExpression(NO_PREFERENCE))
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) expression.Expression {
	nud := p.nuds[p.curToken.Type]
	if nud == nil {
		msg := fmt.Sprintf("no parse function for %s found", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	return nud(p)
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

func parsePrefixExpression(p *Parser) expression.Expression {
	prefixToken := p.curToken
	p.nextToken()
	rightNode := p.parseExpression(PREFIX_PREFERENCE)
	return expression.NewPrefix(prefixToken, rightNode)
}
