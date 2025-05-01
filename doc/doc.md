# Compidown

Notre projet vise à réaliser un interpréteur de Markdown. Notre objectif est de traduire un fichier source Markdown en HTML pour l'afficher dans un navigateur. 

Les figures imposées sont les suivantes : 
- lexer
- parser
- AST
- visiteur

Nous avons de plus ajouté un *traducteur* qui transforme l'AST en fichier `.html`

## Lexer

Le Lexer est écrit dans le programme [`lexer.go`](../go/pkg/lexer/lexer.go).

Le lexer va parcourir le fichier source et réaliser le découpage en lexemes. Ceux ci sont décrits dans [`lexemes.go`](../go/pkg/lexemes/lexemes.go). On retrouve deux *niveaux* de lexemes : les lines et les inlines.

### Lines

Par lexèmes de lignes, on entend tout ceux qui s'écrivent sur une ligne donc : 

- titres
- citations
- blocs de code
- paragraphes
- listes
- citations

La première étape est donc de parcourir lignes par lignes notre fichier source et d'utiliser des expressions rationnelles pour identifier les lexèmes. 

Expressions rationnelles des lines : 

```go
const (
	ThemeBreak_r  = `^( ){0,3}(-{3}|_{3}|\*{3})$`
	Title_r       = `^(#)+( )(.)*$`
	IndentCode_r  = `^(   )( )*.*$`
	FencedCode_r  = "^(`){3}[a-zA-Z]*( )*$"
	BlankLine_r   = `^\s*$`
	Quote_r       = `^( ){0,3}>( )?.*$`
	OrderedList_r = `^( )*[0-9][\.|\)]( )+.*$`
	BulletList_r  = `^( )*[\+|\-|\.]( )+.*$`
)
```

Une fois identifiés, les lexèmes sont transformés en des objets `Blocks` définis : 

```go
// Informations sur un bloc
type BlockInfos struct {
	IsRaw      bool   `json:"is_raw"`
	IsTerminal bool   `json:"is_terminal"`
	Errors     int    `json:"errors"`
	ErrorsMsg  string `json:"errors_msg"`

	// infos optionelles
	TitleLevel int    `json:"title_level"`
	Language   string `json:"language"`
}

// Structure des blocs de Markdown
type Block struct {
	Genre   string     `json:"Genre"`
	Caracts BlockInfos `json:"caracts"`
	Text    string     `json:"text"`
	Content []*Block   `json:"content"`
}

```

> La partie `json` ne sert que à l'affichage

Prenons l'exemple du texte suivant : 

    ## ceci est un titre

Correspond au bloc suivant : 

```json
{
    "Genre" : "Title" ,
    "Caracts" : {
        "IsRaw" : "False",
        "IsTerminal" : "True",
        "Errors" : "0",
        "ErrorsMsg" : "",

        "TitleLevel" : "2",
        "Language" : ""
    } ,
    "Text" : "ceci est un titre",
    "Content" : ""
}

```

Un bloc non *terminal* est un bloc qui peut inclure d'autres blocs. Ce sont : 

- les listes
- les citations

Si `IsRaw` est faux, alors il faut aussi parser l'intérieur à la recherche d'inlines comme des italiques, gras etc.

### Inlines

Il existe actuellement 5 *inlines* différents : 

- gras
- italiques
- emphases
- liens
- images

Dans le cas où l'intérieur d'un texte devrait aussi être interprété (si `IsRaw` est False), on parcourt le texte avec la fonction `LexerInline` (dans [`lexer.go`](../go/pkg/lexer/lexer.go)). 

La stratégie employée est d'utiliser une *stack* où l'on empile les marqueurs de *inlines* (comme un `*` pour un italique). Le lexer parcourt la ligne carcatère par caractère à la recherche d'un marqueu spécial. Si il s'agit d'un marqueur spécial, on regare le sommet de la pile. Si le sommet de la pile est le même caractère, c'est que l'on clos l'inline en cours. Sinon on ajoute le caractère à la pile et on attends le suivant. 

> Notre compilateur prends en compte le caractère d'échappement `\` pour pouvoir écrire des caractères spéciaux.

Comme pour les lines, les *inlines* sont instancés dans des structures `Inline` : 

```go
type Inline struct {
	Genre string `json:"Genre"`
	Text  string `json:"text"`
}
```

Par exemple la ligne suivant : 

    hello *world* 

est interprété comme une liste de deux structures `Inlines` : 

```json

{
    "Genre" : "Text" ,
    "Text" : "hello"
}

{
    "Genre" : "Italic" ,
    "Text" : "world"
}

