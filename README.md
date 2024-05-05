# Projet de Développement : Application Web de Recommandation de Films

## Description

Ce projet consiste en une application web de recommandation de films, avec une architecture client-serveur. Le backend est développé en utilisant le langage Go (Golang) pour créer une API REST. Le frontend est construit en utilisant Svelte, un framework JavaScript moderne, et Vite comme outil de build. 

## Technologies Utilisées

- **Backend (API REST en Golang)**
  - Langage de Programmation : Go (Golang)
  - Framework pour le Serveur Web : Standard Library
  - Test Unitaires : Go Test avec Mocking de la Base de Données
  - Base de Données : PostgreSQL
  - Migration de la Base de Données : Liquibase

- **Frontend (Application Web en Svelte)**
  - Framework JavaScript : Svelte
  - Outil de Build : Vite
  - Tests Unitaires : Vitest

## Tests Unitaires

Les tests unitaires pour le backend sont réalisés avec Go Test, avec une utilisation de mocking pour la base de données. Pour le frontend, les tests unitaires sont réalisés avec Vitest.

## Lancer la Stack localement 

- Ajouter les variables d'env dans le .env de movie-recommendation-frontend

`VITE_API_URL= localhost:8080`  
`VITE_TMDB_TOKEN=` token api de tmdb

`docker compose up --build `
