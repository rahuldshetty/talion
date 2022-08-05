package lexer

import "github.com/rahuldshetty/talion/token"

type Lexer struct {
	input        string
	position     int  // current position of character
	readPosition int  // current reading positon in input (after reading character)
	ch           byte // current char under examination under position
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// Fetch next Character and increment position
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII Code for NUL to handle no input or reached EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip whitespace tokens while parsing
	l.skipWhitespace()

	switch l.ch {
	case '=': 
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';': tok = newToken(token.SEMICOLON, l.ch)
	case '(': tok = newToken(token.LPAREN, l.ch)
	case ')': tok = newToken(token.RPAREN, l.ch)
	case '{': tok = newToken(token.LBRACE, l.ch)
	case '}': tok = newToken(token.RBRACE, l.ch)
	case ',': tok = newToken(token.COMMA, l.ch)
	case '+': tok = newToken(token.PLUS, l.ch)
	case '-': tok = newToken(token.MINUS, l.ch)
	case '*': tok = newToken(token.MULTIPLY, l.ch)
	case '/': tok = newToken(token.DIVIDE, l.ch)
	case '[': tok = newToken(token.LBRACKET, l.ch)
	case ']': tok = newToken(token.RBRACKET, l.ch)
	case ':': tok = newToken(token.COLON, l.ch)
	case '!': 
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.NOT, l.ch)
		}
	case '<':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: 
		// If character is letter then read the complete text identifier
		if isLetter(l.ch){
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch){
			tok = l.readDecimal()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// lookup next character but don't update position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input){
		return 0
	} else {
		return l.input[l.readPosition]
	}
}


// fetch alphabetical text from position
func (l *Lexer) readIdentifer() string {
	position := l.position

	// Identifer should start with alphabet
	if isLetter(l.ch){
		l.readChar()
	}

	// Identifier can be formed with difit or _ combination
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_'{
		l.readChar()
	}
	return l.input[position:l.position]
}

// read decimal number
func (l *Lexer) readDecimal() token.Token{
	integer := l.readNumber()

	// if come across . then try to parse next string as integer to form decimal value
	if l.ch == '.' && isDigit(l.peekChar()){
		l.readChar()
		fraction := l.readNumber()
		return token.Token{Type: token.FLOAT, Literal: integer + "." + fraction}
	}

	return token.Token{Type: token.INT, Literal: integer}
}


// fetch number text
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}

// read string
// TODO: Support for escape characters
func (l *Lexer) readString() string{
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}


// Skip whitespace Tokens
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}  
}

// String check
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch &&  ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}