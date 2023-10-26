# Trucs important à faire !!!

- Pensez à créer le lien qui lies les salles intermediaire, de départ et de fin. Pour cela utilisé les func 
contenu dans "IA".

-Trier les listes de chemins pour avoir une liste avec 

je veux que tu me trie cette liste :

L1-3 L2-2 L3-1 L4-3 L5-1 
L1-4 L2-5 L3-4 L4-6 L5-2
L2-6 L4-7 L5-6
L2-7 L5-8

Pour qu'elle s'affiche de cette manière:

L1-3 L2-2
L1-4 L2-5 L3-3 L4-3
L2-6 L3-4 L4-6 L5-2
L2-7 L4-7 L5-6
L5-6
L5-8

Maintenant que j'ai ca : 

L1-3 L1-4
L2-2 L2-5 L2-6 L2-7
L3-1 L3-4
L4-3 L4-6 L4-7
L5-1 L5-2 L5-6 L5-8

j'aimerias ca : 

L1-3 L2-2
L1-4 L2-5 L3-3 L4-3
L2-6 L3-4 L4-6 L5-2
L2-7 L4-7 L5-6
L5-6
L5-8


j'aimerais que mes étapes ce stocke de cette manière : 

L1-3
L1-4

puis,

L1-3 L2-2
L1-4 L2-5
L2-6
L2-7

puis,

L1-3 L2-2 
L1-4 L2-5 L3-3
L2-6 L3-4
L2-7

ect...