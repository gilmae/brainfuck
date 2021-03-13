package compiler

import (
	"github.com/gilmae/brainfuck/instruction"
)

type Compiler struct {
	code         string
	ip           int
	length       int
	instructions []*instruction.Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{code: code, length: len(code), instructions: []*instruction.Instruction{}}
}

func (c *Compiler) Compile() []*instruction.Instruction {
	loopStack := []int{}
	for c.ip < c.length {
		ins := c.code[c.ip]

		switch ins {
		case '+':
			c.foldInstruction(instruction.Plus)
		case '-':
			c.foldInstruction(instruction.Minus)
		case '>':
			c.foldInstruction(instruction.Next)
		case '<':
			c.foldInstruction(instruction.Prev)
		case '.':
			c.foldInstruction(instruction.PutChar)
		case ',':
			c.foldInstruction(instruction.ReadChar)
		case '[':
			insPos := c.emit(instruction.JumpIfZero)
			loopStack = append(loopStack, insPos)
		case ']':
			openInsPos := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			closeInsPos := c.emitWithArg(instruction.JumpIfNotZero, openInsPos)

			c.instructions[openInsPos].Argument = closeInsPos
		}

		c.ip++
	}

	return c.instructions
}

// func (c *Compiler) Optimise(instructions []*instruction.Instruction) []*instruction.Instruction {
// 	for ix := 0;ix < len(instructions);ix++ {
// 		ins := instructions[ix]
// 		if ins.Type == instruction.JumpIfZero {
// 			if instructions[ix+1].Type == instruction.Minus && instructions[ix+2].Type == instruction.JumpIfNotZero {
// 				// Replace with Clear

// 				instructions[ix].Type = instruction.Clear
// 				instructions[ix].Argument = 0
// 				instructions =  removeInstructions(instructions, ix+1, 2)
// 			}
// 		}
// 	}

// 	return instructions
// }

func (c *Compiler) foldInstruction(insType instruction.InsType) {
	ch := c.code[c.ip]
	count := 1
	for c.ip < c.length-1 && c.code[c.ip+1] == ch {
		count++
		c.ip++
	}

	c.emitWithArg(insType, count)
}

func (c *Compiler) emit(insType instruction.InsType) int {
	ins := &instruction.Instruction{Type: insType}
	c.instructions = append(c.instructions, ins)

	return len(c.instructions) - 1 // Position of emitted instruction
}

func (c *Compiler) emitWithArg(insType instruction.InsType, arg int) int {
	ins := &instruction.Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)

	return len(c.instructions) - 1 // Position of emitted instruction
}

func removeInstructions(instructions []*instruction.Instruction, position int, count int) []*instruction.Instruction{
	ret := make([]*instruction.Instruction, 0)
    ret = append(ret, instructions[:position]...)
    return append(ret, instructions[position+count:]...)
}
