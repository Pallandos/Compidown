# Projet COMPIDOWN

Comme dans le cours, je suivrais les étapes suivantes : 

1. Lexer
2. Parser
3. AST
4. Visitor

# Lexer

La première étape dans le développement du lexer est de réaliser une liste **exhaustive** des lexèmes existants. Puis on écrira des expressions rationelles correspondants à ces lexèmes. 

## Lexèmes

Suit ici une liste qui se veut exhaustive des lexèmes de Markdown. Il est important de considérer que les éléments en Markdown peuvent être séparés en 2 catégories : **blocks** et **inline**. Un *block* peut être peut dans un cas contenir d'autres blocs : il est dit **container**, ou bien être **terminal**. 

### Les **blocks**

Comme dit plus haut, les blocs peuvent dans un cas contenir d'autres blocs, dans d'autres cas non. 

#### Bloc terminal 

Un bloc terminal est un bloc qui ne peux contenir d'autres blocs, et qui occupe seul toute sa ligne. 

##### Break line

Une **break line **est une ligne simple. Elle est composée de 0 à 3 espaces suivis de 3 caractères (ou plus) identiques parmi : `-`, `_` et `*`.

Une REGEX python correspondante pourrait être : 

    "( ){0,3}(-{3}|_{3}|\*{3})"

##### Titres

Un **titre** en Markdown, ou **ATX heading** est un block qui ne contient pas d'éléments. Il peuvent être de 6 niveaux d'importance différents.

On le définit par 1 à 6 éléments hashtag `#`, suivi d'un espace nécessaire, suivi du nom du titre lui même. On peut aussi clore la séquence par d'autres `#` (peu importe leur nombre) mais cela n'est pas nécessaire. 

Une REGEX python qui décrit ce lexème est :

    "(#){1,6}( )(.)*"

##### Setext

???

##### Bloc de code indenté

Un bloc de code indenté est composé d'une indentation (ou plus) suivi du texte.

##### Fenced code block

Un **fenced code block** est un bloc de code qui s'étire sur plusieurs lignes. Il commence par au moins 3 `` ` `` consécutifs. Il se termine lorsque le même nombre de *ticks* est placé sur une ligne. On peut, dans certains cas, placer des caractères après les premiers *ticks*, pour indiquer le langage. 

Une REGEX python qui décrit ce lexème :

    "(`){3}[a-zA-Z]*\n(.*\n)*(`){3}"

Il faut néanmoins décrire l'ouverture du bloc : 

    "(`){3}[a-zA-Z]*"

et sa fermeture : 

    "(`){3}"

##### HTML blocks

Les blocs HTML sont des éléments qui seront traités comme du HTML et donc PAS comme du Markdown. Il existe de nombreux types de blocs HTML et sont décrits ci après : 

1. **start** : l'un de ces blocs : `<script`, `<pre` ou `<style` suivi d'espaces éventuels, le tout fermé par `>`
    
    **end**  : le bloc de fermeture correspondant, à savoir `</script>`, `</pre>` ou `</style>`.
2. 

##### Paragraphes


Les **paragraphes** sont tous les blocs qui ne peuvent être interprétés comme autre chose. Le seul moyen de terminer un paragraphe est de rajouter une ligne vide. Rajouter plus d'une ligne vide n'a aucun effet.