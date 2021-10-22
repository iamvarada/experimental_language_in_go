package token

type TokenType string

/* Data structure defining a token in the Monkey language */
type Token struct {
	Type    TokenType
	Literal string
}

/* Map of keyword's literal to its token */
var keywords = map[string]TokenType {
	"fn" : FUNCTION,
	"let" : LET,
	"true" : TRUE,
	"false" : FALSE,
	"if" : IF,
	"else" : ELSE,
	"return" : RETURN,
}

/* Compares the input string to see whether it is a keyword or a variable 
 * \note: set of keywords are defined in the "keywords" map variable
 * \param[in] ident_literal inpout string literal
 * \return token type for the input literal
 */
func LoopUpIdent(ident_literal string) TokenType {
	if tok, key_found := keywords[ident_literal]; key_found {
		return tok
	}
	return IDENT
}

/* List of tokens */
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	/* Identifiers */
	IDENT = "IDENT"
	INT   = "INT"

	/* Operators */
	ASSIGN = "="
	PLUS   = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	/* Delimiters */
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	/* Keywords */
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)
