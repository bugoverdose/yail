package parser

import (
	"yail/ast"
	"yail/token"
)

func isVariableBindingStatement(p *Parser) bool {
	return p.curTokenIs(token.VAR) || p.curTokenIs(token.VAL)
}

func parseVariableBindingStatement(p *Parser) *ast.VariableBindingStatement {
	curToken := p.curToken
	if !p.nextTokenAndValidate(token.IDENTIFIER) {
		return nil
	}
	name := ast.NewIdentifierFrom(p.curToken.Literal)

	if !p.nextTokenAndValidate(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	value := p.parseExpression(NO_PRIORITY)

	if !p.nextTokenAndValidate(token.SEMICOLON) {
		return nil
	}
	return ast.NewVariableBinding(curToken, name, value)
}

func isReassignmentStatement(p *Parser) bool {
	return p.curTokenIs(token.IDENTIFIER) && p.peekTokenIs(token.ASSIGN)
}

func parseReassignmentStatement(p *Parser) *ast.ReassignmentStatement {
	curToken := p.curToken
	if !p.curTokenIs(token.IDENTIFIER) {
		return nil
	}
	name := ast.NewIdentifierFrom(p.curToken.Literal)

	if !p.nextTokenAndValidate(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	value := p.parseExpression(NO_PRIORITY)

	if !p.nextTokenAndValidate(token.SEMICOLON) {
		return nil
	}
	return ast.NewReassignment(curToken, name, value)
}
