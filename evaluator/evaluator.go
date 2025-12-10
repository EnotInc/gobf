package evaluator

import (
	"fmt"
	"gobf/nast"
	"gobf/token"
)

const DATASIZE = 16

var pointer int16
var data []int16

func Eval(p *nast.Program) {
	data = make([]int16, DATASIZE)

	for i := range p.Nodes {
		n := p.Nodes[i]

		switch n.Token.Type {
		case token.INCREASE:
			data[pointer] += n.Streak
		case token.DECREASE:
			data[pointer] -= n.Streak
		case token.MOV_L:
			pointer -= n.Streak
		case token.MOV_R:
			pointer += n.Streak
		case token.WRITE:
			fmt.Printf("%c", data[pointer])
		case token.EOF:
			return
		}
	}
}
