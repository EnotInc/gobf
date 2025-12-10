package evaluator

import (
	"bufio"
	"fmt"
	"gobf/nast"
	"gobf/token"
	"os"
)

const DATASIZE = 16

var pointer int16
var data []int16

func Eval(p *nast.Program) {
	reader := bufio.NewReader(os.Stdin)
	data = make([]int16, DATASIZE)

	for i := 0; i < len(p.Nodes); i++ {
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
		case token.READ:
			value, _ := reader.ReadByte()
			data[pointer] = int16(value)
		case token.L_B:
			if data[pointer] == 0 {
				i = int(p.Pointers[len(p.Pointers)-1].Right)
			}
		case token.R_B:
			if data[pointer] != 0 {
				i = int(p.Pointers[len(p.Pointers)-1].Left)
			}
		case token.EOF:
			return
		}
	}
}
