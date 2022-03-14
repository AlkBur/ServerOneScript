package vm

type Stack []Object

func (s *Stack) Push(v Object) {
	*s = append(*s, v)
}

func (s *Stack) Pop() Object {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
