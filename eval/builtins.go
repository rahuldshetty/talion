package eval

import (
	"fmt"

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

				case *object.List:
					return &object.Integer{Value: int64(len(arg.Elements))}

				case *object.Hash:
					return &object.Integer{Value: int64(len(arg.Pairs))}

				default: 
					return newError("Argument to `len` not supported, got %s", args[0].Type())
			}

		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2{
				return newError("Wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.LIST_OBJ{
				return newError("Argument to push must be LIST, got %s", args[0].Type())
			}
			arr := args[0].(*object.List)
			arr.Elements = append(arr.Elements, args[1])
			return nil
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object{
			for _, arg := range args{
				fmt.Println(arg.Inspect())
			}
			return nil
		},
	},
}