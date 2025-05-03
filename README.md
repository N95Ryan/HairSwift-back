# HairSwift 💈

Backend API pour l'application de gestion de rendez-vous de coiffure HairSwift.

## Description du projet 👨‍🏫

HairSwift est une application backend développée dans le cadre d'un projet scolaire. Elle vise à simplifier la gestion d'un salon de coiffure en offrant une API robuste pour gérer les rendez-vous, les clients, les coiffeurs, les salons et les créneaux horaires. Ce projet a été conçu pour explorer les capacités de Go (Golang) en développement backend tout en encourageant les développeurs frontend à se familiariser avec les technologies backend.

L'application permet une gestion efficace des opérations d'un salon grâce à une architecture performante et une base de données moderne.

## Fonctionnalités ✨

- **Gestion des clients** : Création, lecture, mise à jour et suppression (CRUD).
- **Gestion des salons** : Administration des informations des salons (CRUD).
- **Gestion des coiffeurs** : Gestion des profils des coiffeurs (CRUD).
- **Gestion des créneaux horaires** : Planification et modification des disponibilités (CRUD).
- **Gestion des réservations** : Prise, modification et annulation de rendez-vous (CRUD).

## Stack technique 🖥️

### Back-end 🛠️

Le backend de HairSwift repose sur les technologies suivantes :

- **Go (Golang)**  
  Le choix de [Go](https://golang.org/) garantit une performance élevée et une syntaxe concise, idéal pour un serveur robuste et fiable. Sa simplicité facilite la maintenance et l'évolutivité du code.

- **Supabase**  
  [Supabase](https://supabase.com/) est utilisé comme solution de base de données PostgreSQL. Il offre des fonctionnalités modernes telles que l'authentification intégrée, des API RESTful et un stockage en temps réel, simplifiant la gestion des données.

### Repository GitHub 📂

- Backend : [https://github.com/N95Ryan/hairswift-back](https://github.com/N95Ryan/hairswift-back)
- Frontend : [https://github.com/N95Ryan/hairswift](https://github.com/N95Ryan/hairswift)

## Prérequis 📋

Avant de commencer, assurez-vous d'avoir installé :

- [Go](https://golang.org/doc/install) (version 1.16 ou supérieure)
- Un compte [Supabase](https://supabase.com/)
- [Git](https://git-scm.com/)

## Installation 📥

Suivez ces étapes pour configurer et exécuter le projet localement :

1. **Clonez le dépôt** :

   ```bash
   git clone https://github.com/N95Ryan/hairswift-back.git
   cd hairswift-back
   ```

2. **Installez les dépendances** :

   ```bash
   go mod download
   ```

3. **Configurez Supabase** :

   - Créez un projet sur [Supabase](https://supabase.com/).
   - Configurez les variables d'environnement dans un fichier `.env` avec vos identifiants Supabase (par exemple, `SUPABASE_URL` et `SUPABASE_KEY`).
   - Importez le schéma de base de données fourni dans le projet.

4. **Lancez l'application** :
   ```bash
   go run main.go
   ```

## Contribution 🤝

Nous accueillons toutes les contributions ! Pour contribuer :

1. Forkez le projet.
2. Créez une branche pour votre fonctionnalité (`git checkout -b feature/nouvelle-fonctionnalite`).
3. Commitez vos changements (`git commit -m "Ajout de nouvelle fonctionnalité"`).
4. Poussez vers la branche (`git push origin feature/nouvelle-fonctionnalite`).
5. Ouvrez une Pull Request.

## Auteurs 👥

- [@Ryan PINA-SILASSE](https://github.com/N95Ryan)
- [@Nathan PINARD](https://github.com/YOUGBOY95)
- [@Anne-Catherine MICHAUD](https://github.com/annemhd)
