package parser

import (
	"testing"

	"github.com/rahuldshetty/talion/ast"
	"github.com/rahuldshetty/talion/lexer"
)

func TestVarStatements(t *testing.T){
	input := `
	var x = 5;
	var y = 10 ;
	var foo = 83838;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements together. got=%d", len(program.Statements))
	}
	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i,tt := range tests{
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifier){
			return 
		}
	}
}


func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenListeral not 'var'. got=%q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not '%s'. got=%s", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("varStmt.Name not '%s'. got=%s", name, varStmt.Name)
		return false
	}
	return true;
}