package ast

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Number
	Array
	Function
	FunctionCall
	ProcedureCall
	Map
	String
	Struct
	BinaryOperation
	Variable
	Literal
	Date
	Undefined
	Continue
	Null
)

func (t Kind) Equal(k Kind) bool {
	return t == k
}
