package main

import (
	"fmt"
	"log"
	"net/http"

	"hairswift-back/controllers"
	"hairswift-back/database"
)

func main() {
	// Initialisation du client Supabase
	client, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialisation des controllers
	clientController := controllers.NewClientController(client)
	salonController := controllers.NewSalonController(client)
	coiffeurController := controllers.NewCoiffeurController(client)
	creneauController := controllers.NewCreneauController(client)
	reservationController := controllers.NewReservationController(client)

	// Routes pour les clients
	http.HandleFunc("/api/clients", clientController.GetClients)
	http.HandleFunc("/api/clients/add", clientController.AddClient)
	http.HandleFunc("/api/clients/update", clientController.UpdateClient)
	http.HandleFunc("/api/clients/delete", clientController.DeleteClient)

	// Routes pour les salons
	http.HandleFunc("/api/salons", salonController.GetSalons)
	http.HandleFunc("/api/salons/add", salonController.AddSalon)
	http.HandleFunc("/api/salons/update", salonController.UpdateSalon)
	http.HandleFunc("/api/salons/delete", salonController.DeleteSalon)

	// Routes pour les coiffeurs
	http.HandleFunc("/api/coiffeurs", coiffeurController.GetCoiffeurs)
	http.HandleFunc("/api/coiffeurs/add", coiffeurController.AddCoiffeur)
	http.HandleFunc("/api/coiffeurs/update", coiffeurController.UpdateCoiffeur)
	http.HandleFunc("/api/coiffeurs/delete", coiffeurController.DeleteCoiffeur)

	// Routes pour les créneaux
	http.HandleFunc("/api/creneaux", creneauController.GetCreneaux)
	http.HandleFunc("/api/creneaux/add", creneauController.AddCreneau)
	http.HandleFunc("/api/creneaux/update", creneauController.UpdateCreneau)
	http.HandleFunc("/api/creneaux/delete", creneauController.DeleteCreneau)

	// Routes pour les réservations
	http.HandleFunc("/api/reservations", reservationController.GetReservations)
	http.HandleFunc("/api/reservations/add", reservationController.AddReservation)
	http.HandleFunc("/api/reservations/update", reservationController.UpdateReservation)
	http.HandleFunc("/api/reservations/delete", reservationController.DeleteReservation)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
