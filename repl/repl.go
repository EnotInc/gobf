package repl

import (
	"bufio"
	"fmt"
	"gobf/evaluator"
	"gobf/lexer"
	"gobf/nast"
	"gobf/parser"
	"io"
)

const (
	COLOR  = "\033[92m"
	CLEAR  = "\033[0m"
	PROMPT = "\n " + COLOR + "~$ " + CLEAR // ~$
)

func Start(in io.Reader, out io.Writer) { //Do I really need `out` here?
	scanner := bufio.NewScanner(in)
	var program *nast.Program

	E := evaluator.New()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program = p.ParseProgram()
		E.Eval(program)
	}
}
