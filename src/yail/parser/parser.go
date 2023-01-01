package parser

import (
	"fmt"
	"yail/ast"
	"yail/ast/statement"
	"yail/lexer"
	"yail/token"
)

type Parser struct {
	lexer  *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: []string{},
	}
	p.curToken = p.lexer.NextToken()
	p.peekToken = p.lexer.NextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []statement.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() statement.Statement {
	if isVariableBindingStatement(p) {
		return parseVariableBindingStatement(p)
	}
	if isReassignmentStatement(p) {
		return parseReassignmentStatement(p)
	}
	return parseExpressionStatement(p)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) nextTokenAndValidate(t token.TokenType) bool {
	p.nextToken()
	if p.curTokenIs(t) {
		return true
	}
	msg := fmt.Sprintf("missing token: %s", t)
	p.errors = append(p.errors, msg)
	return false
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
