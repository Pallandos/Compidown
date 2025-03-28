package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Pallandos/Compidown/pkg/lexer"
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
		retour := lexer.LexerInline(scanner.Text())

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
}
