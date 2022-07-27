# Tokens available for Tokenizer/Lexing

ILLEGAL = "ILLEGAL"
EOF = "EOF"

// Identifiers + literals
IDENT = "IDENT" // foo, bar, add, x, y
INT = "INT" // 3, 5151, 45

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
GT = ">"


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
TRUE = "TRUE"
FALSE = "FALSE"
IF = "IF"
ELSE = "ELSE"
RETURN = "RETURN"
