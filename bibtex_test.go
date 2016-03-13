package bibtex

import (
	"fmt"
	"testing"
)

// Tests basic usage of bibtex library.
func TestBasic(t *testing.T) {
	bibtex := NewBibTeX()
	entry := NewBibEntry("article", "abcd1234")
	entry.AddField("title", "HelloWorld")
	bibtex.AddEntry(entry)

	expected := `@article{abcd1234,
  title = {{HelloWorld}}
}
`
	if bibtex.String() != expected {
		fmt.Printf("%s\n%s\n", bibtex.String(), expected)
		t.Error("Output does not match.")
	}
}

// Tests usage of bibtex library of strings with spaces.
// The tag content should be unchanged, other strings will have spaces removed.
func TestSpaces(t *testing.T) {
	bibtex := NewBibTeX()
	entry := NewBibEntry("ar t icle", "ab cd 1234")
	entry.AddField("title", "Hello World  ")
	bibtex.AddEntry(entry)

	expected := `@article{abcd1234,
  title = {{Hello World}}
}
`
	if bibtex.String() != expected {
		fmt.Printf("%s\n%s\n", bibtex.String(), expected)
		t.Error("Output does not match.")
	}
}
