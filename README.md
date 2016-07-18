# bibtex

## `nickng/bibtex` is a bibtex parser and library for Go.

[![Build Status](https://travis-ci.org/nickng/bibtex.svg?branch=master)](https://travis-ci.org/nickng/bibtex)

The bibtex format is not very well documented, this parser follows the descriptions found
[here](http://maverick.inria.fr/~Xavier.Decoret/resources/xdkbibtex/bibtex_summary.html).
Please file any issues with a minimal working example.

To get:

    go get -u github.com/nickng/bibtex

To run a test parser:

    cd $GOPATH/src/github.com/nickng/bibtex
    go run parser/main.go < example/simple.bib
