// Command prettybib pretty prints a bib entry from command line.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/nickng/bibtex"
)

var (
	infile  = flag.String("in", "", "Input file (default: stdin)")
	outfile = flag.String("out", "", "Output file (default: stdout)")
	config  = flag.String("conf", "", "Filter config to use")

	reader = os.Stdin
	writer = os.Stdout
)

// Config holds a configuration for filtering rules.
type Config struct {
	BibType map[string]struct {
		Required []string
		Remove   []string
	}
}

func main() {
	flag.Parse()
	if *infile != "" || len(flag.Args()) > 0 {
		if len(flag.Args()) > 0 {
			*infile = flag.Arg(0)
		}
		rdFile, err := os.Open(*infile)
		if err != nil {
			log.Fatal(err)
		}
		defer rdFile.Close()
		reader = rdFile
	}

	if *outfile != "" {
		wrFile, err := os.OpenFile(*outfile, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer wrFile.Close()
		writer = wrFile
	}

	parsed, err := bibtex.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	if *config != "" {
		var conf Config
		if _, err := toml.DecodeFile(*config, &conf); err != nil {
			log.Fatalf("Cannot read config: %s", err)
		}
		filter(parsed, &conf)
	}
	fmt.Fprintf(writer, parsed.PrettyString())
}

func filter(bib *bibtex.BibTex, conf *Config) {
	for _, entry := range bib.Entries {
		if rule, ok := conf.BibType[entry.Type]; ok {
			for _, required := range rule.Required {
				if _, found := entry.Fields[required]; !found {
					entry.Fields[required] = bibtex.NewBibConst("")
				}
			}
			for _, remove := range rule.Remove {
				if _, found := entry.Fields[remove]; found {
					delete(entry.Fields, remove)
				}
			}
		}
	}
}
