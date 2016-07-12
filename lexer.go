//go:generate go tool yacc -p bibtex -o bibtex.y.go bibtex.y

package bibtex

import (
	"io"
	"log"
)

// Lexer for bibtex.
type Lexer struct {
	scanner *Scanner
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{scanner: NewScanner(r)}
}

// Lex is provided for yacc-compatible parser.
func (l *Lexer) Lex(yylval *bibtexSymType) int {
	token, strval := l.scanner.Scan()
	yylval.strval = strval
	return int(token)
}

// Error handles error.
func (l *Lexer) Error(err string) {
	log.Fatalf("parse error: %s", err)
}
