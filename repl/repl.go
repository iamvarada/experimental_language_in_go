package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey_lang/lexer"
	"monkey_lang/token"
)

const PROMPT = ">>"

/* Reads the input and prints the tokens using the Lexer
 * \param[in] in console input from the user
 * \param[out] out output written opn the console 
 */
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
