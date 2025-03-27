package lexemes

import "regexp"

const (
	ThemeBreak_r  = `^( ){0,3}(-{3}|_{3}|\*{3})$`
	Title_r       = `^(#){1,6}( )(.)*$`
	IndentCode_r  = `^(   )( )*.*$`
	FencedCode_r  = "^(`){3}[a-zA-Z]*$"
	BlankLine_r   = `^\s*$`
	Quote_r       = `^( ){0,3}>( )?.*$`
	OrderedList_r = `^( )*[0-9][\.|\)]( )+.*$`
	BulletList_r  = `^( )*[\+|\-|\.]( )+.*$`
)

var (
	ThemeBreakRegexp  = regexp.MustCompile(ThemeBreak_r)
	TitleRegexp       = regexp.MustCompile(Title_r)
	IndentCodeRegexp  = regexp.MustCompile(IndentCode_r)
	FencedCodeRegexp  = regexp.MustCompile(FencedCode_r)
	BlankLineRegexp   = regexp.MustCompile(BlankLine_r)
	QuoteRegexp       = regexp.MustCompile(Quote_r)
	OrderedListRegexp = regexp.MustCompile(OrderedList_r)
	BulletListRegexp  = regexp.MustCompile(BulletList_r)
)

type BlockOptions struct {
	is_raw      bool
	is_terminal bool
}

type Block struct {
	BlockOptions

	name    string
	text    string
	content []Block
}

func NewBlock(name string, text string, content []Block) Block {

	var is_raw bool
	var is_terminal bool

	if name == "IndentCode" || name == "FencedCode" || name == "ThemeBreak" || name == "BlankLine" {
		is_raw = true
	} else {
		is_raw = false
	}

	if name == "OrderedList" || name == "BulletList" || name == "Quote" {
		is_terminal = false
	} else {
		is_terminal = true
	}

	return Block{

		BlockOptions: BlockOptions{
			is_raw:      is_raw,
			is_terminal: is_terminal,
		},

		name:    name,
		text:    text,
		content: content,
	}
}
