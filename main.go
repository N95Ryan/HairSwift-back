package main

import (
	"fmt"
	"log"
	"net/http"

	"hairswift-back/database"
	"hairswift-back/handlers"
)

func main() {
	// Initialisation du client Supabase
	client, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialisation des handlers avec le client Supabase
	handlers.InitHandlers(client)

	// Routes pour les clients
	http.HandleFunc("/api/clients", handlers.GetClientsHandler)
	http.HandleFunc("/api/clients/add", handlers.AddClientHandler)
	http.HandleFunc("/api/clients/update", handlers.UpdateClientHandler)
	http.HandleFunc("/api/clients/delete", handlers.DeleteClientHandler)

	// Routes pour les salons
	http.HandleFunc("/api/salons", handlers.GetSalonsHandler)
	http.HandleFunc("/api/salons/add", handlers.AddSalonHandler)
	http.HandleFunc("/api/salons/update", handlers.UpdateSalonHandler)
	http.HandleFunc("/api/salons/delete", handlers.DeleteSalonHandler)

	// Routes pour les coiffeurs
	http.HandleFunc("/api/coiffeurs", handlers.GetCoiffeursHandler)
	http.HandleFunc("/api/coiffeur/add", handlers.AddCoiffeurHandler)
	http.HandleFunc("/api/coiffeur/update", handlers.UpdateCoiffeurHandler)
	http.HandleFunc("/api/coiffeur/delete", handlers.DeleteCoiffeurHandler)

	// Routes pour les créneaux
	http.HandleFunc("/api/creneaux", handlers.GetCreneauxHandler)
	http.HandleFunc("/api/creneaux/add", handlers.AddCreneauHandler)
	http.HandleFunc("/api/creneaux/update", handlers.UpdateCreneauHandler)
	http.HandleFunc("/api/creneaux/delete", handlers.DeleteCreneauHandler)

	// Routes pour les réservations
	http.HandleFunc("/api/reservations", handlers.GetReservationsHandler)
	http.HandleFunc("/api/reservations/add", handlers.AddReservationHandler)
	http.HandleFunc("/api/reservations/update", handlers.UpdateReservationHandler)
	http.HandleFunc("/api/reservations/delete", handlers.DeleteReservationHandler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
