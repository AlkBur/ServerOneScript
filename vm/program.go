package vm

import (
	"errors"
	"fmt"
	"github.com/AlkBur/ServerOneScript/ast"
)

type Program struct {
	variables *Variables
	functions *Functions
	ast       *ast.Tree
	result    *Result
}

type Result struct {
	Msg []string
	Err error
	Obj Object
}

func NewProgram(tree *ast.Tree) *Program {
	return &Program{
		variables: NewVariables(),
		functions: NewFunctions(),
		ast:       tree,
		result: &Result{
			Msg: make([]string, 0),
		},
	}
}

func (prog *Program) AddFunction(fn *Function) {
	prog.functions.Add(fn)
}

func (result *Result) error(err string, arg ...string) {
	result.Err = errors.New(fmt.Sprintf(err, arg))
}

func (prog *Program) error(err string, arg ...interface{}) {
	prog.result.Err = errors.New(fmt.Sprintf(err, arg...))
}

func (prog *Program) IsError() bool {
	return prog.result.Err != nil
}

func (prog *Program) loadObject(fn *Function, varObj Object) Object {
	if varObj.Type() == TypeVariant {
		vr := varObj.(*Variant)
		obj, ok := fn.variables.data[vr.name]
		if ok {
			return obj.value
		}
		obj, ok = prog.variables.data[vr.name]
		if ok {
			return obj.value
		}
		prog.error("Ошибка в получении переменной. Переменная не найдена %v", vr.name)
	}
	return varObj
}

func (prog *Program) setVariantObject(fn *Function, varObj Object, obj Object) {
	switch varObj.Type() {
	case TypeVariant:
		vr := varObj.(*Variant)
		setVr, ok := fn.variables.data[vr.name]
		if ok {
			setVr.value = obj
			return
		}
		setVr, ok = prog.variables.data[vr.name]
		if ok {
			setVr.value = obj
			return
		}
		prog.error("Ошибка в установке переменной. Переменная не найдена %v", vr.name)
	default:
		prog.error("Ошибка в установке переменной")
	}
}
