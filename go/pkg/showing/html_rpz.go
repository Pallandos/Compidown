package html_rpz

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Pallandos/Compidown/pkg/lexemes"
)

func Rpz_title(line string, level int) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                                                                           // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<h" + strconv.Itoa(level) + ">" + line + "</h" + strconv.Itoa(level) + "><HR>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
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

func Rpz_fencedcode(line string, language string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                                                      // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<div class='fenced_code_language'>" + language + "</div>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	defer file.Close()                                                               // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<div class='fenced_code_block'>" + line + "</div>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_intentedcode(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                                  // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<pre><code>" + line + "</code></pre>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_quote(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                                   // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<blockquote>" + line + "</blockquote>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_orderedlist(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                         // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString(line + "<br>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func Rpz_bulletlist(line string) {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()                                   // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<li>" + line + "</li>\n") // écrire dans le fichier
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
	defer file.Close()                                 // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<p>" + line + "</p>\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func showblock(line lexemes.Inline) {
	// on affiche le bloc
	fmt.Println("Genre : " + line.Genre)
	fmt.Println("Texte : " + line.Text)
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
	fmt.Println("Affichage des inlines :")
	Import_style()
	for _, inline := range lines {
		showblock(inline)
	}
	Rpz_fencedcode("fencedcode", "python")
	Rpz_intentedcode("intentedcode")
	Rpz_quote("quote")
	Rpz_title("title", 1)
	Rpz_themebreak()
	Rpz_blankline()
	Rpz_bulletlist("bulletlist")
	Rpz_orderedlist("orderedlist")
	Rpz_paragraph("paragraph")

}

func ShowBlock(lines []lexemes.Block) {
	fmt.Println("Affichage des blocs :")
	Import_style()
	var last string = "First"
	var temp string = ""
	for _, block := range lines {
		fmt.Println("///////")
		if last != "BlankLine" && last != "First" {
			fmt.Println("Genre dans le non Blankline : " + block.Genre)
			temp = last
			last = temp
		}
		fmt.Println("Genre : " + block.Genre)
		fmt.Println("Texte : " + block.Text)
		switch block.Genre {
		case "Title":
			Rpz_title(block.Text, block.Caracts.TitleLevel)
		case "BlankLine":
			continue
		case "ThemeBreak":
			Rpz_themebreak()
		case "FencedCode":
			Rpz_fencedcode(block.Text, block.Caracts.Language)
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
			fmt.Println("Genre inconnu")
		}
		last = block.Genre
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
