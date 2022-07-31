package ast

import (
	"bytes"
	"strings"

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


// Assignment Statement
// Operator stores string info - '', '+', '-', '/' operation
type AssignStatement struct{
	Token token.Token // = token
	Name *Identifier // hold the identifier of the binding
	Operator string // 
	Value Expression // expression that produces the value
}

func (as *AssignStatement) expressionNode() {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }

func (as *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Name.String())
	out.WriteString(as.Operator)
	if as.Value != nil {
		out.WriteString(as.Value.String())
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


// String Literal
type StringLiteral struct{
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string { return sl.Token.Literal }


// PrefixExpression
type PrefixExpression struct{
	Token token.Token // prefix token like ! -
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode(){}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	
	return out.String()
}

// Infix Expression
type InfixExpression struct{
	Token token.Token // The operator token
	Left Expression
	Operator string
	Right Expression
}
func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString( " " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean
type Boolean struct{
	Token token.Token //true or false
	Value bool
}

func (b *Boolean) expressionNode(){}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string { return b.Token.Literal }

// If Expression
type IfExpression struct{
	Token token.Token // if token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}
func (ie *IfExpression) expressionNode(){}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil{
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// Block Statement
type BlockStatement struct{
	Token token.Token // { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode(){}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements{
		out.WriteString(s.String())
	}
	return out.String()
}

// Function 
type FunctionLiteral struct {
	Token token.Token // if token
	Parameters []*Identifier
	Body *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters{
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())
	
	return out.String()
}

// FunctionCall: add(1, 2, fn(x,y){x+y;});
type CallExpression struct{
	Token token.Token // ( token
	Function Expression // Identifer of FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode(){}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments{
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	
	return out.String()
}

// List Literals
type ListLiteral struct{
	Token token.Token // [
	Elements []Expression
}

func (ll *ListLiteral) expressionNode(){}
func (ll *ListLiteral) TokenLiteral() string { return ll.Token.Literal }
func (ll *ListLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}

	for _, el := range ll.Elements{
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Index Expression - my_list[index]
type IndexExpression struct{
	Token token.Token // The [ token
	Left Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())	
	out.WriteString("])")

	return out.String()
}

