package bibtex

import "fmt"

// SyntaxError is an error at the scanner level.
type SyntaxError struct {
	What string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("Syntax error: %s\n", e.What)
}
