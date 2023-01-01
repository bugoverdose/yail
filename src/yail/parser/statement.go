package parser

import (
	"yail/ast/expression"
	"yail/ast/statement"
	"yail/token"
)

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
	value := p.parseExpression(NO_PREFERENCE)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement.NewVariableBinding(curToken, name, value)
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
	value := p.parseExpression(NO_PREFERENCE)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement.NewReassignment(curToken, name, value)
}
