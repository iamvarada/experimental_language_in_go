package lexer

import "monkey_lang/token"

/* Structure for a lexer data structure */
type Lexer struct {
	input string 		// input string
	position int 		// current position being read
	readPosition int 	// position right after the current position
	ch byte 			// character at the current position
}

/* Function to initialize an instance of Lexer struct 
 * \param[in] input input string for which lexing needs to be done 
 * \return pointer to a Lexer instance 
 */
func New(input string) *Lexer {
	l := &Lexer{input : input}
	l.readChar()
	return l
}

/* function to assign token for every input character 
 * \note: Function on the Lexer structure
 * \param[in] l pointer to the instance of the Lexer instance 
 * \return token for the current character in the input string
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal : literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal : literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type  = token.LoopUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

/* function to read the current character and populate the instance of the Lexer data structure 
 * \note: Function on the Lexer structure
 * \param[in] l pointer to the instance of the Lexer instance 
 */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

/* Populate and return the token for a given character in the input string
 * \param[in] tokenType token type for the current character
 * \param[in] ch character for which the token is desired
 * \return token for the current character in the input string
 */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/* Reads an identifier/keyword 
 * \note: Gets the sub-string from the input until it encounters a non-letter character
 * \param[in] tokenType token type for the current character
 * \return sub-set string corresponding to the keyword or identifier
 */
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/* Funtion defines what is considered as a letter in this langugage and is allowed in the keywords and identifiers 
 * \note: (foo_bar, Foo_Bar, foo?bar, Foorbar! are all valid) 
 * \param[in] input character
 * \return flag indicating whether the inoput character is a letter
 */
func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'|| ch >= 'A' && ch <= 'Z' || ch == '_' || ch == '?' || ch == '!'
}

/* Skip whitespaces while lexing 
 * \param[in] l pointer to the instance of the Lexer instance 
 */
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/* Returns the number sub-string from the input 
 * \param[in] l pointer to the instance of the Lexer instance 
 * \return sub-set string corresponding to the number
 */
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/* Function determins whether the input character is a digit (b/w 0 and 9) 
 * \note: only converts posotive integers into tokens as of now (not floats/hex/oct and so on) 
 * \param[in] input character
 * \return flag indicating whether the inoput character is a digit
 */
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

/* Function to look ahead to the next character 
 * \param[in] l pointer to the instance of the Lexer instance 
 * \return chracter in the string right next to the current character
 */
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
