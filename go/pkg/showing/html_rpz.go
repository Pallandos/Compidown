package main

import (
	"os"
	"strconv"
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

func Import_style() {
	file, err := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close() // on ferme automatiquement à la fin de notre programme
	_, err = file.WriteString("<!DOCTYPE html><html lang='fr'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><link rel='stylesheet' href='style.css'> </head>")
	if err != nil {
		panic(err)
	}
}

func main() {
	Import_style()
	Rpz_title("Carotte", 2)
	Rpz_blankline()
	Rpz_title("Patate", 6)
	Rpz_themebreak()
	Rpz_fencedcode("func main() {", "go")
}
