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

var loopStack []int16

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() (*nast.Program, error) {
	program := &nast.Program{}

	i := 0
	for !p.curTokenIs(token.EOF) {
		var n nast.Node

		switch p.curToken.Type {
		case token.INCREASE, token.DECREASE, token.MOV_L, token.MOV_R, token.WRITE:
			n.Token = p.curToken
			n.Streak = p.countSreak()
		case token.READ:
			n.Token = p.curToken
			n.Streak = 1
		case token.L_B:
			n.Token = p.curToken
			n.Streak = 1
			loopStack = append(loopStack, int16(i))
		case token.R_B:
			n.Token = p.curToken
			if len(loopStack) == 0 {
				panic("the ']' doesn't have the '['")
			}
			var match = loopStack[len(loopStack)-1]

			p := nast.Pointers{Left: match, Right: int16(i)}
			program.Pointers = append(program.Pointers, p)

			loopStack = loopStack[:len(loopStack)-1]
			n.Streak = 1
		}
		program.Nodes = append(program.Nodes, n)
		i++

		p.nextToken()
	}
	if len(loopStack) != 0 {
		panic("the '[' doesn't have the ']'")
	}
	return program, nil
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
func (p *Parser) curTokenIs(tt token.TokenType) bool {
	return tt == p.curToken.Type
}
