package ast

import "github.com/AlkBur/ServerOneScript/runes"

type Tree struct {
	Source *runes.Source
	Nodes  []Node
}

func New(source *runes.Source) *Tree {
	return &Tree{
		Nodes:  make([]Node, 0),
		Source: source,
	}
}

func (t *Tree) AddNode(n Node) {
	t.Nodes = append(t.Nodes, n)
}
