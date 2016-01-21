// Package bibtex provides a simple bibtex parser and data structure to
// represent bibtex records.
package bibtex // go get github.com/nickng/bibtex

import (
	"bytes"
	"fmt"
)

// BibEntry is a record of BibTeX record.
type BibEntry struct {
	Type     string
	CiteName string
	Fields   map[string]string
}

// BibTeX is a list of BibTeX entries.
type BibTeX []*BibEntry

// NewBibTeX creates a new BibTeX data structure.
func NewBibTeX() BibTeX {
	return []*BibEntry{}
}

// AddEntry adds an entry to the BibTeX data structure.
func (bib *BibTeX) AddEntry(entry *BibEntry) {
	*bib = append(*bib, entry)
}

// NewBibEntry creates a new BibTeX entry.
func NewBibEntry(entryType string, citeName string) *BibEntry {
	return &BibEntry{
		Type:     entryType,
		CiteName: citeName,
		Fields:   map[string]string{},
	}
}

// AddField adds a field (key-value) to a BibTeX entry.
func (entry *BibEntry) AddField(name string, value string) {
	entry.Fields[name] = value
}

// String prints a BibTeX data structure as a BibTeX string.
func (bib BibTeX) String() string {
	var bibtex bytes.Buffer

	for _, entry := range bib {
		bibtex.WriteString(fmt.Sprintf("@%s{%s,\n", entry.Type, entry.CiteName))
		for key, val := range entry.Fields {
			bibtex.WriteString(fmt.Sprintf("  %s = {{%s}},\n", key, val))
		}
		bibtex.Truncate(bibtex.Len() - 2)
		bibtex.WriteString(fmt.Sprintf("\n}\n"))
	}
	bibtex.Truncate(bibtex.Len() - 1)
	return bibtex.String()
}
