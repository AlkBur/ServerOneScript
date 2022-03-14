package vm

import (
	"fmt"
	"github.com/AlkBur/ServerOneScript/ast"
)

func (program *Program) Compile() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	fn := NewFunction("")

	for _, stmt := range program.ast.Nodes {
		program.compile(fn, stmt)
		if program.IsError() {
			err = program.result.Err
			return
		}
	}
	if len(fn.instructions) > 0 {
		program.AddFunction(fn)
	}
	if program.IsError() {
		err = program.result.Err
	}
	return
}

func (program *Program) compile(fn *Function, stmt ast.Node) {
	switch stmt.Type() {
	case ast.Variable:
		program.compileVariable(fn, stmt.(*ast.VariableNode))
		return
	case ast.Number:
		obj := NewObjectNumber(stmt.(*ast.NumberNode).Value)
		program.compileValue(fn, obj)
		return
	case ast.Literal:
		program.compileLiteral(fn, stmt.(*ast.LiteralNode))
		return
	case ast.Return:
		program.compile(fn, stmt.(*ast.ReturnNode).Value)
		program.compileReturn(fn)
		return
	case ast.BinaryOperation:
		program.compileBinaryOperation(fn, stmt.(*ast.BinaryNode))
		return
	}
	program.error("Неизвестная интсрукция %v", stmt.Type())
}

func (program *Program) compileVariable(fn *Function, val *ast.VariableNode) {
	vr := NewVariable(val.Name)
	fn.AddVariable(vr)

	program.compile(fn, val.Value)

	fn.AddInstruction(Instruction{
		opcode: OpPopA,
		object: nil,
	})
	fn.AddInstruction(Instruction{
		opcode: OpSet,
		object: NewObjectVariant(val.Name),
	})

}

func (program *Program) compileValue(fn *Function, val Object) {
	fn.AddInstruction(Instruction{
		opcode: OpLoad,
		object: val,
	})
}

func (program *Program) compileLiteral(fn *Function, val *ast.LiteralNode) {
	fn.AddInstruction(Instruction{
		opcode: OpLoad,
		object: NewObjectVariant(val.Name),
	})
}

func (program *Program) compileReturn(fn *Function) {
	fn.AddInstruction(Instruction{
		opcode: OpPopA,
		object: nil,
	})

	fn.AddInstruction(Instruction{
		opcode: OpReturn,
		object: nil,
	})
}

func (program *Program) compileBinaryOperation(fn *Function, bi *ast.BinaryNode) {
	var op Opcode
	switch bi.Operator {
	case "+":
		op = OpAdd
	case "/":
		op = OpDiv
	case "*":
		op = OpMul
	case "-":
		op = OpSub

	default:
		program.error("Неизвестный оператор %v", bi.Operator)
		return
	}
	program.compile(fn, bi.Right)
	program.compile(fn, bi.Left)

	fn.AddInstruction(Instruction{
		opcode: OpPopA,
		object: nil,
	})

	fn.AddInstruction(Instruction{
		opcode: OpPopB,
		object: nil,
	})

	fn.AddInstruction(Instruction{
		opcode: op,
		object: nil,
	})

	fn.AddInstruction(Instruction{
		opcode: OpPush,
		object: nil,
	})
}
