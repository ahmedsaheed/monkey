package lexer

import "monkey/lang-monkey/token"

type Lexer struct {
	input        string 
	position     int    
	readPosition int    
	char         byte  
}

func New(input string) *Lexer {
	L := &Lexer{input: input}
	L.readChar()
	return L
}

func (L *Lexer) readChar() {
	if L.readPosition >= len(L.input) {
		L.char = 0 
	} else {
		L.char = L.input[L.readPosition]
	}
	L.position = L.readPosition
	L.readPosition += 1
}

func (L *Lexer) peekChar() byte {
	if L.readPosition >= len(L.input) {
		return 0
	} else {
		return L.input[L.readPosition]
	}
}

func (L *Lexer) NextToken() token.Token {
	var tok token.Token
	L.skipWhiteSpaces()
	switch L.char {
	default:
		if isLetter(L.char) {
			tok.Literal = L.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(L.char) {
			tok.Type = token.INT
			tok.Literal = L.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, L.char)
		}
	case '=':
		if L.peekChar() == '=' {
			char := L.char
			L.readChar()
			literal := string(char) + string(L.char)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, L.char)
		}
	case ';':
		tok = newToken(token.SEMICOLON, L.char)
	case '(':
		tok = newToken(token.LPAREN, L.char)
	case ')':
		tok = newToken(token.RPAREN, L.char)
	case ',':
		tok = newToken(token.COMMA, L.char)
	case '+':
		tok = newToken(token.PLUS, L.char)
	case '{':
		tok = newToken(token.LBRACE, L.char)
	case '}':
		tok = newToken(token.RBRACE, L.char)
	case '!':
		if L.peekChar() == '=' {
			char := L.char
			L.readChar()
			literal := string(char) + string(L.char)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, L.char)
		}
	case '/':
		tok = newToken(token.SLASH, L.char)
	case '*':
		tok = newToken(token.ASTERISK, L.char)
	case '-':
		tok = newToken(token.MINUS, L.char)
	case '<':
		tok = newToken(token.LT, L.char)
	case '>':
		tok = newToken(token.GT, L.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	L.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (L *Lexer) readIdentifier() string {
	position := L.position
	for isLetter(L.char) {
		L.readChar()
	}
	return L.input[position:L.position]

}

func (L *Lexer) readNumber() string {
	position := L.position
	for isDigit(L.char) {
		L.readChar()
	}
	return L.input[position:L.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9' 
}

func (L *Lexer) skipWhiteSpaces() {
	for L.char == ' ' || L.char == '\t' || L.char == '\n' || L.char == '\r' {
		L.readChar()
	}
}
