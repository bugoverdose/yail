package parser

import (
	"fmt"
	"strconv"
	"yail/ast/expression"
	"yail/ast/statement"
	"yail/token"
)

const (
	NO_PRIORITY int = iota
	EQUALS_PRIORITY
	COMPARISON_PRIORITY
	SUM_SUBTRACT_PRIORITY
	PROD_DIV_PRIORITY
	PREFIX_PRIORITY
)

var priorities = map[token.TokenType]int{
	token.PLUS:             SUM_SUBTRACT_PRIORITY,
	token.MINUS:            SUM_SUBTRACT_PRIORITY,
	token.MULTIPLY:         PROD_DIV_PRIORITY,
	token.DIVIDE:           PROD_DIV_PRIORITY,
	token.MODULO:           PROD_DIV_PRIORITY,
	token.LESS_THAN:        COMPARISON_PRIORITY,
	token.GREATER_THAN:     COMPARISON_PRIORITY,
	token.EQUAL:            EQUALS_PRIORITY,
	token.NOT_EQUAL:        EQUALS_PRIORITY,
	token.LESS_OR_EQUAL:    EQUALS_PRIORITY,
	token.GREATER_OR_EQUAL: EQUALS_PRIORITY,
}

type (
	nullDenotation func(p *Parser) expression.Expression
	leftDenotation func(e expression.Expression, p *Parser) expression.Expression
)

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

func (p *Parser) initLeftDenotations() {
	p.leds = map[token.TokenType]leftDenotation{
		token.PLUS:             parseInfixExpression,
		token.MINUS:            parseInfixExpression,
		token.MULTIPLY:         parseInfixExpression,
		token.DIVIDE:           parseInfixExpression,
		token.MODULO:           parseInfixExpression,
		token.LESS_THAN:        parseInfixExpression,
		token.GREATER_THAN:     parseInfixExpression,
		token.EQUAL:            parseInfixExpression,
		token.NOT_EQUAL:        parseInfixExpression,
		token.LESS_OR_EQUAL:    parseInfixExpression,
		token.GREATER_OR_EQUAL: parseInfixExpression,
	}
}

func parseExpressionStatement(p *Parser) *statement.ExpressionStatement {
	stmt := statement.NewExpressionStatement(p.curToken, p.parseExpression(NO_PRIORITY))
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(priority int) expression.Expression {
	ok, leftExp := p.parseCurToken()
	if !ok {
		return nil
	}
	return p.pratParse(leftExp, priority)
}

func (p *Parser) parseCurToken() (bool, expression.Expression) {
	nud := p.nuds[p.curToken.Type]
	if nud == nil {
		msg := fmt.Sprintf("failed to understand: '%s'", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return false, nil
	}
	return true, nud(p)
}

func (p *Parser) pratParse(leftExp expression.Expression, priority int) expression.Expression {
	for !p.peekTokenIs(token.SEMICOLON) && priority < p.getNextTokenPriority() {
		infix := p.leds[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp, p)
	}
	return leftExp
}

func (p *Parser) getCurTokenPriority() int {
	if p, ok := priorities[p.curToken.Type]; ok {
		return p
	}
	return NO_PRIORITY
}

func (p *Parser) getNextTokenPriority() int {
	if p, ok := priorities[p.peekToken.Type]; ok {
		return p
	}
	return NO_PRIORITY
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
	rightNode := p.parseExpression(PREFIX_PRIORITY)
	return expression.NewPrefix(prefixToken, rightNode)
}

func parseInfixExpression(leftNode expression.Expression, p *Parser) expression.Expression {
	infixToken := p.curToken
	priority := p.getCurTokenPriority()
	p.nextToken()
	rightNode := p.parseExpression(priority)
	return expression.NewInfix(leftNode, infixToken, rightNode)
}
