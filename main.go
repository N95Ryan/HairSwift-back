package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// TYPES
type Client struct {
	ID_client int    `json:"id_client"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Salon struct {
	ID_salon int    `json:"id_salon"`
	Name     string `json:"name"`
}

type Coiffeur struct {
	ID_coiffeur int    `json:"id_coiffeur"`
	ID_salon    int    `json:"id_salon"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
}

type Creneau struct {
	ID_creneau   int    `json:"id_creneau"`
	ID_coiffeur  int    `json:"id_coiffeur"`
	Date         string `json:"date_creneau"`
	Availability bool   `json:"availability"`
}

type Reservation struct {
	ID_reservation int `json:"id_reservation"`
	ID_salon       int `json:"id_salon"`
	ID_coiffeur    int `json:"id_coiffeur"`
	ID_creneau     int `json:"id_creneau"`
}

var (
	db             *sql.DB
	clientsMu      sync.RWMutex
	salonsMu       sync.RWMutex
	coiffeursMu    sync.RWMutex
	creneauxMu     sync.RWMutex
	reservationsMu sync.RWMutex
	nextID         = 1
)

// / MAIN
func main() {
	/// BASE DE DONNÉES
	var err error
	db, err = sql.Open("mysql", "goteam:root@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	/// TABLES
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			id_client INT AUTO_INCREMENT PRIMARY KEY,
			firstname VARCHAR(150),
			lastname VARCHAR(150),
			email VARCHAR(150),
			password VARCHAR(255)
		);
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS salons (
			id_salon INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(150)
		);
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS coiffeurs (
			id_coiffeur INT AUTO_INCREMENT PRIMARY KEY,
			id_salon INT,
			firstname VARCHAR(150),
			lastname VARCHAR(150)
		);
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS creneaux (
			id_creneau INT AUTO_INCREMENT PRIMARY KEY,
			id_coiffeur INT,
			date_creneau VARCHAR(150),
			availability BOOLEAN
		);
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS reservations (
			id_reservation INT AUTO_INCREMENT PRIMARY KEY,
			id_salon INT,
			id_coiffeur INT,
			id_creneau INT
		);
    `)
	if err != nil {
		log.Fatal(err)
	}

	/// ROUTES
	/// Clients
	http.HandleFunc("/api/clients", getClientsHandler)
	http.HandleFunc("/api/clients/add", addClientHandler)
	http.HandleFunc("/api/clients/update", updateClientHandler)
	http.HandleFunc("/api/clients/delete", deleteClientHandler)

	/// Salons
	http.HandleFunc("/api/salons", getSalonsHandler)
	http.HandleFunc("/api/salons/add", addSalonHandler)
	http.HandleFunc("/api/salons/update", updateSalonHandler)
	http.HandleFunc("/api/salons/delete", deleteSalonHandler)

	/// Coiffeurs
	http.HandleFunc("/api/coiffeurs", getCoiffeursHandler)
	http.HandleFunc("/api/coiffeur/add", addCoiffeurHandler)
	http.HandleFunc("/api/coiffeur/update", updateCoiffeurHandler)
	http.HandleFunc("/api/coiffeur/delete", deleteCoiffeurHandler)

	/// Creneaux
	http.HandleFunc("/api/creneaux", getCreneauxHandler)
	http.HandleFunc("/api/creneaux/add", addCreneauHandler)
	http.HandleFunc("/api/creneaux/update", updateCreneauHandler)
	http.HandleFunc("/api/creneaux/delete", deleteCreneauHandler)

	/// Reservations
	http.HandleFunc("/api/reservations", getReservationsHandler)
	http.HandleFunc("/api/reservations/add", addReservationHandler)
	http.HandleFunc("/api/reservations/update", updateReservationHandler)
	http.HandleFunc("/api/reservations/delete", deleteReservationHandler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// CLIENTS
func addClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newClient Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO clients (firstname, lastname, email, password) VALUES (?, ?, ?, ?)", newClient.Firstname, newClient.Lastname, newClient.Email, newClient.Password)
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

func getClientsHandler(w http.ResponseWriter, r *http.Request) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	// Fetch users from the database
	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clientList []Client
	for rows.Next() {
		var client Client
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

func updateClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedClient Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	_, err = db.Exec("UPDATE clients SET firstname=?, lastname=?, email=?, password=? WHERE id_client=?", updatedClient.Firstname, updatedClient.Lastname, updatedClient.Email, updatedClient.Password, updatedClient.ID_client)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedClient)
}

func deleteClientHandler(w http.ResponseWriter, r *http.Request) {
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

// SALONS
func addSalonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newSalon Salon
	err := json.NewDecoder(r.Body).Decode(&newSalon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO salons (name) VALUES (?)", newSalon.Name)
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

	newSalon.ID_salon = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSalon)
}

func getSalonsHandler(w http.ResponseWriter, r *http.Request) {
	salonsMu.RLock()
	defer salonsMu.RUnlock()

	// Fetch users from the database
	rows, err := db.Query("SELECT * FROM salons")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var salonList []Salon
	for rows.Next() {
		var salon Salon
		err := rows.Scan(&salon.ID_salon, &salon.Name)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		salonList = append(salonList, salon)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salonList)
}

func updateSalonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedSalon Salon
	err := json.NewDecoder(r.Body).Decode(&updatedSalon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	salonsMu.RLock()
	defer salonsMu.RUnlock()
	row := db.QueryRow("SELECT id_salon FROM salons WHERE id_salon=?", updatedSalon.ID_salon)
	if err := row.Scan(&updatedSalon.ID_salon); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE salons SET name=? WHERE id_salon=?", updatedSalon.Name, updatedSalon.ID_salon)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSalon)
}

func deleteSalonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id_salon")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	salonsMu.Lock()
	defer salonsMu.Unlock()
	row := db.QueryRow("SELECT id_salon FROM salons WHERE id_salon=?", id)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM salons WHERE id_salon=?", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// COIFFEURS
func addCoiffeurHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newCoiffeur Coiffeur
	err := json.NewDecoder(r.Body).Decode(&newCoiffeur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO coiffeurs (id_salon, firstname, lastname) VALUES (?, ?, ?)", newCoiffeur.ID_coiffeur, newCoiffeur.Firstname, newCoiffeur.Lastname)
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

	newCoiffeur.ID_coiffeur = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCoiffeur)
}

func getCoiffeursHandler(w http.ResponseWriter, r *http.Request) {
	coiffeursMu.RLock()
	defer coiffeursMu.RUnlock()

	// Fetch users from the database
	rows, err := db.Query("SELECT * FROM coiffeurs")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coiffeurList []Coiffeur
	for rows.Next() {
		var coiffeur Coiffeur
		err := rows.Scan(&coiffeur.ID_coiffeur, &coiffeur.ID_salon, &coiffeur.Firstname, &coiffeur.Lastname)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		coiffeurList = append(coiffeurList, coiffeur)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coiffeurList)
}

func updateCoiffeurHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedCoiffeur Coiffeur
	err := json.NewDecoder(r.Body).Decode(&updatedCoiffeur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coiffeursMu.RLock()
	defer coiffeursMu.RUnlock()
	row := db.QueryRow("SELECT id_coiffeur FROM coiffeurs WHERE id_coiffeur=?", updatedCoiffeur.ID_coiffeur)
	if err := row.Scan(&updatedCoiffeur.ID_coiffeur); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE coiffeurs SET id_salon=?, firstname=?, lastname=? WHERE id_coiffeur=?", updatedCoiffeur.ID_salon, updatedCoiffeur.Firstname, updatedCoiffeur.Lastname, updatedCoiffeur.ID_coiffeur)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCoiffeur)
}

func deleteCoiffeurHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id_coiffeur")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coiffeursMu.Lock()
	defer coiffeursMu.Unlock()
	row := db.QueryRow("SELECT id_coiffeur FROM coiffeurs WHERE id_coiffeur=?", id)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM coiffeurs WHERE id_coiffeur=?", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CRENEAU
func addCreneauHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newCreneau Creneau
	err := json.NewDecoder(r.Body).Decode(&newCreneau)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO creneaux (id_coiffeur, date_creneau, availability) VALUES (?, ?, ?)", newCreneau.ID_coiffeur, newCreneau.Date, newCreneau.Availability)
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

	newCreneau.ID_creneau = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCreneau)
}

func getCreneauxHandler(w http.ResponseWriter, r *http.Request) {
	creneauxMu.RLock()
	defer creneauxMu.RUnlock()

	// Fetch users from the database
	rows, err := db.Query("SELECT * FROM creneaux")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var creneauList []Creneau
	for rows.Next() {
		var creneau Creneau
		err := rows.Scan(&creneau.ID_creneau, &creneau.ID_coiffeur, &creneau.Date, &creneau.Availability)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		creneauList = append(creneauList, creneau)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creneauList)
}

func updateCreneauHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedCreneau Creneau
	err := json.NewDecoder(r.Body).Decode(&updatedCreneau)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creneauxMu.RLock()
	defer creneauxMu.RUnlock()
	row := db.QueryRow("SELECT id_creneau FROM creneaux WHERE id_creneau=?", updatedCreneau.ID_creneau)
	if err := row.Scan(&updatedCreneau.ID_creneau); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE creneaux SET id_coiffeur=?, date=?, availability=? WHERE id_creneau=?", updatedCreneau.ID_coiffeur, updatedCreneau.Date, updatedCreneau.Availability, updatedCreneau.ID_creneau)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCreneau)
}

func deleteCreneauHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id_creneau")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creneauxMu.Lock()
	defer creneauxMu.Unlock()
	row := db.QueryRow("SELECT id_creneau FROM creneaux WHERE id_creneau=?", id)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM creneaux WHERE id_creneau=?", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RESERVATION
func addReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newReservation Reservation
	err := json.NewDecoder(r.Body).Decode(&newReservation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO reservations (id_salon, id_coiffeur, id_creneau) VALUES (?, ?, ?)", newReservation.ID_salon, newReservation.ID_coiffeur, newReservation.ID_creneau)
	_, err = db.Exec("UPDATE creneaux SET availability=false WHERE id_creneau=?", newReservation.ID_creneau)
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

	newReservation.ID_reservation = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newReservation)
}

func getReservationsHandler(w http.ResponseWriter, r *http.Request) {
	reservationsMu.RLock()
	defer reservationsMu.RUnlock()

	// Fetch users from the database
	rows, err := db.Query("SELECT * FROM reservations")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reservationList []Reservation
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.ID_reservation, &reservation.ID_salon, &reservation.ID_coiffeur, &reservation.ID_creneau)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reservationList = append(reservationList, reservation)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservationList)
}

func updateReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedReservation Reservation
	err := json.NewDecoder(r.Body).Decode(&updatedReservation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creneauxMu.RLock()
	defer creneauxMu.RUnlock()
	row := db.QueryRow("SELECT id_reservation FROM reservations WHERE id_reservation=?", updatedReservation.ID_reservation)
	if err := row.Scan(&updatedReservation.ID_reservation); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE reservations SET id_salon=?, id_coiffeur=?, id_creneau=? WHERE id_reservation=?", updatedReservation.ID_salon, updatedReservation.ID_coiffeur, updatedReservation.ID_creneau, updatedReservation.ID_reservation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedReservation)
}

func deleteReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id_reservation")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reservationsMu.Lock()
	defer reservationsMu.Unlock()
	row := db.QueryRow("SELECT id_reservation FROM reservations WHERE id_reservation=?", id)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM reservations WHERE id_reservation=?", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
