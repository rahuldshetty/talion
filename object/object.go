package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
)

// Base Object to represent internal system of interprer - used during runtime evaluation
type Object interface {
	Type() ObjectType
	Inspect() string
}


// Integer Data type 
type Integer struct{
	Value int64
}

func (i *Integer) Inspect() string{ return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean Data Type
type Boolean struct{
	Value bool
}

func (b *Boolean) Inspect() string{ return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// NULL Data Type
type Null struct{}
func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

