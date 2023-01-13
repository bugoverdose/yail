package parser

import (
	"fmt"
	"yail/ast"
	"yail/token"
)

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
		token.LEFT_PARENTHESIS: parseFunctionCallExpression,
		token.LEFT_BRACKET:     parseCollectionAccessExpression,
	}
}

func parseInfixExpression(leftNode ast.Expression, p *Parser) ast.Expression {
	infixToken := p.curToken
	priority := p.getCurTokenPriority()
	p.nextToken()
	rightNode := p.parseExpression(priority)
	return ast.NewInfix(leftNode, infixToken, rightNode)
}

func parseFunctionCallExpression(function ast.Expression, p *Parser) ast.Expression {
	functionIdentifier, ok := function.(*ast.IdentifierExpression)
	if !ok {
		msg := fmt.Sprintf("unsupported operation : %s(", functionIdentifier.Value)
		p.errors = append(p.errors, msg)
	}
	args := parseElements(token.RIGHT_PARENTHESIS, p)
	return ast.NewFunctionCall(functionIdentifier, args)
}

func parseCollectionAccessExpression(left ast.Expression, p *Parser) ast.Expression {
	p.nextToken()
	index := p.parseExpression(NO_PRIORITY)
	if !p.nextTokenAndValidate(token.RIGHT_BRACKET) {
		return nil
	}
	return ast.NewCollectionAccess(left, index)
}
