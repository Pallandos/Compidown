package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/Pallandos/Compidown/pkg/lexemes"
	"github.com/Pallandos/Compidown/pkg/lexer"
	"github.com/Pallandos/Compidown/pkg/parser"
	html_rpz "github.com/Pallandos/Compidown/pkg/showing"
)

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
	case "windows":
		cmd = "rundll32"
		args = append(args, "url.dll,FileProtocolHandler")
	case "darwin":
		cmd = "open"
	}

	args = append(args, url)
	exec.Command(cmd, args...).Start()
}

func main() {
	// argumments
	var jc bool = false

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file_path> <options>", os.Args[0])
	}

	if len(os.Args) == 3 {
		if os.Args[2] == "--help" {
			log.Println("Usage: go run main.go <file_path> <options>")
			log.Println("Options:")
			log.Println("  --help : Affiche l'aide")
			log.Println("  -jc : Compile sans ouvrir le navigateur")
			return
		} else if os.Args[2] == "-jc" {
			jc = true
		} else {
			log.Fatalf("Option inconnue : %s", os.Args[2])
		}
	}

	file_path := os.Args[1]
	if file_path == "--help" {
		log.Println("Usage: go run main.go <file_path> <options>")
		log.Println("Options:")
		log.Println("  --help : Affiche l'aide")
		log.Println("  -jc : Compile sans ouvrir le navigateur")
		return
	}
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

	ast := parser.Parse(lexeme_list)
	output_file_path := "../output.html"
	html_rpz.ShowBlock2(output_file_path, lexeme_list)
	ast.Print()

	if !jc {
		// ouverture du navigateur
		openBrowser(output_file_path)
	} else {
		log.Println("Compilation terminée, fichier de sortie : " + output_file_path)
	}
}
