package token

//  different types of token
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // foo, bar, add, x, y
	INT = "INT" // 3, 5151, 45
	STRING = "STRING" // "hello world"
	
	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	MULTIPLY = "*"
	DIVIDE = "/"
	
	NOT = "!"
	EQ = "=="
	NOT_EQ = "!="
	
	LT = "<"
	LTE = "<="
	GT = ">"
	GTE = ">="
	

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// List
	LBRACKET = "["
	RBRACKET = "]"

	COLON = ":"

	// Keywords
	FUNCTION = "FUNCTION" // fn
	VAR = "VAR" // var
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"var": VAR,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

// Check whether given alphabetic text is keyword or identifer
func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok{
		return tok
	}
	return IDENT
}