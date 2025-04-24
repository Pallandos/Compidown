# Compidown

by Jean CAYLUS et Jérémy COUTURE

## Introduction

Ce projet vise à réaliser un interpréteur pour le langage **Markdown**. Je me baserais sur le *GitHub flavored Markdown*, car il pose un cadre plus précis que le markdown original. 

La grammaire que j'ai utilisée pour mon projet est lisible [ici](./doc/grammaire.md).

## Déroulement 

Pour la réalisation du projet, j'ai travaillé selon ce schéma : [déroulement](./doc/work.md)

## Compatibilité

Le dossier `/bin` inclut des binaires pour les OS suivants : 

- windows ARM64
- windows AMD64
- darwin (MacOS) AMD64
- darwin (MacOS) ARM64

Il est de plus possible de compiler dans un des systèmes d'exploitations compatibles avec Go. Pour voir la liste des OS et architecture supportés : 

    go tool dist list