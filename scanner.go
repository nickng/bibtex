package bibtex

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Scanner represents a lexical scanner
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occrus (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.skipWhitespace()
		ch = s.read()
		if isAlphanum(ch) {
			s.unread()
			return s.scanIdent()
		}
	} else if isAlphanum(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return 0, ""
	case '@':
		return ATSIGN, string(ch)
	case '{':
		return LBRACE, string(ch)
	case ':':
		return COLON, string(ch)
	case ',':
		return COMMA, string(ch)
	case '=':
		return EQUAL, string(ch)
	case '"':
		return DQUOTE, string(ch)
	case '}':
		return RBRACE, string(ch)
	}

	fmt.Printf("Token %c unrecognised\n", ch)
	return ILLEGAL, string(ch)
}

func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == ' ' { // Continue with spaces
			_, _ = buf.WriteRune(ch)
		} else if !isAlphanum(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return IDENT, buf.String()
}

// skipWhitespace consumes the current rune and all contigious whitespace.
func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
