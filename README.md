# Exercice de développement

Le but de cet exercice est d'apporter un socle de discussion pour l'entretien technique que nous aurons ensemble.
Il n'y a pas de contraintes de temps mais l'exercice est assez simple et ne devrait idéalement pas dépasser 1h. N'hésitez pas à poser vos questions à `entretien-atlas-crit@groupe-crit.com` si nécessaire.

## Énoncé

On souhaite rajouter un nouvel endpoint à notre application permettant de retourner les informations pour une agence donnée.

Exemple à titre indicatif:
```
Entrée:
code agence == "431"

Sortie attendue:
"CRIT SELESTAT"
```

Le format de l'entrée et de la sortie est laissée en libre cours au developpeur.


## Lancement du code

**Assurez-vous de disposer d'une version de Go supérieure ou égale à la version 1.16.**

Lancez la commande suivante pour démarrer le serveur HTTP :
```shell
go run main.go
```

Par défaut le serveur écoute sur le port 8080.
Pour changer le port utilisez l'option `-port` :
```shell
go run main.go -port 8000
```

Vous pouvez tester le bon fonctionnement de l'URL http://localhost:8080/healthcheck avec un navigateur ou en ligne de commandes.

## Retour attendu
Une archive avec l'exercice fait. N'hésitez pas à rajouter autant d'informations qu'il vous semble utile de rajouter que ce soit dans le code ou en commentaire en considérant que vous travaillez à proposer une Pull Request à l'équipe.
