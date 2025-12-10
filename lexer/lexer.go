package lexer

import (
	"gobf/token"
)

type Lexer struct {
	input    string
	position int
	ch       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.position > len(l.input)-1 {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}

	var tok token.Token

	switch l.ch {
	case '+':
		tok = l.newToken(token.INCREASE, l.ch)
	case '-':
		tok = l.newToken(token.DECREASE, l.ch)
	case '<':
		tok = l.newToken(token.MOV_L, l.ch)
	case '>':
		tok = l.newToken(token.MOV_R, l.ch)
	case '.':
		tok = l.newToken(token.WRITE, l.ch)
	case ',':
		tok = l.newToken(token.READ, l.ch)
	case '[':
		tok = l.newToken(token.L_B, l.ch)
	case ']':
		tok = l.newToken(token.R_B, l.ch)
	case 0:
		tok = l.newToken(token.EOF, l.ch)
	}

	l.position++
	return tok
}

func (l *Lexer) newToken(tt token.TokenType, ch byte) token.Token {
	return token.Token{Type: tt, Literal: ch}
}
