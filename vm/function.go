package vm

type Functions struct {
	data map[string]*Function
}

func NewFunctions() *Functions {
	return &Functions{
		data: make(map[string]*Function),
	}
}

func (fns *Functions) Get(name string) *Function {
	return fns.data[name]
}

func (fns *Functions) Add(fn *Function) {
	fns.data[fn.name] = fn
}

type Function struct {
	name         string
	variables    *Variables
	instructions []Instruction
}

func NewFunction(name string) *Function {
	return &Function{
		name:         name,
		variables:    NewVariables(),
		instructions: make([]Instruction, 0),
	}
}

func (fn *Function) AddVariable(vr *Variable) {
	fn.variables.Add(vr)
}

func (fn *Function) AddInstruction(i Instruction) {
	fn.instructions = append(fn.instructions, i)
}
