package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Pallandos/Compidown/pkg/lexemes"
	"github.com/Pallandos/Compidown/pkg/lexer"
	"github.com/Pallandos/Compidown/pkg/parser"
)

func main() {
	// argumments

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file_path>", os.Args[0])
	}

	file_path := os.Args[1]
	lexeme_list := []lexemes.Block{}

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

		//json, err := json.Marshal(retour)

		if err != nil {
			log.Fatalf("Erreur de conversion en JSON : %v", err)
			return
		}
		lexeme_list = append(lexeme_list, retour)
	}

	// gestion des erreurs de lecture
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erreur de lecture : %v", err)
	}

	// affichage
	//fmt.Println(lexeme_list)
	ast := parser.Parse(lexeme_list)
	// html_rpz.ShowBlock(lexeme_list)
	//html_rpz.Showtext(inlines)
	ast.Print()
}
