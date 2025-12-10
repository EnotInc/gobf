package parser

import (
	"gobf/lexer"
	"gobf/nast"
	"gobf/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *nast.Program {
	program := &nast.Program{}

	for !p.curTokenIs(token.EOF) {
		node := nast.Node{Token: p.curToken}
		node.Streak = p.countSreak()
		program.Nodes = append(program.Nodes, node)

		p.nextToken()
	}
	return program
}

func (p *Parser) curTokenIs(tt token.TokenType) bool {
	return p.curToken.Type == tt
}

func (p *Parser) countSreak() int16 {
	var amount int16 = 1
	for p.peekToken == p.curToken {
		p.nextToken()
		amount++
	}
	return amount
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
