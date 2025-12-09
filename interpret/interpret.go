package interpret

import (
	"bufio"
	"fmt"
	"os"
)

const DATASIZE int = 65535 //some random number that is power of 2, ig

const (
	INCREASE = '+'
	DECREASE = '-'
	MOV_R    = '>'
	MOV_L    = '<'
	L_B      = '['
	R_B      = ']'
	WRITE    = '.'
	READ     = ','
)

type Interpret struct {
	input     string
	position  int // position at given string
	pointer   int // pointer on data arrey
	ch        byte
	data      []int
	loopStack []int
}

func New() *Interpret {
	i := &Interpret{}
	i.data = make([]int, DATASIZE)
	return i
}

func (i *Interpret) Read(input string) {
	i.position = 0
	i.input = input
	fmt.Print(" ")
}

func (i *Interpret) NextToken() bool {
	if i.position >= len(i.input)-1 {
		return false
	}
	//TODO: implement scipWhiteSpace func
	i.position++
	return true
}

func (i *Interpret) ExecToken() bool {

	if i.pointer >= DATASIZE {
		fmt.Printf(" How tf did you manage to run out of space?\nThere was %d cells reserved and somehow you need to use more then that?!", DATASIZE)
		return false
	}

	reader := bufio.NewReader(os.Stdin)

	i.ch = i.input[i.position]

	switch i.ch {
	case INCREASE:
		i.data[i.pointer]++
	case DECREASE:
		i.data[i.pointer]--
	case MOV_R:
		i.pointer++
	case MOV_L:
		i.pointer--
	case L_B:
		i.loopStack = append(i.loopStack, i.position) //saving position to the stack as value
	case R_B:
		if len(i.loopStack) == 0 {
			fmt.Print("the ']' doesn't have the '[' match")
			return false
		}
		var currentLoopStart = i.loopStack[len(i.loopStack)-1] // gettig the saved '[' position
		if i.data[i.pointer] != 0 {
			i.position = currentLoopStart // setting position after the '['
		} else {
			i.loopStack = i.loopStack[:len(i.loopStack)-1] //deleteng the last ']' position from stack
		}
	case WRITE:
		fmt.Printf("%c", i.data[i.pointer])
	case READ:
		value, _ := reader.ReadByte()
		i.data[i.pointer] = int(value)
	default:
		fmt.Printf("Unckown char: %c", i.ch)
		return false
	}

	return true
}

func (i Interpret) Run() {

	for {
		if !i.ExecToken() {
			return
		}
		if !i.NextToken() {
			break
		}
	}
}
