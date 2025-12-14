// Not Abstract Syntax Tree :)
package nast

import (
	"gobf/token"
)

type Program struct {
	Nodes []Node

	Loops []Loop
}

type Loop struct {
	Left  int16
	Right int16
}

type Node struct {
	Token  token.Token
	Streak int16
	//Streak us to count the streak of the same symbols
	//like "+++++" -> Node.Token.TokenType=token.PlUS, Node.Strak = 5
}
