package eval

import (
	"github.com/rahuldshetty/talion/object"
)

var builtins = map[string]*object.Builtin{
	// len of String object
	"len": {
		Fn: func(args ...object.Object) object.Object{
			if len(args) != 1{
				return newError("Wrong number of arguments. got=%d, want=1", len(args))
			}
			
			switch arg := args[0].(type){

				case *object.String:
					return &object.Integer{Value: int64(len(arg.Value))}

				default: 
					return newError("Argument to `len` not supported, got %s", args[0].Type())
			}

		},
	},
}