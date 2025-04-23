# HairSwift-back

Backend API pour l'application de gestion de rendez-vous de coiffure HairSwift.

## Description

HairSwift-back est une API REST développée en Go qui gère les fonctionnalités principales d'un système de réservation de rendez-vous de coiffure. L'API permet de gérer les clients, les salons, les coiffeurs, les créneaux horaires et les réservations.

## Fonctionnalités

- Gestion des clients (CRUD)
- Gestion des salons (CRUD)
- Gestion des coiffeurs (CRUD)
- Gestion des créneaux horaires (CRUD)
- Gestion des réservations (CRUD)

## Prérequis

- Go 1.x
- MySQL
- Git

## Installation

1. Clonez le dépôt :

```bash
git clone [URL_DU_REPO]
```

2. Installez les dépendances :

```bash
go mod download
```

3. Configurez la base de données MySQL :

- Créez une base de données MySQL
- Modifiez les paramètres de connexion dans le code source

4. Lancez l'application :

```bash
go run main.go
```

## Structure du Projet

- `main.go` : Point d'entrée de l'application
- `go.mod` : Fichier de gestion des dépendances
- `go.sum` : Fichier de vérification des dépendances

## API Endpoints

### Clients

- `POST /clients` : Créer un nouveau client
- `GET /clients` : Récupérer tous les clients
- `PUT /clients/{id}` : Mettre à jour un client
- `DELETE /clients/{id}` : Supprimer un client

### Salons

- `POST /salons` : Créer un nouveau salon
- `GET /salons` : Récupérer tous les salons
- `PUT /salons/{id}` : Mettre à jour un salon
- `DELETE /salons/{id}` : Supprimer un salon

### Coiffeurs

- `POST /coiffeurs` : Créer un nouveau coiffeur
- `GET /coiffeurs` : Récupérer tous les coiffeurs
- `PUT /coiffeurs/{id}` : Mettre à jour un coiffeur
- `DELETE /coiffeurs/{id}` : Supprimer un coiffeur

### Créneaux

- `POST /creneaux` : Créer un nouveau créneau
- `GET /creneaux` : Récupérer tous les créneaux
- `PUT /creneaux/{id}` : Mettre à jour un créneau
- `DELETE /creneaux/{id}` : Supprimer un créneau

### Réservations

- `POST /reservations` : Créer une nouvelle réservation
- `GET /reservations` : Récupérer toutes les réservations
- `PUT /reservations/{id}` : Mettre à jour une réservation
- `DELETE /reservations/{id}` : Supprimer une réservation

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :

1. Fork le projet
2. Créer une branche pour votre fonctionnalité
3. Commiter vos changements
4. Pousser vers la branche
5. Ouvrir une Pull Request

## Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.
