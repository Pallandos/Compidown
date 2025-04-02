# Lexer Compidown

Ce document décrit la stratégie du lexer.

## Lexer des *inlines* 

On utlise une pile pour les ouvertures / fermetures de inlines. J'utilise des marqueurs customs pour me simplifier la tâche.

Marqueurs : 

- `*` : italique 
- `b` : gras complet 
- `_` : gras en attente de fermeture : le caractère suivant DOIT être une étoile pour fermer le gras
- `l` : lien
- `i` : une image en formation ( attente de description)
- `.` : image en cours de cloture