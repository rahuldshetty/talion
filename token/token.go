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
	
	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION" // fn
	VAR = "VAR" // var
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"var": VAR,
}

// Check whether given alphabetic text is keyword or identifer
func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok{
		return tok
	}
	return IDENT
}