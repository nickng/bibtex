%{
package bibtex

import (
	"io"
	"fmt"
)

type bibTag struct {
	key string
	val BibString
}

var bib *BibTex // Only for holding current bib
%}

%union {
	bibtex   *BibTex
	strval   string
	bibentry *BibEntry
	bibtag   *bibTag
	bibtags  []*bibTag
	strings  BibString
}

%token tCOMMENT tSTRING tPREAMBLE
%token tATSIGN tCOLON tEQUAL tCOMMA tPOUND tLBRACE tRBRACE tDQUOTE tLPAREN tRPAREN
%token <strval> tBAREIDENT tIDENT
%type <bibtex> bibtex
%type <bibentry> bibentry
%type <bibtag> tag stringentry
%type <bibtags> tags
%type <strings> longstring preambleentry

%%

top : bibtex { }
    ;

bibtex : /* empty */          { $$ = NewBibTex(); bib = $$ }
       | bibtex bibentry      { $$ = $1; $$.AddEntry($2) }
       | bibtex commententry  { $$ = $1 }
       | bibtex stringentry   { $$ = $1; $$.AddStringVar($2.key, $2.val) }
       | bibtex preambleentry { $$ = $1; $$.AddPreamble($2) }
       ;

bibentry : tATSIGN tBAREIDENT tLBRACE tBAREIDENT tCOMMA tags tRBRACE { $$ = NewBibEntry($2, $4); for _, t := range $6 { $$.AddField(t.key, t.val) } }
         | tATSIGN tBAREIDENT tLPAREN tBAREIDENT tCOMMA tags tRPAREN { $$ = NewBibEntry($2, $4); for _, t := range $6 { $$.AddField(t.key, t.val) } }
         ;

commententry : tATSIGN tCOMMENT tLBRACE longstring tRBRACE {}
             | tATSIGN tCOMMENT tLPAREN longstring tRBRACE {}
             ;

stringentry : tATSIGN tSTRING tLBRACE tBAREIDENT tEQUAL longstring tRBRACE { $$ = &bibTag{key: $4, val: $6 } }
            | tATSIGN tSTRING tLPAREN tBAREIDENT tEQUAL longstring tRBRACE { $$ = &bibTag{key: $4, val: $6 } }
            ;

preambleentry : tATSIGN tPREAMBLE tLBRACE longstring tRBRACE { $$ = $4 }
              | tATSIGN tPREAMBLE tLPAREN longstring tRPAREN { $$ = $4 }
              ;

longstring :                  tIDENT     { $$ = NewBibConst($1) }
           |                  tBAREIDENT { v,err := bib.GetStringVar($1); if err != nil { bibtexlex.Error(fmt.Sprintf("%v",err)) } else {$$ = v} }
           | longstring tPOUND tIDENT     { $$ = NewBibComposite($1); $$.(*BibComposite).Append(NewBibConst($3))}
           | longstring tPOUND tBAREIDENT { $$ = NewBibComposite($1); v,err := bib.GetStringVar($3); if err != nil {bibtexlex.Error(fmt.Sprintf("%v",err))}  else {$$.(*BibComposite).Append(v)} }
           ;

tag : /* empty */                { }
    | tBAREIDENT tEQUAL longstring { $$ = &bibTag{key: $1, val: $3} }
    ;

tags : tag            { if $1 != nil { $$ = []*bibTag{$1}; } }
     | tags tCOMMA tag { if $3 == nil { $$ = $1 } else { $$ = append($1, $3) } }
     ;

%%

// Parse is the entry point to the bibtex parser.
func Parse(r io.Reader) (*BibTex, error) {
	l := newLexer(r)
	bibtexParse(l)
	switch {
	case len(l.Errors) > 0: // Non-yacc errors
		return nil, l.Errors[0]
	case len(l.ParseErrors) > 0:
		return nil, l.ParseErrors[0]
	default:
		return bib, nil
	}
}
