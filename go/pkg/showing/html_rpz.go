package html_rpz

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Pallandos/Compidown/pkg/lexemes"
	"github.com/Pallandos/Compidown/pkg/lexer"
)

func Rpz_title(line string, level int) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                          // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<h" + strconv.Itoa(level) + ">") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	for _, inline := range lexer.LexerInline(line) {
		if inline.Genre == "Bold" {
			Rpz_bold(inline.Text, file)
		}
		if inline.Genre == "Italic" {
			Rpz_italic(inline.Text, file)
		}
		if inline.Genre == "Text" {
			file.WriteString(inline.Text)
		}
	}
}

func Rpz_blankline() {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                  // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<br>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_themebreak() {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                  // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<HR>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_fencedcode_language(language string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                                                          // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<div class='fenced_code_language'>" + language[3:] + "</div>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	defer file.Close()                                           // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<div class='fenced_code_block'>") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_fencedcode(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()              // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(line) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_intentedcode(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                              // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<pre><code>" + line) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_quote(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                               // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<blockquote>" + line) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_unopened_quote(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()              // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(line) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_orderedlist(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                  // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(line[2:]) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_bulletlist(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                           // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<li>" + line[2:]) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_bold(line string, file *os.File) {
	file.WriteString("<b>" + line + "</b>")
}

func Rpz_italic(line string, file *os.File) {
	file.WriteString("<i>" + line + "</i	>")
}

func Rpz_paragraph(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()               // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<p>") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	for _, inline := range lexer.LexerInline(line) {
		if inline.Genre == "Bold" {
			Rpz_bold(inline.Text, file)
		}
		if inline.Genre == "Italic" {
			Rpz_italic(inline.Text, file)
		}
		if inline.Genre == "Text" {
			file.WriteString(inline.Text)
		}
		if inline.Genre == "Link" {
			lastSpaceIndex := strings.LastIndex(inline.Text, " ")
			beforeLastSpace := inline.Text[:lastSpaceIndex]  // Substring before the last space
			afterLastSpace := inline.Text[lastSpaceIndex+1:] // Substring after the last space

			file.WriteString("<a href='" + afterLastSpace + "'>" + beforeLastSpace + "</a>")
		}
		if inline.Genre == "Image" {
			file.WriteString("<img src='" + inline.Text + "' alt='" + inline.Text + "'>")
		}
	}
}

func Rpz_unopened_paragraph(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()              // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(line) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_space() {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()             // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(" ") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func findcloser(block lexemes.Block) string {
	switch block.Genre {
	case "Title":
		var retour string = ("</h" + strconv.Itoa(block.Caracts.TitleLevel) + ">\n")
		return (retour)
	case "BlankLine":
		return ("")
	case "ThemeBreak":
		return ("")
	case "FencedCode":
		return ("</div>\n")
	case "IndentCode":
		return ("</code></pre>\n")
	case "Quote":
		return ("</blockquote>\n")
	case "OrderedList":
		return ("<br>\n")
	case "BulletList":
		return ("</li>\n")
	case "Paragraph":
		return ("</p>\n")
	}
	return ("")
}

func writecloser(closer string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(closer) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func showblock(line lexemes.Inline) {
	switch line.Genre {
	case "Title":
		// Rpz_title(line.Text, line.caracts.TitleLevel)
	case "BlankLine":
		Rpz_blankline()
	case "ThemeBreak":
		Rpz_themebreak()
	case "FencedCode":
		//Rpz_fencedcode(line.Text, line.Caracts.Language)
	case "IndentCode":
		Rpz_intentedcode(line.Text)
	case "Quote":
		Rpz_quote(line.Text)
	case "OrderedList":
		Rpz_orderedlist(line.Text)
	case "BulletList":
		Rpz_bulletlist(line.Text)
	case "Paragraph":
		Rpz_paragraph(line.Text)
	}
}

func Showtext(lines []lexemes.Inline) {
	Import_style()
	for _, inline := range lines {
		showblock(inline)
	}
	//Rpz_fencedcode("fencedcode", "python")
	Rpz_intentedcode("intentedcode")
	Rpz_quote("quote")
	Rpz_title("title", 1)
	Rpz_themebreak()
	Rpz_blankline()
	Rpz_bulletlist("bulletlist")
	Rpz_orderedlist("orderedlist")
	Rpz_paragraph("paragraph")

}

func ShowBlock2(lines []lexemes.Block) {
	Import_style()
	last := lexemes.Block{
		Genre:   "First",
		Caracts: lexemes.BlockInfos{IsRaw: false, IsTerminal: true, Errors: 0, ErrorsMsg: "", TitleLevel: 2, Language: ""},
		Text:    "This is an example title",
		Content: nil, // No nested blocks
	}
	var is_fency_open bool = false
	for _, block := range lines {
		fmt.Println(block)
		fmt.Println(" ")
		if is_fency_open {
			if block.Genre == "FencedCode" {
				is_fency_open = false
				writecloser(findcloser(block))
				continue
			}
			Rpz_fencedcode(block.Text)
			continue
		}

		if last.Genre == block.Genre {
			if block.Genre == "Quote" {
				Rpz_unopened_quote(block.Text)
				continue
			}
			if block.Genre == "Paragraph" {
				Rpz_paragraph(block.Text)
				continue
			}
		}

		writecloser(findcloser(last))

		switch block.Genre {
		case "Title":
			Rpz_title(block.Text, block.Caracts.TitleLevel)
		case "BlankLine":
			last = block
			continue
		case "ThemeBreak":
			Rpz_themebreak()
		case "FencedCode":
			if is_fency_open {
				is_fency_open = false
			} else {
				writecloser(findcloser(block))
				Rpz_fencedcode_language(block.Text)
				is_fency_open = true
				continue
			}
		case "IndentCode":
			Rpz_intentedcode(block.Text)
		case "Quote":
			Rpz_quote(block.Text)
		case "OrderedList":
			Rpz_orderedlist(block.Text)
		case "BulletList":
			Rpz_bulletlist(block.Text)
		case "Paragraph":
			Rpz_paragraph(block.Text)
		default:
		}
		last = block
	}
}

func Import_style() {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close() // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<!DOCTYPE html><html lang='fr'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><link rel='stylesheet' href='pkg/showing/style.css'> </head> \n")
	if err != nil {
		panic(err)
	}
}
