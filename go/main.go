package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Pallandos/Compidown/pkg/lexer"
	html_rpz "github.com/Pallandos/Compidown/pkg/showing"
)

func main() {
	file_path := "../source/src.txt"

	// ouverture
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatalf("Erreur d'ouverture : %v", err)
	}
	defer file.Close() // fermeture

	// lecture
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		retour := lexer.LexerBlock(scanner.Text())

		json, err := json.Marshal(retour)

		if err != nil {
			log.Fatalf("Erreur de conversion en JSON : %v", err)
			return
		}

		fmt.Println(string(json))
	}

	// gestion des erreurs de lecture
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erreur de lecture : %v", err)
	}

	// tests des inlines :
	inlines := lexer.LexerInline("ceci *est* un text avec un **petit** peu de italique de gras et un [lien](youhou.com) et une image ![image](lien.com)")
	html_rpz.Showtext(inlines)
	fmt.Println(inlines)
}
