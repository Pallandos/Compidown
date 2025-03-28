package lexer

import (
	"github.com/Pallandos/Compidown/pkg/lexemes"
)

func LexerInline(line string) lexemes.Block {

	// teste des différents cas de lexèmes :

	if lexemes.BlankLineRegexp.MatchString(line) {

		return (lexemes.NewBlock("BlankLine", "", nil))

	} else if lexemes.TitleRegexp.MatchString(line) {

		return (lexemes.NewBlock("Title", line, nil))

	} else if lexemes.ThemeBreakRegexp.MatchString(line) {

		return (lexemes.NewBlock("ThemeBreak", "", nil))

	} else if lexemes.FencedCodeRegexp.MatchString(line) {

		return (lexemes.NewBlock("FencedCode", line, nil))

	} else if lexemes.IndentCodeRegexp.MatchString(line) {

		return (lexemes.NewBlock("IndentCode", line, nil))

	} else if lexemes.QuoteRegexp.MatchString(line) {

		return (lexemes.NewBlock("Quote", line, nil))

	} else if lexemes.OrderedListRegexp.MatchString(line) {

		return (lexemes.NewBlock("OrderedList", line, nil))

	} else if lexemes.BulletListRegexp.MatchString(line) {

		return (lexemes.NewBlock("BulletList", line, nil))

	} else {

		return (lexemes.NewBlock("Paragraph", line, nil))

	}
}
