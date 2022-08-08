package compiler

import (
	"github.com/rahuldshetty/talion/ast"
	"github.com/rahuldshetty/talion/code"
	"github.com/rahuldshetty/talion/object"
)

type Compiler struct{
	instructions code.Instructions
	constants []object.Object
}

type Bytecode struct{
	Instructions code.Instructions
	Constants []object.Object
}

func New() *Compiler{
	return &Compiler{
		instructions: code.Instructions{},
		constants: []object.Object{},
	}
}


func (c *Compiler) Compile(node ast.Node) error{
	return nil
}

func (c *Compiler) Bytecode() *Bytecode{
	return &Bytecode{
		Instructions: c.instructions,
		Constants: c.constants,
	}
}