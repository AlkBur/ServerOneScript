package ast

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Number
	Array
	Func
	Map
	String
	Struct
	BinaryOperation
	Variable
	Date
	Undefined
)

func (t Kind) Equal(k Kind) bool {
	return t == k
}
