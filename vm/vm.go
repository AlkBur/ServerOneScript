package vm

import (
	"fmt"
	"github.com/AlkBur/ServerOneScript/ast"
)

type VM struct {
	stack     []interface{}
	constants []interface{}
	bytecode  []rune
	ip        int
	pp        int
	//scopes    []Scope
	debug  bool
	step   chan struct{}
	curr   chan int
	memory int
	limit  int
}

func Run(program *ast.Tree, env interface{}) (interface{}, error) {
	if program == nil {
		return nil, fmt.Errorf("program is nil")
	}

	vm := VM{}
	return vm.Run(program, env)
}

func (vm *VM) Run(program *ast.Tree, env interface{}) (out interface{}, err error) {
	return
}
