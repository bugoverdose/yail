package parser

import (
	"fmt"
	"yail/ast"
	"yail/token"
)

const (
	NO_PRIORITY int = iota
	EQUALS_PRIORITY
	COMPARISON_PRIORITY
	SUM_SUBTRACT_PRIORITY
	PROD_DIV_PRIORITY
	PREFIX_PRIORITY
	FUNCTION_CALL_PRIORITY
	COLLECTION_ACCESS_PRIORITY
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
	token.LEFT_PARENTHESIS: FUNCTION_CALL_PRIORITY,
	token.LEFT_BRACKET:     COLLECTION_ACCESS_PRIORITY,
}

type (
	nullDenotation func(p *Parser) ast.Expression
	leftDenotation func(e ast.Expression, p *Parser) ast.Expression
)

func parseExpressionStatement(p *Parser) *ast.ExpressionStatement {
	stmt := ast.NewExpressionStatement(p.curToken, p.parseExpression(NO_PRIORITY))
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(priority int) ast.Expression {
	ok, leftExp := p.parseCurToken()
	if !ok {
		return nil
	}
	return p.pratParse(leftExp, priority)
}

func (p *Parser) parseCurToken() (bool, ast.Expression) {
	nud := p.nuds[p.curToken.Type]
	if nud == nil {
		msg := fmt.Sprintf("failed to understand: '%s'", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return false, nil
	}
	return true, nud(p)
}

func (p *Parser) pratParse(leftExp ast.Expression, priority int) ast.Expression {
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

func parseElements(end token.TokenType, p *Parser) []ast.Expression {
	var elements []ast.Expression
	p.nextToken()
	if p.curTokenIs(end) {
		return []ast.Expression{}
	}
	elements = append(elements, p.parseExpression(NO_PRIORITY))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		elements = append(elements, p.parseExpression(NO_PRIORITY))
	}
	if !p.nextTokenAndValidate(end) {
		return nil
	}
	return elements
}
