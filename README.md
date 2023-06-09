# Simulisation
Simulisation est une contraction de simulation et civilisation

Ce projet ressemble à un automate cellulaire, de part son fonctionnement. En revanche, ici il n'est pas question de "cellules" à proprement parler, les entités sont fixées selon des coordonnées réelles.




### Comment exécuter/compiler le programme : 

- Afin d'exécuter le projet, il est possible de télécharger les exécutables déjà compilés dans la section releases

- <p>Pour compiler le projet, il faut tout d'abord se munir du compiler de go : https://go.dev/dl/ .<br/><br/>Ensuite, il faudra installer TDM-GCC (si le compilateur gcc n'est pas déjà présent sur la machine) : https://jmeubank.github.io/tdm-gcc/ <br/><br/>Enfin, il faut exécuter dans un terminal "go build", pour compiler, ou bien "go run main.go" pour exécuter le code directement<br/><i>il est possible de compiler en enlevant l'invite de commande en faisant "go build -ldflags -H=windowsgui .\main.go".</i><br/><br/>lors de la première compilation, la dépendance raylib-go se téléchargera automatiquement, ce qui peut prendre quelques minutes</p>

### Mode d'emploi :

- Un clic droit permet d'invoquer une entité à l'emplacement de la souris
- Rester appuyer sur shift gauche permet de faire invoquer des entités en continu
- Les flèches directionnelles permettent de se déplacer dans le monde et la molette de la souris permet de régler le zoom
- Afin d'ouvir les paramètres, il est possible d'appuyer sur échap ou bien d'appuyer sur le bouton présent en haut à droite
- maintenir clic droit permet de supprimer les entités proches de la souris (le rayon est modifiable dans les paramètres)
- Appuyer sur suppr permet de supprimer toutes les entités
- Appuyer sur f11 permet de mettre le jeu en mode plein écran (et de l'enlever)
- il est possible de clear l'arrière plan du jeu avec "backspace", si l'option "nettoyer les résidus" est activée
- Pour ajouter un matériaux, il est possible de sélectionner "Insérer un matériau", pour basculer le type de matériau, il faut appuyer sur G


### Fonctionnement du jeu

<p>
Ce jeu est constitué d'une liste d'entité, qui ont chacune des propriétés telle que leur position, leur âge et leur valeur Morale. Cette dernière est une valeur abstraite, utilisée pour les interactions entres les entités. Elle va de 0 à 255 et est cyclique, donc la "distance morale" entre 4 et 253 est de 6.<br/>
</p>

A chaque mise à  jour du jeu, les entités sont mises à jour, qui se divise en 5 parties :
- Une fonction de déplacement, qui fait déplacer l'entité et la fait se rapprocher des autres
- Une fonction de "décollision", où l'entité s'éloigne des autres afin d'éviter qu'elles s'empilent toutes en un point
- Une fonction de reproduction
- Une fonction qui met à jour l'âge de l'entité
- Une fonction qui décide s'il faut tuer une entité

Chacune de ces fonctions peut avoir plusieurs versions qui peuvent changer le comportement du jeu.



