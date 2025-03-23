package lexer

import (
	"github.com/Pallandos/Compidown/pkg/lexemes"
)

func LexerInline(line string) string {

	// teste des différents cas de lexèmes :

	if lexemes.BlankLineRegexp.MatchString(line) {
		return ("blank")
	} else if lexemes.TitleRegexp.MatchString(line) {
		return ("title")
	} else if lexemes.ThemeBreakRegexp.MatchString(line) {
		return ("theme")
	} else if lexemes.FencedCodeRegexp.MatchString(line) {
		return ("code")
	} else if lexemes.IndentCodeRegexp.MatchString(line) {
		return ("indent code")
	} else if lexemes.QuoteRegexp.MatchString(line) {
		return ("quote")
	} else if lexemes.OrderedListRegexp.MatchString(line) {
		return ("ordered list")
	} else if lexemes.BulletListRegexp.MatchString(line) {
		return ("bullet list")
	} else {
		return ("paragraph")
	}
}
