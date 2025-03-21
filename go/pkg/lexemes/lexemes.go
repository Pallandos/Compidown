package lexemes

import "regexp"

const (
	ThemeBreak = `( ){0,3}(-{3}|_{3}|\*{3})`
	Title      = `(#){1,6}( )(.)*`
	IndentCode = `(   )( )*.*`
	FencedCode = "(`){3}[a-zA-Z]*"
	BlankLine  = `\s*`
	Quote      = `( ){,3}>( )?.*`
)

var (
	ThemeBreakRegexp = regexp.MustCompile(ThemeBreak)
	TitleRegexp      = regexp.MustCompile(Title)
	IndentCodeRegexp = regexp.MustCompile(IndentCode)
	FencedCodeRegexp = regexp.MustCompile(FencedCode)
	BlankLineRegexp  = regexp.MustCompile(BlankLine)
	QuoteRegexp      = regexp.MustCompile(Quote)
)
