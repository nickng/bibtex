%{
package bibtex

import (
	"io"
)

type bibTag struct {
	key string
	val BibString
}

var bibs = NewBibTex()
%}

%union {
	strval   string
	bibentry *BibEntry
	bibtag   *bibTag
	bibtags  []*bibTag
	strings  BibString
}

%token COMMENT STRING PREAMBLE
%token ATSIGN COLON EQUAL COMMA POUND LBRACE RBRACE DQUOTE LPAREN RPAREN
%token <strval> BAREIDENT IDENT
%type <bibentry> bibentry
%type <bibtag> tag
%type <bibtags> tags
%type <strings> longstring

%%

top : bibtex { }
    ;

bibtex : /* empty */          { }
       | bibtex bibentry      { bibs.AddEntry($2) }
       | bibtex commententry  { }
       | bibtex stringentry   { }
       | bibtex preambleentry { }
       ;

bibentry : ATSIGN BAREIDENT LBRACE BAREIDENT COMMA tags RBRACE { $$ = NewBibEntry($2, $4); for _, t := range $6 { $$.AddField(t.key, t.val) } }
         | ATSIGN BAREIDENT LPAREN BAREIDENT COMMA tags RPAREN { $$ = NewBibEntry($2, $4); for _, t := range $6 { $$.AddField(t.key, t.val) } }
         ;

commententry : ATSIGN COMMENT LBRACE longstring RBRACE {}
             | ATSIGN COMMENT LPAREN longstring RBRACE {}
             ;

stringentry : ATSIGN STRING LBRACE BAREIDENT EQUAL longstring RBRACE { bibs.AddStringVar($4, $6) }
            | ATSIGN STRING LPAREN BAREIDENT EQUAL longstring RBRACE { bibs.AddStringVar($4, $6) }
            ;

preambleentry : ATSIGN PREAMBLE LBRACE longstring RBRACE { bibs.AddPreamble($4) }
              | ATSIGN PREAMBLE LPAREN longstring RPAREN { bibs.AddPreamble($4) }
              ;

longstring :                  IDENT     { $$ = NewBibConst($1) }
           |                  BAREIDENT { $$ = bibs.GetStringVar($1) }
           | longstring POUND IDENT     { $$ = NewBibComposite($1); $$.(*BibComposite).Append(NewBibConst($3))}
           | longstring POUND BAREIDENT { $$ = NewBibComposite($1); $$.(*BibComposite).Append(bibs.GetStringVar($3)) }
           ;

tag : /* empty */                { }
    | BAREIDENT EQUAL longstring { $$ = &bibTag{key: $1, val: $3} }
    ;

tags : tag            { $$ = []*bibTag{$1} }
     | tags COMMA tag { if $3 == nil { $$ = $1 } else { $$ = append($1, $3) } }
     ;

%%

// Parse is the entry point to the bibtex parser.
func Parse(r io.Reader) (*BibTex, error) {
	l := NewLexer(r)
	bibtexParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return bibs, nil
	}
}
