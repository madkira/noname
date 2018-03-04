# noname
projet pour le cours objet connecté et services


## Deploiement

### Serveur

Notre serveur à été deployer en go, il vous faudra donc installer <a href="https://golang.org/">go</a> sur votre machine, grace à votre gestionnaire de paquet favoris (je ne peut que conseiller l'utilisation de <a href="https://wiki.archlinux.fr/yaourt">yaourt</a>), une fois cela fais, rendez vous dans le dossier contenant le fichier Server.go et executez la commande "go build", cela crée un executable, qu'il ne vous reste plus qu'à executer (cela nécessitera des droits sudo, pour écouter sur les ports). Votre serveur tourne a présent sur votre machine ! Bien joué !

### Raspberry
installer les modules suivants :
```
sudo apt install pico2wave aplay python3 python-pip3 mosquitto  mosquitto-clients
```
installer les modules python :
```
sudo pip3 install paho
```
pour executer lancer les deux scripts python :
```
python3 script.py & python3 script_read.py &
```
La detection des clefs est faite par NFC avec un lecteur et un tag acroché aux clefs, pour la demo des tags NFC Type 4 sont utilisés, le type 4 est couramment utilisé pour des applications du type Smart Card, l'avantage du type 4 est qu'il supporte l'anticolision et permet donc de détecter plusieurs tags dans le champ d'un lecteur, dans notre cas ce serais la possibilité de detecter plusieur clefs de différentes persone. Le faite d'utiliser ce type de tag nous permet aussi de faire evoluer la platforme en stockent des informations utilisateur dans les tags des clefs.
