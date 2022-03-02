package ServerOneScript

import (
	"github.com/AlkBur/ServerOneScript/parser"
	"github.com/AlkBur/ServerOneScript/vm"
	"go/ast"
)

// Compile parses and compiles given input expression to bytecode program.
func Compile(input string) (*vm.Program, error) {
	tree, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}

	_, err = checker.Check(tree, config)

	// If we have a patch to apply, it may fix out error and
	// second type check is needed. Otherwise it is an error.
	if err != nil && len(config.Visitors) == 0 {
		return nil, err
	}

	// Patch operators before Optimize, as we may also mark it as ConstExpr.
	compiler.PatchOperators(&tree.Node, config)

	if len(config.Visitors) >= 0 {
		for _, v := range config.Visitors {
			ast.Walk(&tree.Node, v)
		}
		_, err = checker.Check(tree, config)
		if err != nil {
			return nil, err
		}
	}

	err = optimizer.Optimize(&tree.Node, config)
	if err != nil {
		if fileError, ok := err.(*file.Error); ok {
			return nil, fileError.Bind(tree.Source)
		}
		return nil, err
	}

	program, err := compiler.Compile(tree, config)
	if err != nil {
		return nil, err
	}

	return program, nil
}

// Run evaluates given bytecode program.
func Run(program *vm.Program, env interface{}) (interface{}, error) {
	return vm.Run(program, env)
}
