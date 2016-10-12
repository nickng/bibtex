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

// ErrParse is a parse error.
type ErrParse struct {
	Pos TokenPos
	Err string // Error string returned from parser.
}

func (e *ErrParse) Error() string {
	return fmt.Sprintf("Parse failed at %s: %s", e.Pos, e.Err)
}
