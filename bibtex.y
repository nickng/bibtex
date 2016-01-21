%{

package bibtex

import (
	"bufio"
	"os"
)

type bibTag struct {
	key string
	val string
}

var bibs BibTeX

%}

%union {
	str      string
	bibentry *BibEntry
	biblist  BibTeX
	bibtag   *bibTag
	bibtags  []*bibTag
}

%token ATSIGN LBRACE COLON EQUAL COMMA DQUOTE RBRACE
%token <str> IDENT
%type <bibentry> bibentry
%type <biblist> bibtex
%type <bibtag> tag
%type <bibtags> tags

%%

top : bibtex { bibs = $1 }
    ;

bibtex : /*empty*/       { $$ = NewBibTeX(); }
       | bibtex bibentry { $1.AddEntry($2); $$ = $1 }
       ;

bibentry : ATSIGN IDENT LBRACE IDENT COMMA tags RBRACE { $$ = NewBibEntry($2, $4); for _, t := range $6 { $$.AddField(t.key, t.val) } }
         ;

tag : IDENT EQUAL        IDENT        { $$ = &bibTag{key: $1, val: $3}; }
    | IDENT EQUAL LBRACE IDENT RBRACE { $$ = &bibTag{key: $1, val: $4}; }
    | IDENT EQUAL DQUOTE IDENT DQUOTE { $$ = &bibTag{key: $1, val: $4}; }
    ;

tags : tag            { $$ = []*bibTag{$1} }
     | tags COMMA tag { $$ = append($1, $3) }
     ;

%%

func Parse(in *os.File) *BibTeX {
    bibtexParse(NewLexer(bufio.NewReader(in)))
    return &bibs
}
