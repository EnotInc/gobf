package lexer

import (
	"gobf/token"
	"os"
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
	var tok token.Token

	l.readChar()
	l.skepWhiteSpace()

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
	case 'q':
		os.Exit(0) //place it here so it will exit the programm without any other calculations
	case 0:
		tok = l.newToken(token.EOF, l.ch)
	}

	return tok
}

func (l *Lexer) readChar() {
	if l.position > len(l.input)-1 {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
	l.position++
}

func (l *Lexer) newToken(tt token.TokenType, ch byte) token.Token {
	return token.Token{Type: tt, Literal: ch}
}

func (l *Lexer) skepWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
