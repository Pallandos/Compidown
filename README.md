# Compidown

by Jean CAYLUS et Jérémy COUTURE

## Introduction

Ce projet vise à réaliser un interpréteur pour le langage **Markdown**. Je me baserais sur le *GitHub flavored Markdown*, car il pose un cadre plus précis que le markdown original. 

La grammaire que j'ai utilisée pour mon projet est lisible [ici](./doc/grammaire.md).

### Pourquoi Go

Ce projet est écrit en [Go](https://go.dev/), un langage de programmation développé par Google. Nous avons choisi Go car c'est un langage simple, à la fois à la lecture et à l'apprentissage; mais aussi pour sa rapidité d'écriture. En effet, Go est un langage compilé, très inspiré du C tout en simplifiant sa syntaxe et en incluant davantage de fonctions natives. 

Et puis, nous avons aussi agi par curiosité. En mars 2025, Microsoft annonce réécrire le noyau de Typescript en Go, et annonce une vitesse d'execution multipliée par 10! Intrigués, nous avons décidé de découvrir ce langage. 


## Déroulement 

Pour la réalisation du projet, j'ai travaillé selon ce schéma : [déroulement](./doc/work.md)

Au moment du rendu de ce projet, notre compilateur prends en charge tous les items de la [grammaire](./doc/grammaire.md) décrite, à l'exception de : 

- les tableaux : très complexes à afficher en HTML
- les blocs HTML : assez difficiles à délimités, et rarement utilisés 

## Compatibilité

Le dossier `/bin` inclut des binaires pour les OS suivants : 

- windows ARM64
- windows AMD64
- darwin (MacOS) AMD64
- darwin (MacOS) ARM64

Il est de plus possible de compiler dans un des systèmes d'exploitations compatibles avec Go. Pour voir la liste des OS et architecture supportés : 

    go tool dist list