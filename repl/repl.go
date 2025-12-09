package repl

import (
	"bufio"
	"fmt"
	"gobf/interpret"
	"io"
)

const PROMPT = "\n ~$ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	I := interpret.New()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		I.Read(line)

		I.Run()
	}
}
