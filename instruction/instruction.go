package instruction

type InsType byte

const (
	Plus          InsType = '+'
	Minus         InsType = '-'
	Next          InsType = '>'
	Prev          InsType = '<'
	ReadChar      InsType = ','
	PutChar       InsType = '.'
	JumpIfZero    InsType = '['
	JumpIfNotZero InsType = ']'
	Clear         InsType = '0'
)

type Instruction struct {
	Type     InsType
	Argument int
}
