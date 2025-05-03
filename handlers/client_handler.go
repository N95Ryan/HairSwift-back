package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hairswift-back/database"
	"hairswift-back/models"
)

func AddClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newClient models.Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	clientsMu, _, _, _, _ := database.GetMutexes()

	result, err := db.Exec("INSERT INTO clients (firstname, lastname, email, password) VALUES (?, ?, ?, ?)",
		newClient.Firstname, newClient.Lastname, newClient.Email, newClient.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newClient.ID_client = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClient)
}

func GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	clientsMu, _, _, _, _ := database.GetMutexes()
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	db := database.GetDB()
	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clientList []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID_client, &client.Firstname, &client.Lastname, &client.Email, &client.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		clientList = append(clientList, client)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientList)
}

func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedClient models.Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	clientsMu, _, _, _, _ := database.GetMutexes()
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	row := db.QueryRow("SELECT id_client FROM clients WHERE id_client=?", updatedClient.ID_client)
	if err := row.Scan(&updatedClient.ID_client); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE clients SET firstname=?, lastname=?, email=?, password=? WHERE id_client=?",
		updatedClient.Firstname, updatedClient.Lastname, updatedClient.Email, updatedClient.Password, updatedClient.ID_client)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedClient)
}

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id_client")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	clientsMu, _, _, _, _ := database.GetMutexes()
	clientsMu.Lock()
	defer clientsMu.Unlock()

	row := db.QueryRow("SELECT id_client FROM clients WHERE id_client=?", id)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM clients WHERE id_client=?", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