```

## Parser

Le parser est défini dans [`parser.go`](../go/pkg/parser/parser.go). Il récupère la liste des blocs produite par le lexer et va construire l'AST. Sa première étape est de vérifier si une erreur est présente dans le texte et de l'afficher si il y en a une.

Par exemple le texte suivant comporte une erreur : 

	this is an example

	re

	####### titre avec trop de #

Le retour affiché est : 

```

2025/04/30 10:53:36 Erreur / dans :  ####### titre avec trop de # / Niveau de titre trop élevé : doit être inférieur à 6
exit status 1


```

Ensuite, les informations sont passées aux visiteurs qui vont construire l'AST.

## AST et visiteur

Un ensemble de visiteurs récursifs vont générer l'AST. La fonction `Parse()` de [`parser.go`](../go/pkg/parser/parser.go) parcourt la liste des lexemes construits et appelle la fonction `Visit()` qui elle même appelle les visiteurs externes. Cette séparation permet une indépendance des fonctions. 

Puisque Go n'est **pas** un langage orienté objet, nous avons dû trouver des méthodes alternatives pour coder les visiteurs. Nous avons donc codé des *méthodes* Go : ce sont des fonctions avec un récepteur : un argument qui permet d'appeler la fonction sur une instance d'un type donné, à la manière d'une méthode.

La structure de défintion d'une telle méthode est la suivante : 

```go
func (node *Node) VisitFencedCode(lexemes []lexemes.Block, indice int) (int, *Node)
```

Ces fonctions construisent récursivement l'AST, qui est une structure de la forme : 

```go 

type AST struct {
	Root *Node
}

type Node struct {
	Genre     string
	Text      string
	Childrens []*Node
	Parent    *Node
}

```

> On retrouve la structure classique d'un arbre, avec l'ajout du double lien avec le parent car on a besoin de "remonter" dans l'arbre lors du parsing

Ensuite, la méthode `Print()` permet d'afficher l'arbre syntaxique en pretty printer. 

Prenons l'exemple du code source suivant : 

``````
this is an example

re

### titre

>wtf
> etet

fiex

    code indenté

- liste
- listeee
finex

```py

bon

```
yeahhhhh
``````

L'AST affiché en pretty printer nous donne : 

```txt
Document: 
    Paragraph: this is an example
    BlankLine: 
    Paragraph: re
    BlankLine: 
    Title:  titre
    BlankLine: 
    Quote: 
        Quote: wtf
        Quote:  etet
    BlankLine: 
    Paragraph: fiex
    BlankLine: 
    IndentCode: code indenté
    BlankLine: 
    BulletList: 
        BulletList: - liste
        BulletList: - listeee
    Paragraph: finex
    BlankLine: 
    FencedCode: py
        RawText: 
        RawText: bon
        RawText: 
    Paragraph: yeahhhhh
```


## Traducteur 

Le traducteur en page html se basant sur les lexèmes transformés en blocs correspond au fichier `html_rpz.go`. On parcourt l'ensemble des blocs contenus dans la liste de lexèmes et pour chacun d'entre eux on appelle une fonction de représentation. Ces fonctions écrivent dans le fichier `output.html` le code html correspondant à la représentation du fichier source markdown en html. La correspondance des deux langages provient de la documentation du *GitHub flavored Markdown*. 

Pour chaque block dont la variable `IsRaw` est false on parse l'intérieur et appelle la fonction de représentation correspondante. Afin de gérer les différentes balises imbriquées les unes dans les autres, chaque ligne du inline est responsable, en fonction de son genre et de celui de la précédente ligne, de fermer ou non la prédente balise et d'ouvrir ou non la sienne.

Par exemple

```
# Chapitre 1 

Il était *une* fois un csn devant **son écran**
```

Sera analysé comme suit 

``````
{Title {false true 0  1 }  Chapitre 1  []}
 
{BlankLine {true true 0  0 }  []}
 
{Paragraph {false true 0  0 } Il était *une* fois un csn devant **son écran** []}
 
{Text Il était }
{Italic une}
{Text  fois un csn devant }
{Bold son écran}
``````

Et écrira dans le fichier de sortie

```html
<!DOCTYPE html><html lang='fr'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><link rel='stylesheet' href='go/pkg/showing/style.css'> </head> 
<h1> Chapitre 1 </h1>
<p>Il était <i>une</i	> fois un csn devant <b>son écran</b>
```

Ainsi on joue sur le `inline.Genre` pour faire appel à la bonne fonction de représentation.

Un css a été écrit afin de rendre le fichier `html_rpz.go` plus lisible et la représentation html plus proche d'une preview markdown, il est écrit en premier par la fonction `Import_style`.