package eval

import (
	"testing"

	"github.com/rahuldshetty/talion/lexer"
	"github.com/rahuldshetty/talion/object"
	"github.com/rahuldshetty/talion/parser"
)

func TestEvalIntegerExpression(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},

		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests{
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object{
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok{
		t.Errorf("object is not Integer. got=%T (%+v))", obj, obj)
		return false
	}

	if result.Value != expected{
		t.Errorf("object has wrong value. got=%d, expected=%d", result.Value, expected)
		return false
	}
	return true
}


func TestEvalBooleanExpression(t *testing.T){
	tests := []struct{
		input string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}
	for _,tt := range tests{
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok{
		t.Errorf("object is not Boolean. got=%T", obj)
		return false
	}
	if result.Value != expected{
		t.Errorf("object has wrong value. got=%t, expected=%t", result.Value, expected)
		return false
	}
	return true
}

func TestNOTOperator(t *testing.T){
	tests := []struct{
		input string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},

		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 <= 1", true},
		{"1 >= 1", true},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 <= 2", true},
		{"900 >= -4", true},

		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}
	for _, tt := range tests{
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T){
	tests:=[]struct{
		input string
		expected interface{}
	} {
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}
	for _, tt := range tests{
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2*5; 9;", 10},
		{"9; return 2*5; 9;", 10},
		{"if(10>1){ if(10>1){ return 10; } return 1; }", 10},
	}
	for _, tt := range tests{
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}


func TestErrorHandling(t *testing.T) {
	tests := []struct {
	input string
	expectedMessage string
	}{
		{
			"5 + true;",
			"Type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"Type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"Unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}
				return 1;
			}
			`,
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"Identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"Unknown operator: STRING - STRING",
		},
		{
			`{"name":"hello"}[fn(x){x}]`,
			"Type not support as hash key: FUNCTION",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}
		
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestVarStatements(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"var a = 5; a;", 5},
		{"var a = 5 * 5; a;", 25},
		{"var a = 5; var b=a; b;", 5},
		{"var a = 5; var b=a; var c = a + b + 5; c", 15},
	}
	for _, tt := range tests{
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestAssignExpressions(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"a=5;a;", 5},
		{"var a = 5; a = 10; a;", 10},
		{"var a = 5; a = a + 10; a;", 15},
		{"var a = [1, 2, 3]; a = 1; a;", 1},
		{"var a = [1, 2, 3]; b=a; b[2];", 3},
	}
	for _, tt := range tests{
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestFunctionObject(t *testing.T){
	input := "fn(x) { x + 2; };"
	evaluated := testEval(input)

	fn, ok := evaluated.(*object.Function)
	if !ok{
		t.Fatalf("object is not Function. got=%T (%+v)" , evaluated, evaluated)
	}

	if len(fn.Parameters) != 1{
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x"{
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody{
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"var identity = fn(x) { x; }; identity(5);", 5},
		{"var identity = fn(x) { return x; }; identity(5);", 5},
		{"var double = fn(x) { x * 2; }; double(5);", 10},
		{"var add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"var add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},	
		{"fn (x, y) { if (x > y) { return x; } else { return y; }; }(5, 15)", 15},	
	}

	for _, tt := range tests{
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestClosures(t *testing.T){
	input := `
	var newAdder = fn(x){
		return fn(y){ x + y };
	};
	var addTwo = newAdder(2);
	addTwo(2);
	`
	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T){
	input := `"Hello World";`

	evaluated := testEval(input)
	str,ok := evaluated.(*object.String)
	
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World"{
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringIndexing(t *testing.T){
	input := `"Hello World"[0];`

	evaluated := testEval(input)
	str,ok := evaluated.(*object.String)
	
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "H"{
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}


func TestStringConcatenation(t *testing.T){
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)

	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!"{
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T){
	tests := []struct{
		input string
		exprected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "Argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "Wrong number of arguments. got=2, want=1"},
		{`len([1, 2, 3, 4])`, 4},
		{`len([])`, 0},
		{`len(["hello", "world"])`, 2},
		{`var ls = []; push(ls, 1); ls[0];`, 1},
	}

	for _, tt := range tests{
		evaluated := testEval(tt.input)

		switch expected := tt.exprected.(type){
			case int:
				testIntegerObject(t, evaluated, int64(expected))
			case string:
				errObj, ok := evaluated.(*object.Error)
				if !ok{
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}
				if errObj.Message != expected {
					t.Errorf("Wrong error message. expected=%q got=%q", expected, errObj.Message)
				}
		}

	}
}

func TestListLiterals(t *testing.T){
	input := "[1, 2*2, 3+3]"
	evaluated := testEval(input)

	result, ok := evaluated.(*object.List)
	if !ok{
		t.Fatalf("object is not List Type. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3{
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestIndexExpressions(t *testing.T){
	inputs := []struct{
		input string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"var i = 0; [1][i]",
			1,
		},
		{
			"[1, 2, 3][1 + 1]",
			3,
		},
		{
			"var myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"var myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"var myArray = [1, 2, 3]; var i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			3,
		},
		{
			"[1, 2, 3][-2]",
			2,
		},
		{
			"[1, 2, 3][-3]",
			1,
		},
		{
			"[1, 2, 3][-4]",
			nil,
		},
		{
			`{"1":1, "2":2, "3": 3}["1"]`,
			1,
		},
		{
			`{"1":1, "2":2, "3": 3}["2"]`,
			2,
		},
		{
			`{"1":1, "2":2, "3": 3}["3"]`,
			3,
		},
		{
			`{"1":1, "2":2, "foo": 3}["foo"]`,
			3,
		},
		{
			`{"bar":1, "2":2, "foo": 3}["bar"]`,
			1,
		},
	}

	for _, tt := range inputs{
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}

}


func TestHashLiterals(t *testing.T) {
	input := `two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}
	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey(): 1,
		(&object.String{Value: "two"}).HashKey(): 2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey(): 4,
		TRUE.HashKey(): 5,
		FALSE.HashKey(): 6,
	}
	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}
	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}	
		testIntegerObject(t, pair.Value, expectedValue)
	}
}