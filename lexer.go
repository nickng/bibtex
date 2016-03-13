//go:generate go tool yacc -p bibtex bibtex.y

package bibtex

import (
	"io"
	"log"
)

// Lexer for bibtex
type Lexer struct {
	scanner *Scanner
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{scanner: NewScanner(r)}
}

// Lex is provided for yacc-compatible parser.
func (l *Lexer) Lex(yylval *bibtexSymType) int {
	tok, lit := l.scanner.Scan()
	yylval.str = lit

	return int(tok)
}

func (l *Lexer) Error(err string) {
	log.Printf("parse error: %s", err)
}
