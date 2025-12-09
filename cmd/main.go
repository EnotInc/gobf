package main

import (
	"fmt"
	"gobf/interpret"
	"gobf/repl"
	"os"
)

const (
	GREETINGS string = " Oh hello there! Be my guest and f_ck around! You get it? It's a brainf_ck :)"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		file := args[1]
		fileContent, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Cannot open file %s", file)
			return
		}

		I := interpret.New()
		I.Read(string(fileContent))
		I.Run()
	} else {
		fmt.Print(GREETINGS)
		repl.Start(os.Stdin, os.Stdout)
	}
}
