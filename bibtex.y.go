//line bibtex.y:2
package bibtex

import __yyfmt__ "fmt"

//line bibtex.y:2
import (
	"io"
)

type bibTag struct {
	key string
	val string
}

var bibs BibTeX

//line bibtex.y:16
type bibtexSymType struct {
	yys      int
	strval   string
	bibentry *BibEntry
	biblist  BibTeX
	bibtag   *bibTag
	bibtags  []*bibTag
}

const ATSIGN = 57346
const LBRACE = 57347
const COLON = 57348
const EQUAL = 57349
const COMMA = 57350
const RBRACE = 57351
const IDENT = 57352

var bibtexToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ATSIGN",
	"LBRACE",
	"COLON",
	"EQUAL",
	"COMMA",
	"RBRACE",
	"IDENT",
}
var bibtexStatenames = [...]string{}

const bibtexEofCode = 1
const bibtexErrCode = 2
const bibtexInitialStackSize = 16

//line bibtex.y:51

// Parse is the entry point to the bibtex parser.
func Parse(r io.Reader) *BibTeX {
	bibtexParse(NewLexer(r))
	return &bibs
}

//line yacctab:1
var bibtexExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const bibtexNprod = 9
const bibtexPrivate = 57344

var bibtexTokenNames []string
var bibtexStates []string

const bibtexLast = 16

var bibtexAct = [...]int{

	10, 16, 11, 13, 12, 14, 7, 5, 8, 6,
	4, 1, 9, 2, 15, 3,
}
var bibtexPact = [...]int{

	-1000, -1000, 6, -1000, -3, 4, -4, 0, -8, -5,
	-1000, -2, -1000, -8, -9, -1000, -1000,
}
var bibtexPgo = [...]int{

	0, 15, 13, 0, 12, 11,
}
var bibtexR1 = [...]int{

	0, 5, 2, 2, 1, 3, 3, 4, 4,
}
var bibtexR2 = [...]int{

	0, 1, 0, 2, 7, 0, 3, 1, 3,
}
var bibtexChk = [...]int{

	-1000, -5, -2, -1, 4, 10, 5, 10, 8, -4,
	-3, 10, 9, 8, 7, -3, 10,
}
var bibtexDef = [...]int{

	2, -2, 1, 3, 0, 0, 0, 0, 5, 0,
	7, 0, 4, 5, 0, 8, 6,
}
var bibtexTok1 = [...]int{

	1,
}
var bibtexTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var bibtexTok3 = [...]int{
	0,
}

var bibtexErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	bibtexDebug        = 0
	bibtexErrorVerbose = false
)

type bibtexLexer interface {
	Lex(lval *bibtexSymType) int
	Error(s string)
}

type bibtexParser interface {
	Parse(bibtexLexer) int
	Lookahead() int
}

type bibtexParserImpl struct {
	lval  bibtexSymType
	stack [bibtexInitialStackSize]bibtexSymType
	char  int
}

func (p *bibtexParserImpl) Lookahead() int {
	return p.char
}

func bibtexNewParser() bibtexParser {
	return &bibtexParserImpl{}
}

const bibtexFlag = -1000

