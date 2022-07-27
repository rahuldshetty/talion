package ast

import (
	"bytes"

	"github.com/rahuldshetty/talion/token"
)

//Node will have 3 fields: identifier, expression, token
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program node - root node of every AST produces by parser
// Every program contains sequence of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// STATEMENTS

// var <identifier> = <expression> ;
// VAR
type VarStatement struct{
	Token token.Token // token.VAR token
	Name *Identifier // hold the identifier of the binding
	Value Expression // expression that produces the value
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	out.WriteString(" = ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Return 
// return <expression>;
type ReturnStatement struct {
	Token token.Token // token.RETURN
	ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if  rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

// Identifier
type Identifier struct{
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

type ExpressionStatement struct{
	Token token.Token // the first token of the expression
	Expression Expression
}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if  es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral
type IntegerLiteral struct{
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }
