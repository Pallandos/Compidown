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

	echaped := false

	//TODO : inline lexer

	//parcours des caractères :
	for i, char := range text {

		// échappement
		if echaped {
			echaped = false
			current_text += string(char)
			continue
		}

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

			if len(stack) > 0 {
				// on est dans un inline

				top := stack[len(stack)-1] //sommet

				if top == 'i' {
					// on est dans une image qui passe en attente de description
					// on ne fait rien
				}

			} else {

				// on clos l'inline précédent
				inlines = append(inlines, lexemes.Inline{Genre: "Text", Text: current_text})
				current_text = ""

				// push on ouvre un lien
				stack = append(stack, 'l') // l pour un lien
			}

		case ']':

			if len(stack) == 0 {

				// pas de inline
				current_text += string(char)

			} else {

				top := stack[len(stack)-1] //sommet

				if top == 'l' {
					// on a fini le texte du lien
				} else if top == 'i' {
					// on est dans une image, qui passe en attente d'url
					stack = stack[:len(stack)-1] // pop
					stack = append(stack, '.')   // push d'un . pour indique qu'il ne manque plus que un ) pour fermer l'image
				}
			}

		case '(':
			// on vérifie si on est dans un lien ou dans une image
			if len(stack) > 0 {
				top := stack[len(stack)-1] //sommet

				if top == '.' {
					// on est dans une image, qui passe en attente d'url
					current_text += string(' ')
					stack = stack[:len(stack)-1] // pop
					stack = append(stack, '.')   // push d'un . pour indique qu'il ne manque plus que un ) pour fermer l'image
				} else {
					// on est dans un lien
					current_text += string(' ')
				}
			} else {
				// on n'est pas dans un lien
				current_text += string(char)
			}

		case ')':
			// on vérifie si on est dans un lien
			if len(stack) > 0 {
				top := stack[len(stack)-1] //sommet

				if top == 'l' {
					// on a fini le texte du lien
					inlines = append(inlines, lexemes.Inline{Genre: "Link", Text: current_text})
					current_text = ""
					stack = stack[:len(stack)-1] // pop
				} else if top == '.' {
					// on a fini l'image
					inlines = append(inlines, lexemes.Inline{Genre: "Image", Text: current_text})
					current_text = ""
					stack = stack[:len(stack)-1] // pop
				} else {
					// on n'est pas dans un lien
					current_text += string(char)
				}
			}

		case '!':
			// on est potentiellement dans une image

			if len(stack) != 0 {
				// on est pas dans une image
				current_text += string(char)
			} else {

				if i+1 < len(text) && text[i+1] == '[' {
					// on est dans une image

					// on clos l'inline précédent
					inlines = append(inlines, lexemes.Inline{Genre: "Text", Text: current_text})
					current_text = ""
					// push on ouvre une image
					stack = append(stack, 'i') // i pour une image
				} else {
					// on n'est pas dans une image
					current_text += string(char)
				}
			}

		case '`':
			// dans une emphase

			if len(stack) > 0 {

				top := stack[len(stack)-1]

				if top == '`' {
					// on est dans une emphase
					stack = stack[:len(stack)-1] // pop
					inlines = append(inlines, lexemes.Inline{Genre: "Emphase", Text: current_text})
					current_text = ""

				} else {
					// on n'est pas dans une emphase
					current_text += string(char)
				}

			} else {
				// on clos l'inline précédent
				inlines = append(inlines, lexemes.Inline{Genre: "Text", Text: current_text})
				current_text = ""

				// push on ouvre une emphase
				stack = append(stack, char)
			}

		case '\\':
			// échappement, on ignore le prochain caractère

			if i+1 < len(text) {
				echaped = true
			} else {
				//TODO warning : \ sans caractère suivant
				continue
			}

		default:
			current_text += string(char)
		}
	}

	// on ferme les inlines restants

	if len(current_text) > 0 {
		inlines = append(inlines, lexemes.Inline{Genre: "Text", Text: current_text})
		current_text = ""

		//TODOD warning ? il reste des inlines ouverts
	}

	return inlines
}
