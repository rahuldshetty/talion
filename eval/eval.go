package eval

import (
	"fmt"

	"github.com/rahuldshetty/talion/ast"
	"github.com/rahuldshetty/talion/object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type){
		// Statements
		case *ast.Program:
			return evalProgram(node, env)

		case *ast.ExpressionStatement:
			return Eval(node.Expression, env)

		case *ast.BlockStatement:
			return evalBlockStatement(node, env)

		case *ast.IfExpression:
			return evalIfExpression(node, env)

		case *ast.Identifier:
			return evalIdentifer(node, env)
		
		case *ast.VarStatement:
			val := Eval(node.Value, env)
			if isError(val){
				return val
			}
			env.Set(node.Name.Value, val)

		case *ast.AssignStatement:
			return evalAssignmentExpression(node, env)

		case *ast.ReturnStatement:
			val := Eval(node.ReturnValue, env)
			if isError(val){
				return val
			}
			return &object.ReturnValue{Value: val}

		// List
		case *ast.ListLiteral:
			elements := evalExpressions(node.Elements, env)
			if len(elements) == 1 && isError(elements[0]){
				return elements[0]
			}
			return &object.List{Elements: elements}

		case *ast.IndexExpression:
			left := Eval(node.Left, env)
			if isError(left){
				return left
			}
			index := Eval(node.Index, env)
			if isError(index){
				return index
			}
			return evalIndexExpression(left, index)

		// Function Definition
		case *ast.FunctionLiteral:
			params := node.Parameters
			body := node.Body
			return &object.Function{Parameters: params, Env: env, Body: body}

		// Function Call
		case *ast.CallExpression:
			function := Eval(node.Function, env)
			if isError(function){
				return function
			}
			args := evalExpressions(node.Arguments, env)
			if len(args) == 1 && isError(args[0]){
				return args[0]
			}
			// Invoke Function Call 
			return applyFunction(function, args)

		// Expression
		case *ast.IntegerLiteral: 
			return &object.Integer{Value: node.Value}

		case *ast.StringLiteral:
			return &object.String{Value: node.Value}

		case *ast.Boolean: 
			return nativeBoolToBooleanObject(node.Value)

		// Operator Expression
		case *ast.InfixExpression:
			left := Eval(node.Left, env)
			if isError(left){
				return left
			}
			right := Eval(node.Right, env)
			if isError(right){
				return right
			}
			return evalInfixExpression(node.Operator, left, right)
		
		case *ast.PrefixExpression:
			right := Eval(node.Right, env)
			if isError(right){
				return right
			}
			return evalPrefixExpression(node.Operator, right)

	}
	return nil
}

