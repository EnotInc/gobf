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

func (p *Parser) ParseProgram() *nast.Program {
	program := &nast.Program{
		Forward:  make(map[int16]int16),
		Backword: make(map[int16]int16),
	}

	var i int16 = 0 //counting the tokens that we iterating
	for !p.curTokenIs(token.EOF) {
		var n nast.Node

		switch p.curToken.Type {
		case token.INCREASE, token.DECREASE, token.MOV_L, token.MOV_R:
			n.Streak = p.countSreak()
		case token.L_B:
			n.Streak = 1
			loopStack = append(loopStack, int16(i))
		case token.R_B:
			if len(loopStack) == 0 {
				panic("the ']' doesn't have the '[' pair")
			}

			left := loopStack[len(loopStack)-1]
			program.Forward[left] = i
			program.Backword[i] = left

			loopStack = loopStack[:len(loopStack)-1]
			n.Streak = 1
		case token.WRITE, token.READ:
			n.Streak = 1
		}
		n.Token = p.curToken

		program.Nodes = append(program.Nodes, n)
		i++

		p.nextToken()
	}
	if len(loopStack) != 0 {
		panic("the '[' doesn't have the ']'")
	}

	return program
}

func (p *Parser) countSreak() uint16 {
	var amount uint16 = 1
	for p.curToken == p.peekToken {
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
