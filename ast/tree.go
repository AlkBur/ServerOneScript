package ast

type Tree struct {
	Nodes []Node
}

func New() *Tree {
	return &Tree{
		Nodes: make([]Node, 0, 100),
	}
}

func (t *Tree) AddNode(n Node) {
	t.Nodes = append(t.Nodes, n)
}
