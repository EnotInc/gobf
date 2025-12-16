package main

import (
	"fmt"
	"gobf/evaluator"
	"gobf/lexer"
	"gobf/parser"
	"gobf/repl"
	"os"
)

const (
	GREETINGS string = " Oh hello there! Be my guest and f_ck around with your brain! You get it? It's a brainf_ck :)"
)

func main() {
	args := os.Args

	E := evaluator.New()

	if len(args) == 2 {
		file := args[1]
		fileContent, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Cannot open file %s.\nError: %v", file, err)
			return
		}

		fileStr := string(fileContent)

		l := lexer.New(fileStr)
		p := parser.New(l)

		program := p.ParseProgram()
		E.Eval(program)

	} else {
		fmt.Print(GREETINGS)
		repl.Start(os.Stdin)
	}
}
