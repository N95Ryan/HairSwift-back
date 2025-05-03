package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hairswift-back/models"
	"hairswift-back/services"

	supa "github.com/supabase-community/supabase-go"
)

type ClientController struct {
	clientService *services.ClientService
}

func NewClientController(client *supa.Client) *ClientController {
	return &ClientController{
		clientService: services.NewClientService(client),
	}
}

func (c *ClientController) AddClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newClient models.Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		log.Printf("Erreur lors du décodage du JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON invalide"})
		return
	}

	insertedClient, err := c.clientService.CreateClient(newClient)
	if err != nil {
		log.Printf("Erreur lors de la création du client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erreur lors de la création du client"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedClient)
}

func (c *ClientController) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := c.clientService.GetAllClients()
	if err != nil {
		log.Printf("Erreur lors de la récupération des clients: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erreur lors de la récupération des clients"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func (c *ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedClient models.Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		log.Printf("Erreur lors du décodage du JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON invalide"})
		return
	}

	updated, err := c.clientService.UpdateClient(updatedClient)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erreur lors de la mise à jour du client"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (c *ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id_client")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("ID invalide: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalide"})
		return
	}

	err = c.clientService.DeleteClient(id)
	if err != nil {
		log.Printf("Erreur lors de la suppression du client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erreur lors de la suppression du client"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client supprimé avec succès"})
}
