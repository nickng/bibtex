package bibtex

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
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
	examples, err := filepath.Glob("example/*.bib")
	if err != nil {
		t.Fatal(err)
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

func TestPrettyStringRoundTrip(t *testing.T) {
	examples, err := filepath.Glob("example/*.bib")
	if err != nil {
		t.Fatal(err)
	}

	for _, ex := range examples {
		// Read input.
		b, err := ioutil.ReadFile(ex)
		if err != nil {
			t.Fatal(err)
		}

		// Parse into BibTeX.
		bib, err = Parse(bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}

		// Pretty print it and parse it again.
		s := bib.PrettyString()

		bib2, err := Parse(strings.NewReader(s))
		if err != nil {
			t.Fatal(err)
		}

		// Check equality.
		AssertEntryListsEqual(t, bib.Entries, bib2.Entries)
	}
}

func AssertEntryListsEqual(t *testing.T, a, b []*BibEntry) {
	t.Helper()

	if len(a) != len(b) {
		t.Fatalf("length mismatch")
	}

	for i := range a {
		AssertEntriesEqual(t, a[i], b[i])
	}
}

func AssertEntriesEqual(t *testing.T, a, b *BibEntry) {
	if a.Type != b.Type {
		t.Error("type mismatch")
	}
	if a.CiteName != b.CiteName {
		t.Error("cite name mismatch")
	}
	if len(a.Fields) != len(b.Fields) {
		t.Fatal("different number of fields")
	}
	for key := range a.Fields {
		if a.Fields[key].String() != b.Fields[key].String() {
			t.Fatalf("mismatch on field %q", key)
		}
	}
}
