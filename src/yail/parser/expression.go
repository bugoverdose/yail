package parser

import (
	"fmt"
	"strconv"
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
}

type (
	nullDenotation func(p *Parser) ast.Expression
	leftDenotation func(e ast.Expression, p *Parser) ast.Expression
)

func (p *Parser) initNullDenotations() {
	p.nuds = map[token.TokenType]nullDenotation{
		token.IDENTIFIER:       parseIdentifier,
		token.INTEGER:          parseIntegerLiteral,
		token.TRUE:             parseBoolean,
		token.FALSE:            parseBoolean,
		token.NULL:             parseNull,
		token.NOT:              parsePrefixExpression,
		token.MINUS:            parsePrefixExpression,
		token.LEFT_PARENTHESIS: parseGroupedExpression,
		token.IF:               parseIfExpression,
		token.FUNCTION:         parseFunctionLiteral,
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
		token.LEFT_PARENTHESIS: parseCallExpression, // function call: identifier(parameters)
	}
}

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

func parseIdentifier(p *Parser) ast.Expression {
	return ast.NewIdentifier(p.curToken)
}

func parseIntegerLiteral(p *Parser) ast.Expression {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	return ast.NewIntegerLiteral(p.curToken, value)
}

func parseBoolean(p *Parser) ast.Expression {
	return ast.GetPooledBoolean(p.curTokenIs(token.TRUE))
}

func parseNull(_ *Parser) ast.Expression {
	return ast.NULL
}

func parsePrefixExpression(p *Parser) ast.Expression {
	prefixToken := p.curToken
	p.nextToken()
	rightNode := p.parseExpression(PREFIX_PRIORITY)
	return ast.NewPrefix(prefixToken, rightNode)
}

func parseGroupedExpression(p *Parser) ast.Expression {
	p.nextToken()
	exp := p.parseExpression(NO_PRIORITY) // always parse inside the `(~)` first
	if !p.nextTokenAndValidate(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return exp
}

func parseIfExpression(p *Parser) ast.Expression {
	if !p.nextTokenAndValidate(token.LEFT_PARENTHESIS) {
		return nil
	}
	p.nextToken()
	condition := p.parseExpression(NO_PRIORITY)
	if !p.nextTokenAndValidate(token.RIGHT_PARENTHESIS) {
		return nil
	}
	if !p.nextTokenAndValidate(token.LEFT_BRACKET) {
		return nil
	}
	consequence := parseBlockStatement(p)
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.nextTokenAndValidate(token.LEFT_BRACKET) {
			return nil
		}
		alternative := parseBlockStatement(p)
		return ast.NewIfElse(condition, consequence, alternative)
	}
	return ast.NewIf(condition, consequence)
}

func parseFunctionLiteral(p *Parser) ast.Expression {
	if !p.nextTokenAndValidate(token.LEFT_PARENTHESIS) {
		return nil
	}
	params := parseFunctionParameters(p)
	if !p.nextTokenAndValidate(token.LEFT_BRACKET) {
		return nil
	}
	body := parseBlockStatement(p)
	return ast.NewFunctionLiteral(params, body)
}

func parseFunctionParameters(p *Parser) []*ast.IdentifierExpression {
	var identifiers []*ast.IdentifierExpression
	p.nextToken()
	if p.curTokenIs(token.RIGHT_PARENTHESIS) {
		return identifiers
	}
	identifiers = append(identifiers, ast.NewIdentifier(p.curToken))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		identifiers = append(identifiers, ast.NewIdentifier(p.curToken))
	}
	if !p.nextTokenAndValidate(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return identifiers
}

func parseCallExpression(functionIdentifier ast.Expression, p *Parser) ast.Expression {
	p.nextToken()
	if p.curTokenIs(token.RIGHT_PARENTHESIS) {
		return ast.NewFunctionCall(functionIdentifier, []ast.Expression{})
	}
	return ast.NewFunctionCall(functionIdentifier, parseCallArguments(p))
}

func parseCallArguments(p *Parser) []ast.Expression {
	var args []ast.Expression
	args = append(args, p.parseExpression(NO_PRIORITY))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(NO_PRIORITY))
	}
	if !p.nextTokenAndValidate(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return args
}

func parseInfixExpression(leftNode ast.Expression, p *Parser) ast.Expression {
	infixToken := p.curToken
	priority := p.getCurTokenPriority()
	p.nextToken()
	rightNode := p.parseExpression(priority)
	return ast.NewInfix(leftNode, infixToken, rightNode)
}
