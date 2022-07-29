package eval

import (
	"github.com/rahuldshetty/talion/ast"
	"github.com/rahuldshetty/talion/object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type){
		// Statements
		case *ast.Program:
			return evalStatements(node.Statements)

		case *ast.ExpressionStatement:
			return Eval(node.Expression)

		case *ast.BlockStatement:
			return evalStatements(node.Statements)

		case *ast.IfExpression:
			return evalIfExpression(node)

		// Expression
		case *ast.IntegerLiteral: 
			return &object.Integer{Value: node.Value}
		case *ast.Boolean: 
			return nativeBoolToBooleanObject(node.Value)

		// Operator Expression
		case *ast.InfixExpression:
			left := Eval(node.Left)
			right := Eval(node.Right)
			return evalInfixExpression(node.Operator, left, right)
		
		case *ast.PrefixExpression:
			right := Eval(node.Right)
			return evalPrefixExpression(node.Operator, right)

	}
	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements{
		result = Eval(statement)		
	}

	return result
}

// If - else Expression
func evalIfExpression(ie *ast.IfExpression) object.Object{
	condition := Eval(ie.Condition)

	if isTruthy(condition){
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil{
		return Eval(ie.Alternative)
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

// Unary Operator Switching
func evalPrefixExpression(operator string, right object.Object) object.Object{
	switch operator{
		case "!": 
			return evalNOTOperatorExpression(right)
		case "-":
			return evalMinusPrefixOperatorExpression(right)
		default:
			return NULL
	}
}

// Binary Operator Switching
func evalInfixExpression(operator string, left, right object.Object) object.Object{
	switch{
		case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
			return evalIntegerInfixExpression(operator, left, right)
		
		// pointer comparison here to check for equality between booleans
		// if left & right is boolean type then they are referenced already
		// so equality is checked on pointer is pointing to correct boolean types
		case operator == "==":
			return nativeBoolToBooleanObject(left == right)
		case operator == "!=":
			return nativeBoolToBooleanObject(left != right)
		

		default: return NULL
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
		
		default: return NULL
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object{
	if right.Type() != object.INTEGER_OBJ{
		return NULL
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