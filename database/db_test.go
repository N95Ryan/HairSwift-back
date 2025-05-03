package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Définir directement les variables d'environnement pour les tests
	os.Setenv("SUPABASE_URL", "https://ftnqaicrlvzvdgyeyphm.supabase.co")
	os.Setenv("SUPABASE_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZ0bnFhaWNybHZ6dmRneWV5cGhtIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDYyNjc2NTksImV4cCI6MjA2MTg0MzY1OX0.JGF83IFJDbhNDJJJr_7WHKHkMtipBjc1XOU18Jw1xY4")

	// Exécuter les tests
	code := m.Run()
	os.Exit(code)
}

func TestInitDB(t *testing.T) {
	// Vérifier que les variables d'environnement sont définies
	if os.Getenv("SUPABASE_URL") == "" || os.Getenv("SUPABASE_KEY") == "" {
		t.Skip("Les variables d'environnement SUPABASE_URL et SUPABASE_KEY ne sont pas définies")
	}

	// Tester l'initialisation du client Supabase
	client, err := InitDB()
	if err != nil {
		t.Fatalf("Échec de l'initialisation du client Supabase: %v", err)
	}

	// Vérifier que le client n'est pas nil
	if client == nil {
		t.Fatal("Le client Supabase est nil")
	}
}

func TestGetDB(t *testing.T) {
	// Vérifier que les variables d'environnement sont définies
	if os.Getenv("SUPABASE_URL") == "" || os.Getenv("SUPABASE_KEY") == "" {
		t.Skip("Les variables d'environnement SUPABASE_URL et SUPABASE_KEY ne sont pas définies")
	}

	// Initialiser le client Supabase
	_, err := InitDB()
	if err != nil {
		t.Fatalf("Échec de l'initialisation du client Supabase: %v", err)
	}

	// Tester la récupération du client
	client := GetDB()
	if client == nil {
		t.Fatal("GetDB() a retourné nil")
	}
}
