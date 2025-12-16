// Not Abstract Syntax Tree :)
package nast

import (
	"gobf/token"
)

type Program struct {
	Nodes []Node

	Forward  map[int16]int16
	Backword map[int16]int16
}

type Node struct {
	Token  token.Token
	Streak int16
	//Streak us to count the streak of the same symbols
	//like "+++++" -> Node.Token.TokenType=token.PlUS, Node.Strak = 5
}
