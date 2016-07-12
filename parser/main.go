// Command parser reads bibtex from stdin and print out int stdout.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nickng/bibtex"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println(bibtex.Parse(r).String())
}
