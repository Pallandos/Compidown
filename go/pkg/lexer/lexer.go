package lexer

import (
	"github.com/Pallandos/Compidown/pkg/lexemes"
)

func LexerBlock(line string) lexemes.Block {

	// teste des différents cas de lexèmes :

	if lexemes.BlankLineRegexp.MatchString(line) {

		return (lexemes.NewBlock("BlankLine", "", nil))

	} else if lexemes.TitleRegexp.MatchString(line) {

		return (lexemes.NewBlock("Title", line, nil))

	} else if lexemes.ThemeBreakRegexp.MatchString(line) {

		return (lexemes.NewBlock("ThemeBreak", "", nil))

	} else if lexemes.FencedCodeRegexp.MatchString(line) {

		return (lexemes.NewBlock("FencedCode", line, nil))

	} else if lexemes.IndentCodeRegexp.MatchString(line) {

		return (lexemes.NewBlock("IndentCode", line, nil))

	} else if lexemes.QuoteRegexp.MatchString(line) {

		return (lexemes.NewBlock("Quote", line, nil))

	} else if lexemes.OrderedListRegexp.MatchString(line) {

		return (lexemes.NewBlock("OrderedList", line, nil))

	} else if lexemes.BulletListRegexp.MatchString(line) {

		return (lexemes.NewBlock("BulletList", line, nil))

	} else {

		return (lexemes.NewBlock("Paragraph", line, nil))

	}
}

func LexerInline(text string) []lexemes.Inline {

	stack := []rune{}
	inlines := []lexemes.Inline{}

	current_text := ""

	//TODO : inline lexer

	//parcours des caractères :
	for i, char := range text {

		switch char {

		case '*':
			if len(stack) > 0 {
				// dans ce cas, on est soit dans un italique, soit dans un gras

				top := stack[len(stack)-1] //sommet

				if top == '*' {

					// on vérifie si on commence un gras
					if text[i-1] == '*' {
						// on commence un gras

						// pop du dernier * pour le remplacer par un b qui marque le gras COMPLET
						stack = stack[:len(stack)-1]
						stack = append(stack, 'b')

					} else {

						// nous sommes dans un italique
						stack = stack[:len(stack)-1] // pop
						inlines = append(inlines, lexemes.Inline{Genre: "Italic", Text: current_text})
						current_text = ""
					}
				} else if top == 'b' {

					// on attends la fin d'un gras COMPLET
					stack = stack[:len(stack)-1] // pop
					stack = append(stack, '_')   // push d'un _ pour indique qu'il ne manque plus que un * pour fermer le gras

				} else if top == '_' {

					// on avait un gras partiel, on le ferme
					stack = stack[:len(stack)-1] // pop
					inlines = append(inlines, lexemes.Inline{Genre: "Bold", Text: current_text})
					current_text = ""

				}

			} else {
				// on clos l'inline précédent
				inlines = append(inlines, lexemes.Inline{Genre: "Text", Text: current_text})
				current_text = ""

				// push on ouvre un italique
				stack = append(stack, char)
			}

		case '[':

		case ']', ')':

			if len(stack) == 0 {

				// pas de inline
				current_text += string(char)

			} else {

				// top := stack[len(stack)-1]   //sommet
				stack = stack[:len(stack)-1] //pop

			}

		default:
			current_text += string(char)
		}
	}

	// on ferme les inlines restants
	inlines = append(inlines, lexemes.Inline{Genre: "Text", Text: current_text})

	return inlines
}
