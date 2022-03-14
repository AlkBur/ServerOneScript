package vm

type Opcode int

const (
	OpPush = iota
	OpReturn
	OpLoad
	OpPopA
	OpPopB
	OpPopNop
	OpSet
	OpAdd
	OpDiv
	OpMul
	OpSub
)

type Instruction struct {
	opcode Opcode
	object Object
}
