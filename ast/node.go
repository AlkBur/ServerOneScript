package ast

import (
	"errors"
	"github.com/AlkBur/ServerOneScript/runes"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type Node interface {
	Location() runes.Location
	SetLocation(runes.Location)
	Type() Kind
	SetType(Kind)
	Equal(node Node) bool
}

//BaseNode - базовый класс
type BaseNode struct {
	location runes.Location
	nodeType Kind
}

func NewBaseNode(t Kind, l runes.Location) *BaseNode {
	return &BaseNode{
		nodeType: t,
		location: l,
	}
}

func (n *BaseNode) Location() runes.Location {
	return n.location
}

func (n *BaseNode) SetLocation(l runes.Location) {
	n.location = l
}

func (n *BaseNode) Type() Kind {
	return n.nodeType
}

func (n *BaseNode) SetType(t Kind) {
	n.nodeType = t
}

func (n *BaseNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if _, ok := node.(*BaseNode); ok {
		return true
	}
	return false
}

//NumberNode - Класс узла значения числа.
type NumberNode struct {
	BaseNode
	Value decimal.Decimal
}

func NewNumberNode(val string) (*NumberNode, error) {
	result := &NumberNode{
		BaseNode: BaseNode{
			nodeType: Number,
		},
	}
	err := result.SetValue(val)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *NumberNode) SetValue(val string) error {
	dec, err := decimal.NewFromString(val)
	if err != nil {
		return err
	}
	n.Value = dec
	return nil
}

func (n *NumberNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*NumberNode); ok {
		return m.Value.Equal(n.Value)
	}
	return false
}

//DateNode - Класс узла значения даты.
type DateNode struct {
	BaseNode
	Value time.Time
}

func NewDateNode(val string) (*DateNode, error) {
	result := &DateNode{
		BaseNode: BaseNode{
			nodeType: Date,
		},
	}
	err := result.SetValue(val)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *DateNode) SetValue(val string) error {
	t, err := time.Parse("20060102150405", val)
	if err != nil {
		return err
	}
	n.Value = t
	return nil
}

func (n *DateNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*DateNode); ok {
		return m.Value.Equal(n.Value)
	}
	return false
}

//BoolNode - Класс узла значения булево.
type BoolNode struct {
	BaseNode
	Value bool
}

func NewBoolNode(val string) (*BoolNode, error) {
	result := &BoolNode{
		BaseNode: BaseNode{
			nodeType: Bool,
		},
	}
	err := result.SetValue(val)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *BoolNode) SetValue(val string) error {
	if strings.EqualFold(val, "Истина") || strings.EqualFold(val, "True") {
		n.Value = true
	} else if strings.EqualFold(val, "Ложь") || strings.EqualFold(val, "False") {
		n.Value = false
	} else {
		return errors.New("Ошибка boolean значения: " + val)
	}
	return nil
}

func (n *BoolNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*BoolNode); ok {
		return m.Value == n.Value
	}
	return false
}

// StringNode - Класс узла значения строки.
type StringNode struct {
	BaseNode
	Value string
}

func NewStringNode(value string) *StringNode {
	return &StringNode{
		BaseNode: BaseNode{
			nodeType: String,
		},
		Value: value,
	}
}

func (n *StringNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*StringNode); ok {
		return m.Value == n.Value
	}
	return false
}

// BinaryNode - Класс узла формулы.
type BinaryNode struct {
	BaseNode
	Operator string
	Left     Node
	Right    Node
}

func NewBinaryNode(operator string, left, right Node) *BinaryNode {
	return &BinaryNode{
		BaseNode: BaseNode{
			nodeType: BinaryOperation,
		},
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (n *BinaryNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*BinaryNode); ok {
		if m.Operator != n.Operator {
			return false
		}
		isEmpty := isNil(isNil(n.Left))
		if isNil(m.Left) != isEmpty {
			return false
		} else if !isEmpty && !n.Left.Equal(m.Left) {
			return false
		}
		isEmpty = isNil(isNil(n.Right))
		if isNil(m.Right) != isEmpty {
			return false
		} else if !isEmpty && !n.Right.Equal(m.Right) {
			return false
		}
		return true
	}
	return false
}

// VariableNode - Класс определения переменной.
type VariableNode struct {
	BaseNode
	Name     string
	Value    Node
	IsExport bool
	IsRef    bool
}

func NewVariableNode(name string, val Node) *VariableNode {
	return &VariableNode{
		BaseNode: BaseNode{
			nodeType: Variable,
		},
		Name:  name,
		Value: val,
	}
}

func (n *VariableNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*VariableNode); ok {
		if m.Name != n.Name {
			return false
		} else if m.IsExport != n.IsExport {
			return false
		} else if m.IsRef != n.IsRef {
			return false
		}

		isEmpty := isNil(n.Value)
		if isNil(m.Value) != isEmpty {
			return false
		} else if !isEmpty && !n.Value.Equal(m.Value) {
			return false
		}
		return true
	}
	return false
}

// ReturnNode - Класс определения переменной.
type ReturnNode struct {
	BaseNode
	Value Node
}

func NewReturnNode(val Node) *ReturnNode {
	return &ReturnNode{
		BaseNode: BaseNode{
			nodeType: Variable,
		},
		Value: val,
	}
}

func (n *ReturnNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*ReturnNode); ok {
		isEmpty := isNil(n.Value)
		if isNil(m.Value) != isEmpty {
			return false
		} else if !isEmpty && !n.Value.Equal(m.Value) {
			return false
		}
		return true
	}
	return false
}

// LiteralNode - Класс переменной.
type LiteralNode struct {
	BaseNode
	Name  string
	IsRef bool
}

func NewLiteralNode(name string) *LiteralNode {

	return &LiteralNode{
		BaseNode: BaseNode{
			nodeType: Literal,
		},
		Name: name,
	}
}

func (n *LiteralNode) Equal(node Node) bool {
	if !n.nodeType.Equal(node.Type()) {
		return false
	}
	if !n.location.Equal(node.Location()) {
		return false
	}
	if m, ok := node.(*LiteralNode); ok {
		if m.Name != n.Name {
			return false
		} else if m.IsRef != n.IsRef {
			return false
		}
		return true
	}
	return false
}

// FunctionCallNode - Класс вызова.
type FunctionCallNode struct {
	BaseNode
	Name   string
	Args   []Node
	Export bool
}

func NewFunctionCallNode(name string, args []Node) *FunctionCallNode {

	return &FunctionCallNode{
		BaseNode: BaseNode{
			nodeType: FunctionCall,
		},
		Name: name,
		Args: args,
	}
}

func (n *FunctionCallNode) Equal(node Node) bool {
	if m, ok := node.(*FunctionCallNode); ok {
		if !m.BaseNode.Equal(&n.BaseNode) {
			return false
		} else if m.Name != n.Name {
			return false
		} else if len(n.Args) != len(m.Args) {
			return false
		} else if n.Export != m.Export {
			return false
		}
		for i := 0; i < len(n.Args); i++ {
			if !n.Args[i].Equal(m.Args[i]) {
				return false
			}
		}
		return true
	}
	return false
}
