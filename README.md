# HairSwift-back 💈

Backend API pour l'application de gestion de rendez-vous de coiffure HairSwift.

## Présentation

HairSwift est un projet scolaire conçu pour explorer les capacités de Go (Golang) en développement backend, tout en encourageant les développeurs frontend à se plonger dans les technologies backend.
L'application vise à gérer efficacement un salon de coiffure.

## Stack 🖥️

### Développement Back-end avec Go 🛠

Le backend de notre application est développé en utilisant Go, garantissant une robustesse et une efficacité optimales dans le traitement des requêtes serveur.
Go, grâce à sa syntaxe concise et ses performances élevées, s'impose comme le choix idéal pour assurer la fiabilité de notre infrastructure côté serveur.

### Base de données avec Supabase 🗄️

Nous utilisons Supabase comme solution de base de données, offrant une infrastructure PostgreSQL robuste avec des fonctionnalités modernes comme l'authentification, les API RESTful et le stockage en temps réel.

## Fonctionnalités

- Gestion des clients (CRUD)
- Gestion des salons (CRUD)
- Gestion des coiffeurs (CRUD)
- Gestion des créneaux horaires (CRUD)
- Gestion des réservations (CRUD)

## Prérequis

- Go 1.x
- Compte Supabase
- Git

## Installation

1. Clonez le dépôt :

```bash
git clone https://github.com/N95Ryan/hairswift-back.git
```

2. Installez les dépendances :

```bash
go mod download
```

3. Configurez Supabase :

- Créez un projet sur Supabase
- Configurez les variables d'environnement avec vos identifiants Supabase
- Importez le schéma de base de données fourni

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

## 👥 Auteurs

- [Ryan PINA-SILASSE](https://github.com/N95Ryan)
- [Nathan PINARD](https://github.com/YOUGBOY95)
- [Anne-Catherine MICHAUD](https://github.com/annemhd)
