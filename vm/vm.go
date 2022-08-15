package vm

import (
	"fmt"

	"github.com/rahuldshetty/talion/code"
	"github.com/rahuldshetty/talion/compiler"
	"github.com/rahuldshetty/talion/object"
)

const StackSize = 2048

type VM struct{
	constants []object.Object
	instructions code.Instructions

	stack []object.Object
	sp int // point to the next value. top of stack is stack[sp-1]
}

func New(bytecode *compiler.Bytecode) *VM{
	return &VM{
		instructions: bytecode.Instructions,
		constants: bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp: 0,
	}
}


func (vm *VM) StackTop() object.Object{
	if vm.sp == 0{
		return nil
	}
	return vm.stack[vm.sp - 1]
}

func (vm *VM) Run() error {
	// ip - instruction pointer
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op{

			case code.OpConstant:
				constIndex := code.ReadUint16(vm.instructions[ip+1:])
				ip += 2
				err := vm.push(vm.constants[constIndex])
				if err != nil{
					return err
				}

			case code.OpAdd:
				right := vm.pop()
				left := vm.pop()

				leftValue := left.(*object.Integer).Value
				rightValue := right.(*object.Integer).Value

				result := leftValue + rightValue
				vm.push(&object.Integer{Value: result})
		}
	}

	return nil
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize{
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

func (vm *VM) pop() object.Object{
	obj := vm.stack[vm.sp - 1]
	vm.sp--
	return obj
}