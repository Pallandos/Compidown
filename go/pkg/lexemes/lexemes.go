// Décrit les lexemes de Markdown
//
// La première partie traite des lexemes de blocs
package lexemes

import (
	"regexp"
	"strings"
)

const (
	ThemeBreak_r  = `^( ){0,3}(-{3}|_{3}|\*{3})$`
	Title_r       = `^(#)*( )(.)*$`
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

// Informations sur un bloc
type BlockInfos struct {
	IsRaw      bool   `json:"is_raw"`
	IsTerminal bool   `json:"is_terminal"`
	Errors     int    `json:"errors"`
	ErrorsMsg  string `json:"errors_msg"`

	// infos optionelles
	TitleLevel int    `json:"title_level"`
	Language   string `json:"language"`
}

// Structure des blocs de Markdown
type Block struct {
	Name    string     `json:"name"`
	Caracts BlockInfos `json:"caracts"`
	Text    string     `json:"text"`
	Content []Block    `json:"content"`
}

// Crée un nouveau bloc
func NewBlock(name string, text string, content []Block) Block {

	var is_raw bool
	var is_terminal bool
	var errors int = 0
	var errors_msg string = ""

	// infos optionelles
	var title_level int = 0
	var language string = ""

	// séparation des cas :

	switch name {

	case "Title":
		title_level = len(text) - len(strings.TrimLeft(text, "#"))

		if title_level > 6 {
			// les titres sont d'un niveau entre 1 et 6
			errors++
			errors_msg += "Niveau de titre trop élevé : doit être inférieur à 6\n"
			title_level = 0
			name = "Paragraph"
		} else {
			text = strings.TrimLeft(text, "#")
		}

	}

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

		Name: name,

		Caracts: BlockInfos{
			IsRaw:      is_raw,
			IsTerminal: is_terminal,
			Errors:     errors,
			ErrorsMsg:  errors_msg,

			// infos optionelles
			TitleLevel: title_level,
			Language:   language,
		},

		Text:    text,
		Content: content,
	}
}
