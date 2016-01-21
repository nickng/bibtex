//go:generate go tool yacc -p bibtex bibtex.y

package bibtex

import (
	"io"
	"log"
)

type bibtexLex struct {
	scanner *Scanner
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *bibtexLex {
	return &bibtexLex{scanner: NewScanner(r)}
}

func (l *bibtexLex) Lex(yylval *bibtexSymType) int {
	tok, lit := l.scanner.Scan()
	yylval.str = lit

	return int(tok)
}

func (l *bibtexLex) Error(err string) {
	log.Printf("parse error %s", err)
}
