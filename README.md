# Présentation de HairSwift 💈

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

- Go
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

## Utilisation de Postman pour tester l'API

### Configuration de Postman

1. **Création d'une collection**
   - Ouvrez Postman
   - Créez une nouvelle collection nommée "HairSwift API"
   - Ajoutez les variables d'environnement suivantes :
     ```
     base_url: http://localhost:8080
     ```

### Endpoints et exemples de requêtes

#### Clients

1. **Créer un client**

   ```
   POST {{base_url}}/clients
   Content-Type: application/json

   {
       "firstname": "John",
       "lastname": "Doe",
       "email": "john.doe@example.com",
       "password": "password123"
   }
   ```

2. **Obtenir tous les clients**

   ```
   GET {{base_url}}/clients
   ```

3. **Mettre à jour un client**

   ```
   PUT {{base_url}}/clients/{id}
   Content-Type: application/json

   {
       "firstname": "John",
       "lastname": "Doe",
       "email": "john.doe.updated@example.com",
       "password": "newpassword123"
   }
   ```

4. **Supprimer un client**
   ```
   DELETE {{base_url}}/clients/{id}
   ```

#### Salons

1. **Créer un salon**

   ```
   POST {{base_url}}/salons
   Content-Type: application/json

   {
       "name": "Salon de beauté"
   }
   ```

2. **Obtenir tous les salons**

   ```
   GET {{base_url}}/salons
   ```

3. **Mettre à jour un salon**

   ```
   PUT {{base_url}}/salons/{id}
   Content-Type: application/json

   {
       "name": "Nouveau nom du salon"
   }
   ```

4. **Supprimer un salon**
   ```
   DELETE {{base_url}}/salons/{id}
   ```

#### Coiffeurs

1. **Créer un coiffeur**

   ```
   POST {{base_url}}/coiffeurs
   Content-Type: application/json

   {
       "id_salon": 1,
       "firstname": "Marie",
       "lastname": "Dupont"
   }
   ```

2. **Obtenir tous les coiffeurs**

   ```
   GET {{base_url}}/coiffeurs
   ```

3. **Mettre à jour un coiffeur**

   ```
   PUT {{base_url}}/coiffeurs/{id}
   Content-Type: application/json

   {
       "id_salon": 1,
       "firstname": "Marie",
       "lastname": "Dupont-Updated"
   }
   ```

4. **Supprimer un coiffeur**
   ```
   DELETE {{base_url}}/coiffeurs/{id}
   ```

#### Créneaux

1. **Créer un créneau**

   ```
   POST {{base_url}}/creneaux
   Content-Type: application/json

   {
       "id_coiffeur": 1,
       "date_creneau": "2024-03-20T10:00:00Z",
       "availability": true
   }
   ```

2. **Obtenir tous les créneaux**

   ```
   GET {{base_url}}/creneaux
   ```

3. **Mettre à jour un créneau**

   ```
   PUT {{base_url}}/creneaux/{id}
   Content-Type: application/json

   {
       "id_coiffeur": 1,
       "date_creneau": "2024-03-20T11:00:00Z",
       "availability": false
   }
   ```

4. **Supprimer un créneau**
   ```
   DELETE {{base_url}}/creneaux/{id}
   ```

#### Réservations

1. **Créer une réservation**

   ```
   POST {{base_url}}/reservations
   Content-Type: application/json

   {
       "id_salon": 1,
       "id_coiffeur": 1,
       "id_creneau": 1
   }
   ```

2. **Obtenir toutes les réservations**

   ```
   GET {{base_url}}/reservations
   ```

3. **Mettre à jour une réservation**

   ```
   PUT {{base_url}}/reservations/{id}
   Content-Type: application/json

   {
       "id_salon": 1,
       "id_coiffeur": 1,
       "id_creneau": 2
   }
   ```

4. **Supprimer une réservation**
   ```
   DELETE {{base_url}}/reservations/{id}
   ```

### Conseils pour les tests

1. **Ordre des tests** :

   - Commencez par créer un salon
   - Puis créez un coiffeur associé à ce salon
   - Ensuite, créez des créneaux pour ce coiffeur
   - Enfin, créez des réservations

2. **Vérification des réponses** :

   - Les réponses sont au format JSON
   - Codes de statut HTTP :
     - 200 : Requêtes GET réussies
     - 201 : Créations réussies
     - 204 : Suppressions réussies
     - 400 : Erreurs de validation
     - 500 : Erreurs serveur

3. **Gestion des erreurs** :

   - En cas d'erreur 500, vérifiez les logs du serveur
   - En cas d'erreur 400, vérifiez le format de vos données

4. **Variables d'environnement** :
   - Assurez-vous que votre fichier `.env` est correctement configuré
   - Vérifiez que le serveur est bien démarré avec `go run main.go`
