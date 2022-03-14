package vm

type Variables struct {
	data map[string]*Variable
}
type Variable struct {
	name  string
	value Object
}

func NewVariable(name string) *Variable {
	return &Variable{
		name:  name,
		value: nil,
	}
}

func NewVariables() *Variables {
	return &Variables{
		data: make(map[string]*Variable),
	}
}

func (vrs *Variables) Add(vr *Variable) {
	vrs.data[vr.name] = vr
}
