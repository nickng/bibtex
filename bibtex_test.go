package bibtex

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

// Test that the parser accepts all valid bibtex files in the example/ dir.
func TestParser(t *testing.T) {
	examples := []string{
		"example/biblatex-examples.bib",
		"example/embeddedtex.bib",
		"example/field-error.bib",
		"example/quoted.bib",
		"example/simple.bib",
		"example/space.bib",
		"example/symbols.bib",
		"example/var.bib",
	}

	for _, ex := range examples {
		t.Logf("Parsing example: %s", ex)
		b, err := ioutil.ReadFile(ex)
		if err != nil {
			t.Errorf("Cannot read %s: %v", ex, err)
		}
		_, err = Parse(bytes.NewReader(b))
		if err != nil {
			t.Errorf("Cannot parse valid bibtex file %s: %v", ex, err)
		}
	}
}

// Tests that multiple parse returns different instances of the parsed BibTex.
// Otherwise the number of entries will pile up. (Issue #4)
func TestMultiParse(t *testing.T) {
	examples := []string{
		"example/simple.bib",
		"example/simple.bib",
		"example/simple.bib",
	}

	var bibs []*BibTex
	for _, ex := range examples {
		t.Logf("Parsing example: %s", ex)
		b, err := ioutil.ReadFile(ex)
		if err != nil {
			t.Errorf("Cannot read %s: %v", ex, err)
		}
		s, err := Parse(bytes.NewReader(b))
		if err != nil {
			t.Errorf("Cannot parse valid bibtex file %s: %v", ex, err)
		}
		if want, got := 2, len(s.Entries); want != got {
			t.Errorf("Expecting %d entries but got %d", want, got)
		}
		bibs = append(bibs, s)
	}
	for _, bib := range bibs {
		if want, got := 2, len(bib.Entries); want != got {
			t.Errorf("Expecting %d entries but got %d", want, got)
		}
	}
}
