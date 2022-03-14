package vm

import (
	"github.com/AlkBur/ServerOneScript/parser"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTest(t *testing.T, code string, result Object) {
	assert := assert.New(t)

	p := parser.NewFromString(code, "")
	ast, err := p.Parse()
	assert.NoError(err, "parser")
	prog := NewProgram(ast)
	err = prog.Compile()
	assert.NoError(err, "compiler")

	out := Run(prog, "")
	if assert.NotNil(out) {

		assert.NoError(out.Err, "vm")
		assert.Equal(out.Obj.String(), result.String(), "return1")
		assert.True(out.Obj.Equal(result), "return2")
	}
}

func TestArithmetic(t *testing.T) {
	code := `a = 1;
	return a;`

	result := NewObjectNumber(decimal.NewFromInt(1))
	RunTest(t, code, result)

	code = `a = 1 + 2;
	return a;`

	result = NewObjectNumber(decimal.NewFromInt(3))
	RunTest(t, code, result)

	code = `a = 1 / 2;
	return a;`

	result = NewObjectNumber(decimal.NewFromFloat(0.5))
	RunTest(t, code, result)

	code = `a =  5-3;
	return a;`

	result = NewObjectNumber(decimal.NewFromFloat(2))
	RunTest(t, code, result)

	code = `a =  (15-3)/2;
	return a;`

	result = NewObjectNumber(decimal.NewFromFloat(6))
	RunTest(t, code, result)

}
