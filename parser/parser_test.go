package parser

import (
	"github.com/AlkBur/ServerOneScript/ast"
	"github.com/AlkBur/ServerOneScript/runes"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	//p := NewFromString("a", "")
	//tree, err := p.Parse()
	//node := ast.NewVariableNode("a")

	parseTests := []struct {
		input    string
		expected ast.Node
	}{
		{
			`Перем a`,
			&ast.VariableNode{Name: "a", BaseNode: *ast.NewBaseNode(ast.Variable, runes.NewLocation(1, 7)),
				Value: ast.NewBaseNode(ast.Undefined, runes.NewLocation(1, 7))},
		},
		{
			`a = "str"`,
			&ast.VariableNode{Name: "a", BaseNode: *ast.NewBaseNode(ast.Variable, runes.NewLocation(1, 1)),
				Value: &ast.StringNode{Value: "str", BaseNode: *ast.NewBaseNode(ast.String, runes.NewLocation(1, 5))},
			},
		},
		{
			`a = 3`,
			&ast.VariableNode{Name: "a", BaseNode: *ast.NewBaseNode(ast.Variable, runes.NewLocation(1, 1)),
				Value: &ast.NumberNode{Value: decimal.NewFromInt(3), BaseNode: *ast.NewBaseNode(ast.Number, runes.NewLocation(1, 5))},
			},
		},
		{
			`a = 2.5`,
			&ast.VariableNode{Name: "a", BaseNode: *ast.NewBaseNode(ast.Variable, runes.NewLocation(1, 1)),
				Value: &ast.NumberNode{Value: decimal.NewFromFloat(2.5), BaseNode: *ast.NewBaseNode(ast.Number, runes.NewLocation(1, 5))},
			},
		},
		{
			`a = 10000000`,
			&ast.VariableNode{Name: "a", BaseNode: *ast.NewBaseNode(ast.Variable, runes.NewLocation(1, 1)),
				Value: &ast.NumberNode{Value: decimal.NewFromInt(10000000), BaseNode: *ast.NewBaseNode(ast.Number, runes.NewLocation(1, 5))},
			},
		},

		/*
							{
								"true",
								ast.NewBoolNode("true"),
							},
							{
								"false",
								ast.NewBoolNode("false"),
							},
							{
								"nil",
								ast.NewNilNode(),
							},
							{
								"-3",
								ast.NewNumberNode("-3"),
							},
							{
								"1 - 2",
								ast.NewBinaryNode(Operator: "-", Left: &ast.NewIntegerNode(Value: ), Right: &ast.NewIntegerNode(Value: 2),
							},
							{
								"(1 - 2) * 3",
								&ast.BinaryNode{
									Operator: "*",
									Left: &ast.NewBinaryNode{
			,							Operator: "-", Left: &ast.NewNumberNode("1"),
										Right: &ast.NewNumberNode("2"),
									}, Right: &ast.NewNumberNode("3"),
								},
							},
							{
								"a or b or c",
								ast.NewBinaryNode(Operator: "or", Left: &ast.NewBinaryNode(Operator: "or", Left: &ast.NewIdentifierNode(Value: "a), Right: &ast.NewIdentifierNode(Value: "b"), Right: &ast.NewIdentifierNode(Value: "c"),
							},
							{
								"a or b and c",
								ast.NewBinaryNode(Operator: "or", Left: &ast.NewIdentifierNode(Value: "a), Right: &ast.NewBinaryNode(Operator: "and", Left: &ast.NewIdentifierNode(Value: "b), Right: &ast.NewIdentifierNode(Value: "c"}),
							},
							{
								"(a or b) and c",
								ast.NewBinaryNode(Operator: "and", Left: &ast.NewBinaryNode(Operator: "or", Left: &ast.NewIdentifierNode(Value: "a), Right: &ast.NewIdentifierNode(Value: "b"), Right: &ast.NewIdentifierNode(Value: "c"),
							},
							{
								"2**4-1",
								ast.NewBinaryNode(Operator: "-", Left: &ast.NewBinaryNode(Operator: "**", Left: &ast.NewIntegerNode(Value: ), Right: &ast.NewIntegerNode(Value: 4), Right: &ast.NewIntegerNode(Value: 1),
							},
							{
								"foo(bar())",
								ast.NewFunctionNode(Name: "foo", Arguments: []ast.NewNode(&ast.NewFunctionNode(Name: "bar", Arguments: []ast.NewNode(}}),
							},
							{
								`foo("arg1", 2, true)`,
								ast.NewFunctionNode(Name: "foo", Arguments: []ast.NewNode(&ast.NewStringNode(Value: "arg1), &ast.NewIntegerNode(Value: ), &ast.NewBoolNode(Value: true}),
							},
							{
								"foo.bar",
								ast.NewPropertyNode(Node: &ast.NewIdentifierNode(Value: "foo), Property: "bar),
							},
							{
								"foo?.bar",
								ast.NewPropertyNode(Node: &ast.NewIdentifierNode(Value: "foo", NilSafe: tru), Property: "bar", NilSafe: tru),
							},
							{
								"foo['all']",
								ast.NewIndexNode(Node: &ast.NewIdentifierNode(Value: "foo), Index: &ast.NewStringNode(Value: "all"),
							},
							{
								"foo.bar()",
								ast.NewMethodNode(Node: &ast.NewIdentifierNode(Value: "foo), Method: "bar", Arguments: []ast.NewNode(),
							},
							{
								`foo.bar("arg1", 2, true)`,
								ast.NewMethodNode(Node: &ast.NewIdentifierNode(Value: "foo), Method: "bar", Arguments: []ast.NewNode(&ast.NewStringNode(Value: "arg1), &ast.NewIntegerNode(Value: ), &ast.NewBoolNode(Value: true}),
							},
							{
								"foo[3]",
								ast.NewIndexNode(Node: &ast.NewIdentifierNode(Value: "foo), Index: &ast.NewIntegerNode(Value: 3),
							},
							{
								"true ? true : false",
								ast.NewConditionalNode(Cond: &ast.NewBoolNode(Value: tru), Exp1: &ast.NewBoolNode(Value: tru), Exp2: &ast.NewBoolNode(),
							},
							{
								"foo.bar().foo().baz[33]",
								ast.NewIndexNode(
									Node: ast.NewPropertyNode(Node: &ast.NewMethodNode(Node: &ast.NewMethodNode(
										Node: ast.NewIdentifierNode(Value: "foo), Method: "bar", Arguments: []ast.NewNode),
									}, Method: "foo", Arguments: []ast.NewNode(), Property: "baz),
									Index: ast.NewIntegerNode(Value: 3),
								},
							},
							{
								"'a' == 'b'",
								ast.NewBinaryNode(Operator: "==", Left: &ast.NewStringNode(Value: "a), Right: &ast.NewStringNode(Value: "b"),
							},
							{
								"+0 != -0",
								ast.NewBinaryNode(Operator: "!=", Left: &ast.NewUnaryNode(Operator: "+", Node: &ast.NewIntegerNode(), Right: &ast.NewUnaryNode(Operator: "-", Node: &ast.NewIntegerNode(}),
							},
							{
								"[a, b, c]",
								ast.NewArrayNode(Nodes: []ast.NewNode(&ast.NewIdentifierNode(Value: "a), &ast.NewIdentifierNode(Value: "b), &ast.NewIdentifierNode(Value: "c"}),
							},
							{
								"{foo:1, bar:2}",
								ast.NewMapNode(Pairs: []ast.NewNode(&ast.NewPairNode(Key: &ast.NewStringNode(Value: "foo), Value: &ast.NewIntegerNode(Value: 1), &ast.NewPairNode(Key: &ast.NewStringNode(Value: "bar), Value: &ast.NewIntegerNode(Value: 2}}),
							},
							{
								"{foo:1, bar:2, }",
								ast.NewMapNode(Pairs: []ast.NewNode(&ast.NewPairNode(Key: &ast.NewStringNode(Value: "foo), Value: &ast.NewIntegerNode(Value: 1), &ast.NewPairNode(Key: &ast.NewStringNode(Value: "bar), Value: &ast.NewIntegerNode(Value: 2}}),
							},
							{
								`{"a": 1, 'b': 2}`,
								ast.NewMapNode(Pairs: []ast.NewNode(&ast.NewPairNode(Key: &ast.NewStringNode(Value: "a), Value: &ast.NewIntegerNode(Value: 1), &ast.NewPairNode(Key: &ast.NewStringNode(Value: "b), Value: &ast.NewIntegerNode(Value: 2}}),
							},
							{
								"[1].foo",
								ast.NewPropertyNode(Node: &ast.NewArrayNode(Nodes: []ast.NewNode(&ast.NewIntegerNode(Value: 1}), Property: "foo),
							},
							{
								"{foo:1}.bar",
								ast.NewPropertyNode(Node: &ast.NewMapNode(Pairs: []ast.NewNode(&ast.NewPairNode(Key: &ast.NewStringNode(Value: "foo), Value: &ast.NewIntegerNode(Value: 1}}), Property: "bar),
							},
							{
								"len(foo)",
								ast.NewBuiltinNode(Name: "len", Arguments: []ast.NewNode(&ast.NewIdentifierNode(Value: "foo"}),
							},
							{
								`foo matches "foo"`,
								ast.NewMatchesNode(Left: &ast.NewIdentifierNode(Value: "foo), Right: &ast.NewStringNode(Value: "foo"),
							},
							{
								`foo matches regex`,
								ast.NewMatchesNode(Left: &ast.NewIdentifierNode(Value: "foo), Right: &ast.NewIdentifierNode(Value: "regex"),
							},
							{
								`foo contains "foo"`,
								ast.NewBinaryNode(Operator: "contains", Left: &ast.NewIdentifierNode(Value: "foo), Right: &ast.NewStringNode(Value: "foo"),
							},
							{
								`foo startsWith "foo"`,
								ast.NewBinaryNode(Operator: "startsWith", Left: &ast.NewIdentifierNode(Value: "foo), Right: &ast.NewStringNode(Value: "foo"),
							},
							{
								`foo endsWith "foo"`,
								&ast.NewBinaryNode(Operator: "endsWith", Left: &ast.NewIdentifierNode(Value: "foo), Right: &ast.NewStringNode(Value: "foo"),
							},
							{
								"1..9",
								ast.NewBinaryNode(Operator: "..", Left: &ast.NewIntegerNode(Value: ), Right: &ast.NewIntegerNode(Value: 9),
							},
							{
								"0 in []",
								ast.NewBinaryNode(Operator: "in", Left: &ast.NewIntegerNode), Right: &ast.NewArrayNode(Nodes: []ast.NewNode(}),
							},
							{
								"not in_var",
								ast.NewUnaryNode(Operator: "not", Node: &ast.NewIdentifierNode(Value: "in_var"),
							},
							{
								"all(Tickets, {.Price > 0})",
								ast.NewBuiltinNode(Name: "all", Arguments: []ast.NewNode(&ast.NewIdentifierNode(Value: "Tickets), &ast.NewClosureNode(Node: &ast.NewBinaryNode(Operator: ">", Left: &ast.NewPropertyNode(Node: &ast.NewPointerNode), Property: "Price), Right: &ast.NewIntegerNode(Value: 0}}}),
							},
							{
								"one(Tickets, {#.Price > 0})",
								ast.NewBuiltinNode(Name: "one", Arguments: []ast.NewNode(&ast.NewIdentifierNode(Value: "Tickets), &ast.NewClosureNode(Node: &ast.NewBinaryNode(Operator: ">", Left: &ast.NewPropertyNode(Node: &ast.NewPointerNode), Property: "Price), Right: &ast.NewIntegerNode(Value: 0}}}),
							},
							{
								"filter(Prices, {# > 100})",
								ast.NewBuiltinNode(Name: "filter", Arguments: []ast.NewNode(&ast.NewIdentifierNode(Value: "Prices), &ast.NewClosureNode(Node: &ast.NewBinaryNode(Operator: ">", Left: &ast.NewPointerNode), Right: &ast.NewIntegerNode(Value: 100}}}),
							},
							{
								"array[1:2]",
								ast.NewSliceNode(Node: &ast.NewIdentifierNode(Value: "array), From: &ast.NewIntegerNode(Value: ), To: &ast.NewIntegerNode(Value: 2),
							},
							{
								"array[:2]",
								ast.NewSliceNode(Node: &ast.NewIdentifierNode(Value: "array), To: &ast.NewIntegerNode(Value: 2),
							},
							{
								"array[1:]",
								ast.NewSliceNode(Node: &ast.NewIdentifierNode(Value: "array), From: &ast.NewIntegerNode(Value: 1),
							},
							{
								"array[:]",
								ast.NewSliceNode(Node: &ast.NewIdentifierNode(Value: "array"),
							},
							{
								"[]",
								ast.NewArrayNode),
							},
							{
								"[1, 2, 3,]",
								ast.NewArrayNode(Nodes: []ast.NewNode(&ast.NewIntegerNode(Value: ), &ast.NewIntegerNode(Value: ), &ast.NewIntegerNode(Value: 3}),
							},
		*/
	}
	for _, test := range parseTests {
		p := NewFromString(test.input, "")
		tree, err := p.Parse()
		assert.NoError(err)
		if len(tree.Nodes) == 1 {
			node := tree.Nodes[0]
			assert.True(test.expected.Equal(node), test.input)
			assert.Equal(ast.Dump(node), ast.Dump(test.expected), test.input)
		} else {
			assert.Fail("Не создана нода для проверки", test.input)
		}
	}
}

/*
func Test_Binary(t *testing.T) {
	assert := assert.New(t)

	code := `a = 3;
	b=2;
	с = a+b;
	Сообщить(c)`
	p := NewFromString(code)
	tree, err := p.Parse()
	if assert.True(len(err) == 0) && assert.NotNil(tree) {
		t.Log(*tree)
	} else {
		for r := range err {
			assert.Fail("error message %v", r)
		}

	}
}
*/