// To handle nested loop need to make generic eval program
func evalProgram(program *ast.Program,  env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements{
		result = Eval(statement, env)		

		switch result := result.(type){
			case *object.ReturnValue: return result.Value
			case *object.Error: return result
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement,  env *object.Environment) object.Object{
	var result object.Object

	for _, statement := range block.Statements{
		result = Eval(statement, env)		

		if result != nil{
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result

}

// If - else Expression
func evalIfExpression(ie *ast.IfExpression,  env *object.Environment) object.Object{
	condition := Eval(ie.Condition, env)

	if isError(condition){
		return condition
	}

	if isTruthy(condition){
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil{
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

// Null & False fails the condition
func isTruthy(obj object.Object) bool {
	switch obj {
		case NULL: return false
		case TRUE: return true
		case FALSE: return false
		default: return true 
	}
}

func applyFunction(fn object.Object, args []object.Object) object.Object{
	switch fn := fn.(type){
		case *object.Function:
			extendedEnv := extendFunctionEnv(fn, args)
			evaluated := Eval(fn.Body, extendedEnv)
			return unwrapReturnValue(evaluated)
		
		case *object.Builtin:
			return fn.Fn(args...)

		default:
			return newError("not a function: %s", fn.Type())
	}	
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment{
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters{
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object{
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps{
		evaluated := Eval(e, env)
		if isError(evaluated){
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalAssignmentExpression(assignment *ast.AssignStatement, env *object.Environment) object.Object{
	evaluated := Eval(assignment.Value, env)
	if isError(evaluated){
		return evaluated
	}

	switch assignment.Operator{
		// TODO: operator assignment

		case "=":
			env.Set(assignment.Name.String(), evaluated)
	}

	// return evaluated
	return nil
}


// Index evaluation
func evalIndexExpression(left, index object.Object) object.Object{
	switch{
		case left.Type() == object.LIST_OBJ && index.Type() == object.INTEGER_OBJ:
			return evalListIndexExpression(left, index)
		case left.Type() == object.STRING_OBJ && index.Type() == object.INTEGER_OBJ:
			return evalStringIndexExpression(left, index)
		default:
			return newError("Index operator not supported: %s", left.Type())
	}
}

// Fetch index element from left list
func evalListIndexExpression(list, index object.Object) object.Object {
	listObject := list.(*object.List)
	idx := index.(*object.Integer).Value
	max := int64(len(listObject.Elements))

	// zero-indexing: 0 to len(list) - 1
	if idx >= 0 && idx < max{
		return listObject.Elements[idx]
	}

	// negative-indexing: -1 to -len(list)
	if idx >= -max && idx <= -1{
		return listObject.Elements[max + idx]
	}
	
	return NULL
}


// Fetch index element from left list
func evalStringIndexExpression(str, index object.Object) object.Object {
	str_val := str.(*object.String).Value
	idx := index.(*object.Integer).Value
	max := int64(len(str_val))

	// zero-indexing: 0 to len(list) - 1
	if idx >= 0 && idx < max{
		return &object.String{Value: string(str_val[idx]) }
	}

	// negative-indexing: -1 to -len(list)
	if idx >= -max && idx <= -1{
		return  &object.String{Value: string(str_val[max + idx]) }
	}
	
	return NULL
}


func evalIdentifer(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok{
		return builtin
	}
	
	return newError("Identifier not found: " + node.Value)
}

// Unary Operator Switching
func evalPrefixExpression(operator string, right object.Object) object.Object{
	switch operator{
		case "!": 
			return evalNOTOperatorExpression(right)
		case "-":
			return evalMinusPrefixOperatorExpression(right)
		default:
			return newError("Unknown operator: %s%s", operator, right.Type())
	}
}

// Binary Operator Switching
func evalInfixExpression(operator string, left, right object.Object) object.Object{
	switch{
		case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
			return evalIntegerInfixExpression(operator, left, right)
		

		case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
			return evalStringInfixExpression(operator, left, right)
		
		// TODO: String comparisor with == and !=

		// pointer comparison here to check for equality between booleans
		// if left & right is boolean type then they are referenced already
		// so equality is checked on pointer is pointing to correct boolean types
		case operator == "==":
			return nativeBoolToBooleanObject(left == right)
		case operator == "!=":
			return nativeBoolToBooleanObject(left != right)
		
		case left.Type() != right.Type():
			return newError("Type mismatch: %s %s %s", left.Type() ,operator, right.Type())

		default: return newError("Unknown operator: %s %s %s", left.Type() ,operator, right.Type())
	}
}

// NOT Operator Logic
func evalNOTOperatorExpression(right object.Object) object.Object {
	switch right{
		case TRUE: return FALSE
		case FALSE: return TRUE
		case NULL: return TRUE
		default: return FALSE
	}
}

// Binary Operator Evaluation - Integer & Integer
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator{
		case "+":  return &object.Integer{Value: leftVal + rightVal}
		case "-":  return &object.Integer{Value: leftVal - rightVal}
		case "*":  return &object.Integer{Value: leftVal * rightVal}
		case "/":  return &object.Integer{Value: leftVal / rightVal}

		case "<":  return nativeBoolToBooleanObject(leftVal < rightVal)
		case "<=":  return nativeBoolToBooleanObject(leftVal <= rightVal)
		case ">":  return nativeBoolToBooleanObject(leftVal > rightVal)
		case ">=":  return nativeBoolToBooleanObject(leftVal >= rightVal)
		case "==":  return nativeBoolToBooleanObject(leftVal == rightVal)
		case "!=":  return nativeBoolToBooleanObject(leftVal != rightVal)
		
		default: return newError("Unknown operator: %s %s %s", left.Type() ,operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object{
	if operator != "+"{
		return newError("Unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object{
	if right.Type() != object.INTEGER_OBJ{
		return newError("Unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean{
	if input{
		return TRUE
	}
	return FALSE
}

func newError(format string, a ...interface{}) *object.Error{
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil{
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}