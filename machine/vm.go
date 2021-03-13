package machine

import (
	"io"
	"github.com/gilmae/brainfuck/instruction"
)

type VM struct {
	code []*instruction.Instruction
	ip   int

	memory [30000]int
	dp     int

	input  io.Reader
	output io.Writer

	buffer []byte
}

func NewVM(code []*instruction.Instruction, input io.Reader, output io.Writer) *VM {
	return &VM{code: code, input: input, output: output, buffer: make([]byte, 1)}
}

func (vm *VM) Run() {
	for vm.ip < len(vm.code) {
		ins := vm.code[vm.ip]

		switch ins.Type {
		case instruction.Plus:
			vm.memory[vm.dp] += ins.Argument
		case instruction.Minus:
			vm.memory[vm.dp] -= ins.Argument
		case instruction.Next:
			vm.dp += ins.Argument
		case instruction.Prev:
			vm.dp -= ins.Argument
		case instruction.ReadChar:
			for ix := 0; ix < ins.Argument; ix++ {
				vm.getChar()
			}
		case instruction.PutChar:
			for ix := 0; ix < ins.Argument; ix++ {
				vm.putChar()
			}
		case instruction.JumpIfZero:
			if vm.memory[vm.dp] == 0 {
				vm.ip = ins.Argument
			}
		case instruction.JumpIfNotZero:
			if vm.memory[vm.dp] != 0 {
				vm.ip = ins.Argument
			}
		case instruction.Clear:
			vm.memory[vm.dp] = 0
		}

		vm.ip++
	}
}

func (vm *VM) getChar() {
	n, err := vm.input.Read(vm.buffer)

	if err != nil {
		panic(err)
	}

	if n != 0 {
		panic("wrong number of bytes read")
	}

	vm.memory[vm.dp] = int(vm.buffer[0])
}

func (vm *VM) putChar() {
	vm.buffer[0] = byte(vm.memory[vm.dp])

	n, err := vm.output.Write(vm.buffer)

	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("wrong number of bytes written")
	}
}
