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

Dans la logique globale du lexe, celui ci regarde quel est le prochain caractère. Si il s'agit d'un caractère spécial (comme `*` pour un italique), on vérifie que les conditions sont réunies pour valider le bloc spécial et si c'est le cas on introduit dans notre pile le marqueur du bloc spécial correspondant. Ensuite, la prochaine fois qu'on tombe sur un caractère spécial, on va regarder le sommet de la pile pour vérifier si on dois clore le bloc spécial en cours. Si c'est le cas on le ferme et on ajoute le bloc aux blocs précédents, sinon on ouvre un nouveau bloc et on continue le processus. 