package bibtex

import (
	"errors"
	"fmt"
)

var (
	// ErrUnexpectedAtsign is an error for unexpected @ in {}.
	ErrUnexpectedAtsign = errors.New("Unexpected @ sign")
	// ErrUnknownStringVar is an error for looking up undefined string var.
	ErrUnknownStringVar = errors.New("Unknown string variable")
)

// SyntaxError is an error at the scanner level.
type SyntaxError struct {
	What string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("Syntax error: %s\n", e.What)
}
