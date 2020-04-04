//go:generate goyacc -p bibtex -o bibtex.y.go bibtex.y

package bibtex

import "io"

// lexer for bibtex.
type lexer struct {
	scanner *scanner
	Errors  chan error
}

// newLexer returns a new yacc-compatible lexer.
func newLexer(r io.Reader) *lexer {
	return &lexer{scanner: newScanner(r), Errors: make(chan error, 1)}
}

// Lex is provided for yacc-compatible parser.
func (l *lexer) Lex(yylval *bibtexSymType) int {
	token, strval := l.scanner.Scan()
	yylval.strval = strval
	return int(token)
}

// Error handles error.
func (l *lexer) Error(err string) {
	l.Errors <- &ErrParse{Err: err, Pos: l.scanner.pos}
}
