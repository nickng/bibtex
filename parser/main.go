// Command parser reads bibtex from stdin and print out int stdout.
package main

import (
	"fmt"
	"os"

	"github.com/nickng/bibtex"
)

func main() {
	fmt.Println(bibtex.Parse(os.Stdin).String())
}
