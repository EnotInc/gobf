package evaluator

import (
	"bufio"
	"fmt"
	"gobf/nast"
	"gobf/token"
	"os"
)

const DATASIZE int16 = 32767

type Evaluator struct {
	pointer int16
	data    []int16
	reader  int16
}

func New() *Evaluator {
	e := &Evaluator{}
	e.data = make([]int16, DATASIZE)
	return e
}

func (e *Evaluator) Eval(p *nast.Program) {
	r := bufio.NewReader(os.Stdin)

	for {
		if e.reader >= int16(len(p.Nodes)) {
			return
		}
		if e.pointer >= DATASIZE {
			fmt.Print("How did you mange to use all of the cells?")
			return
		}

		node := p.Nodes[e.reader]

		switch node.Token.Type {
		case token.INCREASE:
			e.data[e.pointer] += node.Streak
		case token.DECREASE:
			e.data[e.pointer] -= node.Streak
		case token.MOV_L:
			e.pointer -= node.Streak
		case token.MOV_R:
			e.pointer += node.Streak
		case token.WRITE:
			fmt.Printf("%c", e.data[e.pointer])
		case token.READ:
			value, _ := r.ReadByte()
			e.data[e.pointer] = int16(value)
		case token.L_B:
			if e.data[e.pointer] == 0 {
				e.reader = p.Forward[e.reader]
			}
		case token.R_B:
			if e.data[e.pointer] != 0 {
				e.reader = p.Backword[e.reader]
			}
		}

		e.reader++ // Moving to the next node
	}

}
