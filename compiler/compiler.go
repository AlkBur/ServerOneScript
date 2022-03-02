package compiler

import (
	"errors"
	"fmt"
	"github.com/AlkBur/ServerOneScript/ast"
	"github.com/AlkBur/ServerOneScript/vm"
	"reflect"
)

type compiler struct {
	constants []interface{}
	bytecode  []rune
	index     map[interface{}]uint16
	mapEnv    bool
	cast      reflect.Kind
	nodes     []ast.Node
}

func Compile(tree *ast.Tree) (program *vm.Program, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	c := &compiler{
		index: make(map[interface{}]uint16),
	}

	for _, node := range tree.Nodes {
		c.compile(node)
	}

	program = &vm.Program{
		Constants: c.constants,
		Bytecode:  c.bytecode,
	}
	return
}

func (c *compiler) compile(node ast.Node) (err error) {
	c.nodes = append(c.nodes, node)
	switch n := node.(type) {
	case *ast.VariableExpr:
		c.IdentifierNode(n)
	default:
		err = errors.New(fmt.Sprintf("undefined node type (%T)", node))
	}
	return
}
