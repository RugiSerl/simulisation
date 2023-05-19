# simulisation
simulisation est une contraction de simulation et civilisation

ce projet ressemble à un automate cellulaire, de part son fonctionnement. En revanche, ici il n'est pas question de "cellules" à proprement parler, les entités sont fixées selon des coordonnées réelles.

### Comment exécuter le programme : 

- afin d'exécuter le projet, il est possible de télécharger les exécutables déjà compilés dans la section releases

- <p>pour compiler le projet, il faut tout d'abord se munir du compiler de go : https://go.dev/dl/ .<br/><br/>Ensuite, il faudra installer TDM-GCC (si le compiler gcc n'est pas déjà présent sur la machine) : https://jmeubank.github.io/tdm-gcc/ <br/><br/>Enfin, il faut exécuter dans un terminal "go build", pour compiler, ou bien "go run main.go" pour exécuter le code directement<br/><i>il est possible de compiler en enlevant l'invite de commande en faisant "go build -ldflags -H=windowsgui .\main.go".</i><br/><br/>lors de la première compilation, la dépendance raylib-go se téléchargera automatiquement, ce qui peut prendre quelques minutes</p>

### Mode d'emploi :

- un clic droit permet d'invoquer une entité à l'emplacement de la souris
- rester appuyer sur shift gauche permet de faire invoquer des entités en continu
- les flèches directionnelles permettent de se déplacer dans le monde et la molette de la souris permet de régler le zoom
- afin d'ouvir les paramètres, il est possible d'appuyer sur échap ou bien d'appuyer sur le bouton présent en haut à droite



