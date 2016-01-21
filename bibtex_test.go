package bibtex

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	bibtex := NewBibTeX()
	entry := NewBibEntry("article", "abcd1234")
	entry.AddField("title", "HelloWorld")
	bibtex.AddEntry(entry)

	expected := `@article{abcd1234,
  title = {{HelloWorld}}
}`
	if bibtex.String() != expected {
		fmt.Printf("%s\n", bibtex.String())
		t.Error("Output does not match.")
	}
}
