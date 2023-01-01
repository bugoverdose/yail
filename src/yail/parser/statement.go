package parser

import (
	"yail/ast/expression"
	"yail/ast/statement"
	"yail/token"
)

func isVariableBindingStatement(p *Parser) bool {
	return p.curTokenIs(token.VAR) || p.curTokenIs(token.VAL)
}

func parseVariableBindingStatement(p *Parser) *statement.VariableBinding {
	curToken := p.curToken
	if !p.nextTokenAndValidate(token.IDENTIFIER) {
		return nil
	}
	name := expression.NewIdentifierFrom(p.curToken.Literal)

	if !p.nextTokenAndValidate(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	value := p.parseExpression(NO_PRIORITY)

	if !p.nextTokenAndValidate(token.SEMICOLON) {
		return nil
	}
	return statement.NewVariableBinding(curToken, name, value)
}

func isReassignmentStatement(p *Parser) bool {
	return p.curTokenIs(token.IDENTIFIER) && p.peekTokenIs(token.ASSIGN)
}

func parseReassignmentStatement(p *Parser) *statement.Reassignment {
	curToken := p.curToken
	if !p.curTokenIs(token.IDENTIFIER) {
		return nil
	}
	name := expression.NewIdentifierFrom(p.curToken.Literal)

	if !p.nextTokenAndValidate(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	value := p.parseExpression(NO_PRIORITY)

	if !p.nextTokenAndValidate(token.SEMICOLON) {
		return nil
	}
	return statement.NewReassignment(curToken, name, value)
}
