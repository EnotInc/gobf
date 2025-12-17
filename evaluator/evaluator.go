package evaluator

import (
	"bufio"
	"fmt"
	"gobf/nast"
	"gobf/token"
	"os"
)

const DATASIZE uint16 = 65535

type Evaluator struct {
	pointer uint16
	data    []byte
	reader  int16
}

func New() *Evaluator {
	e := &Evaluator{}
	e.data = make([]byte, DATASIZE)
	return e
}

func (e *Evaluator) Eval(p *nast.Program) {
	r := bufio.NewReader(os.Stdin)

	for {

		if e.reader >= int16(len(p.Nodes)) {
			return
		}

		node := p.Nodes[e.reader]

		switch node.Token.Type {
		case token.INCREASE:
			e.data[e.pointer] += byte(node.Streak)
		case token.DECREASE:
			e.data[e.pointer] -= byte(node.Streak)
		case token.MOV_L:
			e.pointer -= node.Streak
		case token.MOV_R:
			e.pointer += node.Streak
		case token.WRITE:
			fmt.Printf("%c", e.data[e.pointer])
		case token.READ:
			value, _ := r.ReadByte()
			e.data[e.pointer] = value
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
