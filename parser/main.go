// Command parser reads bibtex from stdin and print out int stdout.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nickng/bibtex"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	parsed, err := bibtex.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(parsed.PrettyString())
}
