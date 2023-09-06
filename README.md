# Go-mutateur

Permet de controller plusieurs ordinateur en réseau avec une seul souris et un seul clavier

## Installation

Pour pouvoir lancer le projet il faut avoir [Go](https://go.dev/) installé et suivre l'installation de [robotgo](https://github.com/go-vgo/robotgo).

## Lancement

```bash
go run . [server/client]
```

## Configuration

La configuration se fait dans le fichier [config.yml](config.yml)

```yaml
server:
  port: 8082 //port sur lequel le server ecoute

client:
  port: 8082 //port sur lequel le client va se connecter
  address: 192.168.0.37 //address sur lequel le client va se connecter

config:
  smooth_mode: false //permet de rendre le déplacement de la souris plus "humain" (WIP)
  smooth_delay: 10 //permet de gerer le délais des deplacements si l'option si dessus est défini sur true
```

## Comment utiliser

Ce petit programme permet de répercuter les actions d'un ordinateur sur un ou plusieurs autres. Il est cross plateforme et marche sur Windows/Linux/MacOS

Appuyer sur "Echap" pour sortir du controle et pouvoir retourner au menu de selection.

Les action effectuers sur l'ordinateur "serveur" sont effectuer sur le/s client/s mais aussi sur lui même. Si vous ne voulez pas que des action se fasse en parrallele sur votre ordinateur je vous invite a mettre en fullscreen le terminal sur lequelle le programme est lancé.

## La suite

Les fonctionalités prévue sont:

- ouverture d'une fenetre en fullscreen afin de ne pas avoir a mettre le termninal en grand pour eviter d'effectuer les action sur le PC hote
- Ajout d'une option pour partager l'ecran de l'ordinateur controlé

## Contribution

N'hésitez pas a faire des retour sur les possible bug et amélioration a apporté au projet.

## BUG connu

Drag ne marche pas encore.
