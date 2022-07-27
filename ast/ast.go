package ast

import "github.com/rahuldshetty/talion/token"

//Node will have 3 fields: identifier, expression, token
type Node interface {
	TokenLiteral() string
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


// Return 
// return <expression>;
type ReturnStatement struct {
	Token token.Token // token.RETURN
	ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Identifier
type Identifier struct{
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }