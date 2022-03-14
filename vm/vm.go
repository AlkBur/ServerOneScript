package vm

type VM struct {
	stack     Stack
	registers [2]Object
}

func NewVM() *VM {
	return &VM{
		stack: make(Stack, 0, 2),
	}
}

func Run(program *Program, mainFn string, args ...interface{}) *Result {
	if program == nil {
		res := &Result{}
		res.error("program is nil")
		return res
	}
	vm := NewVM()
	return vm.Run(program, mainFn, args...)
}

func (vm *VM) Run(program *Program, mainFn string, args ...interface{}) *Result {
	defer func() {
		if r := recover(); r != nil {
			program.error("%v", r)
		}
	}()

	vm.stack = vm.stack[:0]
	for i := 0; i < len(vm.registers); i++ {
		vm.registers[i] = nil
	}

	fn, ok := program.functions.data[mainFn]
	if !ok {
		program.error("Main function not find %s", mainFn)
		return program.result
	}

	for _, instruction := range fn.instructions {
		switch instruction.opcode {
		case OpPush:
			vm.stack.Push(vm.registers[0])
		case OpPopA:
			vm.registers[0] = vm.stack.Pop()
		case OpPopB:
			vm.registers[1] = vm.stack.Pop()
		case OpPopNop:
			vm.stack.Pop()
		case OpLoad:
			obj := program.loadObject(fn, instruction.object)
			vm.stack.Push(obj)
		case OpSet:
			program.setVariantObject(fn, instruction.object, vm.registers[0])
		case OpReturn:
			program.result.Obj = vm.registers[0]
		case OpAdd:
			result, err := AddObject(vm.registers[0], vm.registers[1])
			if err != nil {
				program.error("Ошибка при перации сложения %v", err)
			}
			vm.registers[0] = result
			vm.registers[1] = nil
		case OpSub:
			result, err := SubObject(vm.registers[0], vm.registers[1])
			if err != nil {
				program.error("Ошибка при перации сложения %v", err)
			}
			vm.registers[0] = result
			vm.registers[1] = nil
		case OpMul:
			result, err := MulObject(vm.registers[0], vm.registers[1])
			if err != nil {
				program.error("Ошибка при перации сложения %v", err)
			}
			vm.registers[0] = result
			vm.registers[1] = nil

		case OpDiv:
			result, err := DivObject(vm.registers[0], vm.registers[1])
			if err != nil {
				program.error("Ошибка при перации сложения %v", err)
			}
			vm.registers[0] = result
			vm.registers[1] = nil

		default:
			program.error("Check opcode %v", instruction.opcode)
			return program.result
		}
		if program.IsError() {
			break
		}
	}

	return program.result
}
