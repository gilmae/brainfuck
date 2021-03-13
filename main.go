package main

import (
	"fmt"
	"github.com/gilmae/brainfuck/machine"
	"github.com/gilmae/brainfuck/compiler"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
	c := compiler.NewCompiler(string(code))
	instructions := c.Compile()
	m := machine.NewVM(instructions, os.Stdin, os.Stdout)
	m.Run()
}
