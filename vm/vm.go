package vm

import (
	"fmt"

	"github.com/rahuldshetty/talion/code"
	"github.com/rahuldshetty/talion/compiler"
	"github.com/rahuldshetty/talion/object"
)

const StackSize = 2048

var True = &object.Boolean{Value: true}
var False = &object.Boolean{Value: false}

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
			case code.OpPop:
				vm.pop()

			case code.OpConstant:
				constIndex := code.ReadUint16(vm.instructions[ip+1:])
				ip += 2
				err := vm.push(vm.constants[constIndex])
				if err != nil{
					return err
				}

			case code.OpAdd, code.OpSub, code.OpDiv, code.OpMul:
				err := vm.executeBinaryOperation(op)
				if err != nil{
					return err
				}

			case code.OpTrue:
				err := vm.push(True)
				if err != nil{
					return err
				}

			case code.OpFalse:
				err := vm.push(False)
				if err != nil{
					return err
				}

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

func (vm *VM) LastPoppedStackElem() object.Object{
	return vm.stack[vm.sp]
}

func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	leftType := left.Type()
	rightType := right.Type()

	if leftType == object.INTEGER_OBJ && rightType == object.INTEGER_OBJ{
		return vm.executeBinaryIntegerOperation(op, left, right)
	}
	return fmt.Errorf("unsupported types of binary operation: %s %s", leftType, rightType)
}

func (vm *VM) executeBinaryIntegerOperation(op code.Opcode, left, right object.Object) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	var result int64

	switch op{
		case code.OpAdd:
			result = leftValue + rightValue
		case code.OpSub:
			result = leftValue - rightValue
		case code.OpMul:
			result = leftValue * rightValue
		case code.OpDiv:
			result = leftValue / rightValue
		default:
			return fmt.Errorf("unknown integer operator: %d", op)
	}
	vm.push(&object.Integer{Value: result})

	return nil
}