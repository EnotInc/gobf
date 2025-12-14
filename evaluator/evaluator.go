package evaluator

import (
	"bufio"
	"fmt"
	"gobf/nast"
	"gobf/token"
	"os"
)

// const DATASIZE = 65355
const DATASIZE = 8

type Evaluator struct {
	pointer int16
	data    []int16
}

func New() *Evaluator {
	e := &Evaluator{pointer: 0}
	e.data = make([]int16, DATASIZE)
	return e
}

func (e *Evaluator) Eval(p *nast.Program) {
	reader := bufio.NewReader(os.Stdin)

	read := 0
	for {
		if read >= len(p.Nodes) {
			return
		}
		if e.pointer >= DATASIZE {
			fmt.Print("How did you mange to use all of the cells?")
			return
		}
		n := p.Nodes[read]
		curLoopId := len(p.Loops) - 1

		switch n.Token.Type {
		case token.INCREASE:
			e.data[e.pointer] += n.Streak
		case token.DECREASE:
			e.data[e.pointer] -= n.Streak
		case token.MOV_L:
			e.pointer -= n.Streak
		case token.MOV_R:
			e.pointer += n.Streak
		case token.WRITE:
		case token.READ:
			value, _ := reader.ReadByte()
			e.data[e.pointer] = int16(value)
		case token.L_B:
			if e.data[e.pointer] == 0 {
				read = int(p.Loops[curLoopId].Right)
			}
		case token.R_B:
			fmt.Printf("%v\tpointer: %d\n", p.Loops, e.pointer)
			if e.data[e.pointer] != 0 {
				read = int(p.Loops[curLoopId].Left)
			} else {
				fmt.Print("\tend\n")
			}
		}
		fmt.Printf("\tPointer: %d\tReader: %d\n", e.pointer, read)
		read++
	}
}