func bibtexTokname(c int) string {
	if c >= 1 && c-1 < len(bibtexToknames) {
		if bibtexToknames[c-1] != "" {
			return bibtexToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func bibtexStatname(s int) string {
	if s >= 0 && s < len(bibtexStatenames) {
		if bibtexStatenames[s] != "" {
			return bibtexStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func bibtexErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !bibtexErrorVerbose {
		return "syntax error"
	}

	for _, e := range bibtexErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + bibtexTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := bibtexPact[state]
	for tok := TOKSTART; tok-1 < len(bibtexToknames); tok++ {
		if n := base + tok; n >= 0 && n < bibtexLast && bibtexChk[bibtexAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if bibtexDef[state] == -2 {
		i := 0
		for bibtexExca[i] != -1 || bibtexExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; bibtexExca[i] >= 0; i += 2 {
			tok := bibtexExca[i]
			if tok < TOKSTART || bibtexExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if bibtexExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += bibtexTokname(tok)
	}
	return res
}

func bibtexlex1(lex bibtexLexer, lval *bibtexSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = bibtexTok1[0]
		goto out
	}
	if char < len(bibtexTok1) {
		token = bibtexTok1[char]
		goto out
	}
	if char >= bibtexPrivate {
		if char < bibtexPrivate+len(bibtexTok2) {
			token = bibtexTok2[char-bibtexPrivate]
			goto out
		}
	}
	for i := 0; i < len(bibtexTok3); i += 2 {
		token = bibtexTok3[i+0]
		if token == char {
			token = bibtexTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = bibtexTok2[1] /* unknown char */
	}
	if bibtexDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", bibtexTokname(token), uint(char))
	}
	return char, token
}

func bibtexParse(bibtexlex bibtexLexer) int {
	return bibtexNewParser().Parse(bibtexlex)
}

func (bibtexrcvr *bibtexParserImpl) Parse(bibtexlex bibtexLexer) int {
	var bibtexn int
	var bibtexVAL bibtexSymType
	var bibtexDollar []bibtexSymType
	_ = bibtexDollar // silence set and not used
	bibtexS := bibtexrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	bibtexstate := 0
	bibtexrcvr.char = -1
	bibtextoken := -1 // bibtexrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		bibtexstate = -1
		bibtexrcvr.char = -1
		bibtextoken = -1
	}()
	bibtexp := -1
	goto bibtexstack

ret0:
	return 0

ret1:
	return 1

bibtexstack:
	/* put a state and value onto the stack */
	if bibtexDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", bibtexTokname(bibtextoken), bibtexStatname(bibtexstate))
	}

	bibtexp++
	if bibtexp >= len(bibtexS) {
		nyys := make([]bibtexSymType, len(bibtexS)*2)
		copy(nyys, bibtexS)
		bibtexS = nyys
	}
	bibtexS[bibtexp] = bibtexVAL
	bibtexS[bibtexp].yys = bibtexstate

bibtexnewstate:
	bibtexn = bibtexPact[bibtexstate]
	if bibtexn <= bibtexFlag {
		goto bibtexdefault /* simple state */
	}
	if bibtexrcvr.char < 0 {
		bibtexrcvr.char, bibtextoken = bibtexlex1(bibtexlex, &bibtexrcvr.lval)
	}
	bibtexn += bibtextoken
	if bibtexn < 0 || bibtexn >= bibtexLast {
		goto bibtexdefault
	}
	bibtexn = bibtexAct[bibtexn]
	if bibtexChk[bibtexn] == bibtextoken { /* valid shift */
		bibtexrcvr.char = -1
		bibtextoken = -1
		bibtexVAL = bibtexrcvr.lval
		bibtexstate = bibtexn
		if Errflag > 0 {
			Errflag--
		}
		goto bibtexstack
	}

bibtexdefault:
	/* default state action */
	bibtexn = bibtexDef[bibtexstate]
	if bibtexn == -2 {
		if bibtexrcvr.char < 0 {
			bibtexrcvr.char, bibtextoken = bibtexlex1(bibtexlex, &bibtexrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if bibtexExca[xi+0] == -1 && bibtexExca[xi+1] == bibtexstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			bibtexn = bibtexExca[xi+0]
			if bibtexn < 0 || bibtexn == bibtextoken {
				break
			}
		}
		bibtexn = bibtexExca[xi+1]
		if bibtexn < 0 {
			goto ret0
		}
	}
	if bibtexn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			bibtexlex.Error(bibtexErrorMessage(bibtexstate, bibtextoken))
			Nerrs++
			if bibtexDebug >= 1 {
				__yyfmt__.Printf("%s", bibtexStatname(bibtexstate))
				__yyfmt__.Printf(" saw %s\n", bibtexTokname(bibtextoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for bibtexp >= 0 {
				bibtexn = bibtexPact[bibtexS[bibtexp].yys] + bibtexErrCode
				if bibtexn >= 0 && bibtexn < bibtexLast {
					bibtexstate = bibtexAct[bibtexn] /* simulate a shift of "error" */
					if bibtexChk[bibtexstate] == bibtexErrCode {
						goto bibtexstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if bibtexDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", bibtexS[bibtexp].yys)
				}
				bibtexp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if bibtexDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", bibtexTokname(bibtextoken))
			}
			if bibtextoken == bibtexEofCode {
				goto ret1
			}
			bibtexrcvr.char = -1
			bibtextoken = -1
			goto bibtexnewstate /* try again in the same state */
		}
	}

	/* reduction by production bibtexn */
	if bibtexDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", bibtexn, bibtexStatname(bibtexstate))
	}

	bibtexnt := bibtexn
	bibtexpt := bibtexp
	_ = bibtexpt // guard against "declared and not used"

	bibtexp -= bibtexR2[bibtexn]
	// bibtexp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if bibtexp+1 >= len(bibtexS) {
		nyys := make([]bibtexSymType, len(bibtexS)*2)
		copy(nyys, bibtexS)
		bibtexS = nyys
	}
	bibtexVAL = bibtexS[bibtexp+1]

	/* consult goto table to find next state */
	bibtexn = bibtexR1[bibtexn]
	bibtexg := bibtexPgo[bibtexn]
	bibtexj := bibtexg + bibtexS[bibtexp].yys + 1

	if bibtexj >= bibtexLast {
		bibtexstate = bibtexAct[bibtexg]
	} else {
		bibtexstate = bibtexAct[bibtexj]
		if bibtexChk[bibtexstate] != -bibtexn {
			bibtexstate = bibtexAct[bibtexg]
		}
	}
	// dummy call; replaced with literal code
	switch bibtexnt {

	case 1:
		bibtexDollar = bibtexS[bibtexpt-1 : bibtexpt+1]
		//line bibtex.y:33
		{
			bibs = bibtexDollar[1].biblist
		}
	case 2:
		bibtexDollar = bibtexS[bibtexpt-0 : bibtexpt+1]
		//line bibtex.y:36
		{
			bibtexVAL.biblist = NewBibTeX()
		}
	case 3:
		bibtexDollar = bibtexS[bibtexpt-2 : bibtexpt+1]
		//line bibtex.y:37
		{
			bibtexDollar[1].biblist.AddEntry(bibtexDollar[2].bibentry)
			bibtexVAL.biblist = bibtexDollar[1].biblist
		}
	case 4:
		bibtexDollar = bibtexS[bibtexpt-7 : bibtexpt+1]
		//line bibtex.y:40
		{
			bibtexVAL.bibentry = NewBibEntry(bibtexDollar[2].strval, bibtexDollar[4].strval)
			for _, t := range bibtexDollar[6].bibtags {
				bibtexVAL.bibentry.AddField(t.key, t.val)
			}
		}
	case 5:
		bibtexDollar = bibtexS[bibtexpt-0 : bibtexpt+1]
		//line bibtex.y:43
		{
		}
	case 6:
		bibtexDollar = bibtexS[bibtexpt-3 : bibtexpt+1]
		//line bibtex.y:44
		{
			bibtexVAL.bibtag = &bibTag{key: bibtexDollar[1].strval, val: bibtexDollar[3].strval}
		}
	case 7:
		bibtexDollar = bibtexS[bibtexpt-1 : bibtexpt+1]
		//line bibtex.y:47
		{
			bibtexVAL.bibtags = []*bibTag{bibtexDollar[1].bibtag}
		}
	case 8:
		bibtexDollar = bibtexS[bibtexpt-3 : bibtexpt+1]
		//line bibtex.y:48
		{
			bibtexVAL.bibtags = append(bibtexDollar[1].bibtags, bibtexDollar[3].bibtag)
		}
	}
	goto bibtexstack /* stack new state and value */
}
