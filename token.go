package bibtex

// Lexer token.
type Token int

const (
	// ILLEGAL stands for an invalid token.
	ILLEGAL Token = iota
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlpha(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}

func isAlphanum(ch rune) bool {
	return isAlpha(ch) || isDigit(ch)
}
