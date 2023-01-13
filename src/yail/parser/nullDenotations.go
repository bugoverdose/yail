package parser

import (
	"fmt"
	"strconv"
	"yail/ast"
	"yail/token"
)

func (p *Parser) initNullDenotations() {
	p.nuds = map[token.TokenType]nullDenotation{
		token.IDENTIFIER:       parseIdentifier,
		token.INTEGER:          parseIntegerLiteral,
		token.STRING:           parseStringLiteral,
		token.TRUE:             parseBooleanLiteral,
		token.FALSE:            parseBooleanLiteral,
		token.NULL:             parseNull,
		token.NOT:              parsePrefixExpression,
		token.MINUS:            parsePrefixExpression,
		token.LEFT_PARENTHESIS: parseGroupedExpression,
		token.IF:               parseIfExpression,
		token.FUNCTION:         parseFunctionLiteral,
		token.LEFT_BRACKET:     parseArrayLiteral,
		token.LEFT_BRACE:       parseHashLiteral,
	}
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

func parseStringLiteral(p *Parser) ast.Expression {
	return ast.NewStringLiteral(p.curToken)
}

func parseBooleanLiteral(p *Parser) ast.Expression {
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
	if !p.nextTokenAndValidate(token.LEFT_BRACE) {
		return nil
	}
	consequence := parseBlockStatement(p)
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.nextTokenAndValidate(token.LEFT_BRACE) {
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
	if !p.nextTokenAndValidate(token.LEFT_BRACE) {
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

func parseArrayLiteral(p *Parser) ast.Expression {
	elements := parseElements(token.RIGHT_BRACKET, p)
	return ast.NewArrayLiteral(elements)
}

func parseHashLiteral(p *Parser) ast.Expression {
	pairs := make(map[ast.Expression]ast.Expression)
	for !p.peekTokenIs(token.RIGHT_BRACE) {
		p.nextToken()
		key := p.parseExpression(NO_PRIORITY)
		if !p.nextTokenAndValidate(token.COLON) {
			return nil
		}
		p.nextToken()
		value := p.parseExpression(NO_PRIORITY)
		pairs[key] = value
		if !p.peekTokenIs(token.RIGHT_BRACE) && !p.nextTokenAndValidate(token.COMMA) {
			return nil
		}
	}
	if !p.nextTokenAndValidate(token.RIGHT_BRACE) {
		return nil
	}
	return ast.NewHashMapLiteral(pairs)
}
