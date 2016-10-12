//go:generate go tool yacc -p bibtex -o bibtex.y.go bibtex.y

package bibtex

import "io"

// Lexer for bibtex.
type Lexer struct {
	scanner *Scanner
	Errors  chan error
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{scanner: NewScanner(r), Errors: make(chan error, 1)}
}

// Lex is provided for yacc-compatible parser.
func (l *Lexer) Lex(yylval *bibtexSymType) int {
	token, strval := l.scanner.Scan()
	yylval.strval = strval
	return int(token)
}

// Error handles error.
func (l *Lexer) Error(err string) {
	l.Errors <- &ErrParse{Err: err, Pos: l.scanner.pos}
}
