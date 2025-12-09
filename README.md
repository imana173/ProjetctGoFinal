Mini-CRM – Gestion de contacts en ligne de commande

Ce projet consiste à créer une application CLI en Go pour gérer des contacts (Mini-CRM).
L’objectif était de transformer progressivement un petit programme utilisant des maps en une application complète, modulaire et persistante.

Fonctionnalités

Ajouter un contact

Lister les contacts

Mettre à jour un contact

Supprimer un contact

Choisir le type de stockage (mémoire, JSON, SQLite)

CLI professionnelle avec Cobra

Configuration externe via Viper

Architecture
contact-cli/
│
├── cmd/                 → Commandes CLI (add, list, update, delete)
│
├── internal/
│   ├── models/          → Struct Contact
│   └── storer/          → Interface Storer + 3 stockages :
│       - MemoryStorer
│       - JSONStorer
│       - GORMStorer (SQLite)
│
├── config.yaml          → Choix du storer
├── main.go              → Point d’entrée
└── go.mod               → Module Go

Installation

Dans un terminal, exécuter :

go mod tidy

Configuration

Le fichier config.yaml permet de choisir le système de stockage :

storer: "json"     # "memory", "json", "gorm"
json_path: "contacts.json"
db_path: "contacts.db"

Commandes disponibles
➤ Ajouter un contact

Interactif :

go run main.go add


Avec flags :

go run main.go add --name="Imane" --email="test@mail.com"

➤ Lister les contacts
go run main.go list

➤ Mettre à jour un contact
go run main.go update --id=1 --name="Nouveau Nom" --email="new@mail.com"

➤ Supprimer un contact
go run main.go delete --id=1

Storers implémentés

Mémoire : aucune persistance

JSON : stockage dans un fichier JSON

SQLite (GORM) : stockage dans une base de données SQLite

Le storer utilisé dépend de config.yaml.
