package bibtex

import (
	"fmt"
	"testing"
)

// Tests basic usage of bibtex library.
func TestBasic(t *testing.T) {
	bibtex := NewBibTex()
	entry := NewBibEntry("article", "abcd1234")
	entry.AddField("title", NewBibConst("HelloWorld"))
	bibtex.AddEntry(entry)

	expected := `@article{abcd1234,
  title = {HelloWorld}
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
	bibtex := NewBibTex()
	entry := NewBibEntry("ar t icle", "ab cd 1234")
	entry.AddField("title", NewBibConst("Hello World  "))
	bibtex.AddEntry(entry)

	expected := `@article{abcd1234,
  title = {Hello World}
}
`
	if bibtex.String() != expected {
		fmt.Printf("%s\n%s\n", bibtex.String(), expected)
		t.Error("Output does not match.")
	}
}

func TestString(t *testing.T) {
	bibtex := NewBibTex()
	bibtex.AddStringVar("cat", &BibVar{Key: "cat", Value: NewBibConst("meowmeow")})
	entry := NewBibEntry("article", "abcd")
	entry.AddField("title", bibtex.GetStringVar("cat"))
	bibtex.AddEntry(entry)

	expected := `@article{abcd,
  title = {meowmeow}
}
`
	if bibtex.String() != expected {
		fmt.Printf("%s\n%s\n", bibtex.String(), expected)
		t.Error("Output does not match.")
	}
}
