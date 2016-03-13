package bibtex

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

// Scanner is a lexical scanner
type Scanner struct {
	r *bufio.Reader
}

const (
	bibEntryMode = iota
	bibTagMode
)

var (
	mode = bibEntryMode
)

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.eof is returned).
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
		s.ignoreWhitespace()
		ch = s.read()
	}

	if isAlphanum(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return 0, ""
	case '@':
		mode = bibEntryMode
		return ATSIGN, string(ch)
	case ':':
		return COLON, string(ch)
	case ',':
		return COMMA, string(ch)
	case '=':
		return EQUAL, string(ch)
	case '"':
		s.unread()
		return s.scanIdent()
	case '{':
		if mode == bibEntryMode { // The first { in toplevel (bibtexMode)
			mode = bibTagMode
			return LBRACE, string(ch)
		}
		s.unread()
		return s.scanIdent()
	case '}':
		return RBRACE, string(ch)
	}

	log.Fatal(SyntaxError{What: fmt.Sprintf("Token %c unrecognised\n", ch)})
	return ILLEGAL, string(ch)
}

func (s *Scanner) scanQuoted() (tok Token, lit string) {
	var buf bytes.Buffer
	var endCh rune

	if ch := s.read(); ch == '{' {
		endCh = '}'
	} else if ch == '"' {
		endCh = '"'
	}

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isAlphanum(ch) || isWhitespace(ch) || isSymbol(ch) {
			_, _ = buf.WriteRune(ch)
		} else if ch == '\\' {
			s.unread()
			if escTok, escLit := s.scanEscape(); escTok != ILLEGAL {
				buf.WriteString(escLit)
			} else {
				break
			}
		} else if ch == '{' {
			s.unread()
			if qTok, qLit := s.scanQuoted(); qTok != ILLEGAL {
				buf.WriteString(qLit)
			} else {
				break
			}
		} else if ch == endCh {
			return IDENT, buf.String()
		}
	}
	return ILLEGAL, buf.String() // Unterminated quote or illegal characters
}

func (s *Scanner) scanComment() (tok Token, lit string) {
	var buf bytes.Buffer

	if ch := s.read(); ch != '%' {
		return ILLEGAL, string(ch)
	}

	for {
		if ch := s.read(); ch == eof || ch == '\n' {
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	return IDENT, buf.String()
}

// scanIdent parses a string, could be quoted
func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer

	switch ch := s.read(); ch {
	case '"', '{':
		s.unread()
		return s.scanQuoted()
	}
	s.unread()

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isAlphanum(ch) && !isBareSymbol(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	return IDENT, buf.String()
}

func (s *Scanner) scanEscape() (tok Token, lit string) {
	var buf bytes.Buffer

	buf.WriteRune(s.read()) // escape character '\'

	switch ch := s.read(); ch {
	case '`', '\'', '^', '"', '~', '=', '.', 'u', 'v', 'H', 't', 'c', 'd', 'b', 'l', 'L':
		buf.WriteRune(ch)
		s := buf.String()
		return IDENT, s
	case 'a': // aa, ae, oe
		switch ch2 := s.read(); ch2 {
		case 'a', 'e', 'o':
			buf.WriteRune(ch)
			buf.WriteRune(ch2)
			return IDENT, buf.String()
		}
	case 'o':
		switch ch2 := s.read(); ch2 {
		case 'e':
			buf.WriteRune(ch)
			buf.WriteRune(ch2)
			return IDENT, buf.String()
		}
	case 's': // ss
		switch ch2 := s.read(); ch2 {
		case 's':
			buf.WriteRune(ch)
			buf.WriteRune(ch2)
			return IDENT, buf.String()
		}
	case 'A': // AA, AE
		switch ch2 := s.read(); ch2 {
		case 'A', 'E':
			buf.WriteRune(ch)
			buf.WriteRune(ch2)
			return IDENT, buf.String()
		}
	case '{', '}', '\\': // Illegal characters
		buf.WriteRune(ch)
		return IDENT, buf.String()
	default:
		buf.WriteRune(ch)
	}

	escapeStr := buf.String()
	log.Println(SyntaxError{What: fmt.Sprintf("Unknown escape string %s\n", escapeStr)})
	return ILLEGAL, escapeStr
}

// ignoreWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) ignoreWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
