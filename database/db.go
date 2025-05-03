package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	supa "github.com/supabase-community/supabase-go"
)

var (
	supabaseClient *supa.Client
	clientsMu      sync.RWMutex
	salonsMu       sync.RWMutex
	coiffeursMu    sync.RWMutex
	creneauxMu     sync.RWMutex
	reservationsMu sync.RWMutex
)

func InitDB() (*supa.Client, error) {
	// Charger les variables d'environnement depuis le fichier .env
	// Ignorer l'erreur si le fichier n'existe pas (pour les tests)
	_ = godotenv.Load()

	// Récupérer les informations de Supabase
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		return nil, fmt.Errorf("SUPABASE_URL or SUPABASE_KEY is not set")
	}

	// Créer le client Supabase
	client, err := supa.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing Supabase client: %v", err)
	}

	supabaseClient = client
	return client, nil
}

func CreateTables() error {
	// Les tables sont gérées directement dans l'interface Supabase
	// Cette fonction reste pour la compatibilité mais n'a plus besoin de créer les tables
	return nil
}

func GetDB() *supa.Client {
	return supabaseClient
}

func GetMutexes() (*sync.RWMutex, *sync.RWMutex, *sync.RWMutex, *sync.RWMutex, *sync.RWMutex) {
	return &clientsMu, &salonsMu, &coiffeursMu, &creneauxMu, &reservationsMu
}
