// construit l'AST à partir de la liste de lexemes
// parse la liste des lexemes

package parser

import (
	"log"
	"os"
	"strings"

	"github.com/Pallandos/Compidown/pkg/lexemes"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

type Node struct {
	Genre     string
	Text      string
	Childrens []*Node
	Parent    *Node
}

type AST struct {
	Root *Node
}

func NewAST() *AST {
	return &AST{
		Root: &Node{
			Genre:     "Document",
			Text:      "",
			Childrens: []*Node{},
			Parent:    nil,
		},
	}
}

func (n *Node) AddChild(genre string, text string) {
	child := &Node{
		Genre:     genre,
		Text:      text,
		Childrens: []*Node{},
		Parent:    n,
	}
	n.Childrens = append(n.Childrens, child)
}

// Parse la liste des lexemes et construit l'AST
func Parse(lexemes []lexemes.Block) *AST {
	ast := NewAST()
	currentNode := ast.Root
	// deepth := 0

	for i := 0; i < len(lexemes); {
		lexeme := lexemes[i]

		// on vérifie si il y a une erreur :
		if lexeme.Caracts.Errors > 0 {
			log.Fatalf("%sErreur%s / dans :  %v / %s%v%s", colorRed, colorReset, lexeme.Text, colorGreen, lexeme.Caracts.ErrorsMsg, colorReset)
		}

		// // on monte si c'est une blank line et qu'on est déja descendu
		// if lexeme.Genre == "BlankLine" && deepth > 0 {
		// 	// on monte d'un niveau
		// 	if currentNode.Parent != nil {
		// 		currentNode = currentNode.Parent
		// 		deepth--
		// 	}
		// 	i++
		// 	continue
		// }
		switch lexeme.Caracts.IsTerminal {
		case true:

			// Si le lexeme est un bloc terminal, on l'ajoute comme enfant du noeud courant
			currentNode.AddChild(lexeme.Genre, lexeme.Text)

		case false:

			// Si le lexeme est un bloc non terminal, on crée un nouveau noeud et on l'ajoute comme enfant du noeud courant
			newNode := &Node{
				Genre:     lexeme.Genre,
				Text:      "",
				Childrens: []*Node{},
				Parent:    currentNode,
			}

			currentNode.Childrens = append(currentNode.Childrens, newNode)
			currentNode = newNode

			i, currentNode = currentNode.Visit(lexemes, i, lexeme.Genre)
			continue
		}
		i++
	}
	return ast
}

// Print l'AST dans un fichier
func (ast *AST) Print() {

	print_path := "../out/print.txt"

	file, err := os.Create(print_path)
	if err != nil {
		log.Fatalf("Erreur de création du fichier : %v", err)
	}

	// pour l'affichage chaque enfant sera affiché avec une indentation en plus
	var printNode func(node *Node, depth int)

	printNode = func(node *Node, depth int) {
		indentation := ""
		for i := 0; i < depth; i++ {
			indentation += "    "
		}
		_, err := file.WriteString(indentation + node.Genre + ": " + node.Text + "\n")
		if err != nil {
			log.Fatalf("Erreur d'écriture dans le fichier : %v", err)
		}
		for _, child := range node.Childrens {
			printNode(child, depth+1)
		}
	}

	printNode(ast.Root, 0)

	err = file.Close()
	if err != nil {
		log.Fatalf("Erreur de fermeture du fichier : %v", err)
	}
}

// --------- visiteurs externes --------------

func (node *Node) Visit(lexemes []lexemes.Block, indice int, genre string) (int, *Node) {

	switch genre {
	case "FencedCode":
		return node.VisitFencedCode(lexemes, indice)
	case "Quote":
		return node.VisitQuote(lexemes, indice)
	case "BulletList":
		return node.VisitBulletList(lexemes, indice)
	case "OrderedList":
		return node.VisitOrderedList(lexemes, indice)
	default:
		return indice, node
	}
}

func (node *Node) VisitFencedCode(lexemes []lexemes.Block, indice int) (int, *Node) {

	for i := indice; i < len(lexemes); i++ {
		lexeme := lexemes[i]

		if i == indice {
			// on ajoute le language au noeud parent

			language := strings.TrimLeft(lexeme.Text, "`")
			node.Text += language

		} else if lexeme.Genre == "FencedCode" {
			// on est à la fin du bloc

			return i + 1, node.Parent
		} else {
			// on ajoute le lexeme au noeud courant
			node.AddChild("RawText", lexeme.Text)
		}
	}

	// douteux :
	return len(lexemes), node
}

func (node *Node) VisitQuote(lexemes []lexemes.Block, indice int) (int, *Node) {
	for i := indice; i < len(lexemes); i++ {
		lexeme := lexemes[i]

		if lexeme.Genre == "BlankLine" {
			// on est à la fin du bloc

			return i, node.Parent
		} else {
			// on ajoute le lexeme au noeud courant
			node.AddChild(lexeme.Genre, lexeme.Text)
		}
	}

	// douteux :
	return len(lexemes), node
}

func (node *Node) VisitBulletList(lexemes []lexemes.Block, indice int) (int, *Node) {
	for i := indice; i < len(lexemes); i++ {
		lexeme := lexemes[i]

		if lexeme.Genre != "BulletList" {
			// on est à la fin du bloc

			return i, node.Parent
		} else {
			// on ajoute le lexeme au noeud courant
			node.AddChild(lexeme.Genre, lexeme.Text)
		}
	}

	// douteux :
	return len(lexemes), node
}

func (node *Node) VisitOrderedList(lexemes []lexemes.Block, indice int) (int, *Node) {
	for i := indice; i < len(lexemes); i++ {
		lexeme := lexemes[i]

		if lexeme.Genre != "OrderedList" {
			// on est à la fin du bloc

			return i, node.Parent
		} else {
			// on ajoute le lexeme au noeud courant
			node.AddChild(lexeme.Genre, lexeme.Text)
		}
	}

	// douteux :
	return len(lexemes), node
}
