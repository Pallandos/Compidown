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

##### Titres

Un **titre** en Markdown, ou **ATX heading** est un block qui ne contient pas d'éléments. Il peuvent être de 6 niveaux d'importance différents.

On le définit par 1 à 6 éléments hashtag `#`, suivi d'un espace nécessaire, suivi du nom du titre lui même. On peut aussi clore la séquence par d'autres `#` (peu importe leur nombre) mais cela n'est pas nécessaire. 

##### Setext

???

##### Bloc de codes indentés
