package compiler

import (
	"testing"
	"github.com/gilmae/brainfuck/instruction"
)

type compilerTestCase struct {
	input string
	expectedInstructions []*instruction.Instruction
}

func TestCompile(t *testing.T) {
	tests := []compilerTestCase {
		{
			input: `+++[---[+]>>>]<<<`,
			expectedInstructions: []*instruction.Instruction {
				&instruction.Instruction{Type: instruction.Plus, Argument: 3},
				&instruction.Instruction{Type: instruction.JumpIfZero, Argument: 7},
				&instruction.Instruction{Type: instruction.Minus, Argument: 3},
				&instruction.Instruction{Type: instruction.JumpIfZero, Argument: 5},
				&instruction.Instruction{Type: instruction.Plus, Argument: 1},
				&instruction.Instruction{Type: instruction.JumpIfNotZero, Argument: 3},
				&instruction.Instruction{Type: instruction.Next, Argument: 3},
				&instruction.Instruction{Type: instruction.JumpIfNotZero, Argument: 1},
				&instruction.Instruction{Type: instruction.Prev, Argument: 3},
			},
		},
		{
			input: `,.`,
			expectedInstructions: []*instruction.Instruction {
				&instruction.Instruction{Type: instruction.ReadChar, Argument: 1},
				&instruction.Instruction{Type: instruction.PutChar, Argument: 1},
			},
		},
	}
	runCompilerTests(t,tests)
}

// func TestOptimise(t *testing.T) {
// 	tests := []compilerTestCase{
// 		{
// 			input:`[-]`,
// 			expectedInstructions: []*instruction.Instruction {
// 				&instruction.Instruction{Type:instruction.Clear, Argument:0},
// 			},

// 		},
// 		{
// 			input:`>[-]<`,
// 			expectedInstructions: []*instruction.Instruction {
// 				&instruction.Instruction{Type:instruction.Next, Argument:1},
// 				&instruction.Instruction{Type:instruction.Clear, Argument:0},
// 				&instruction.Instruction{Type:instruction.Prev, Argument:1},
// 			},
// 		},
// 	}

// 	runOptimiserTests(t, tests)
// }

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	for _,tt := range tests {
		c := NewCompiler(tt.input)
		instructions := c.Compile()

		if len(instructions) != len(tt.expectedInstructions) {
			t.Fatalf("wrong number of instructions, want=%d, got=%d", len(tt.expectedInstructions), len(instructions))
		}

		for ix,ins := range tt.expectedInstructions {
			testExpectedInstruction(t, ins, instructions[ix])
		}
	}
}

// func runOptimiserTests(t *testing.T, tests []compilerTestCase) {
// 	for _,tt := range tests {
// 		c := NewCompiler(tt.input)
// 		instructions := c.Compile()
// 		instructions = c.Optimise(instructions)

// 		if len(instructions) != len(tt.expectedInstructions) {
// 			t.Fatalf("wrong number of instructions, want=%d, got=%d", len(tt.expectedInstructions), len(instructions))
// 		}

// 		for ix,ins := range tt.expectedInstructions {
// 			testExpectedInstruction(t, ins, instructions[ix])
// 		}
// 	}
// }

func testExpectedInstruction(t *testing.T, expected *instruction.Instruction, actual *instruction.Instruction) {
	if expected.Type != actual.Type {
		t.Errorf("Wrong instruction, want=%d, got=%d", expected.Type, actual.Type)
	}

	if expected.Argument != actual.Argument {
		t.Errorf("Wrong argument, want=%d, got=%d", expected.Argument, actual.Argument)
	}

}