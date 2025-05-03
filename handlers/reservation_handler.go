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

func AddReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newReservation models.Reservation
	err := json.NewDecoder(r.Body).Decode(&newReservation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, _, _, _, reservationsMu := database.GetMutexes()

	// Vérifier si le créneau est disponible
	var availability bool
	err = db.QueryRow("SELECT availability FROM creneaux WHERE id_creneau=?", newReservation.ID_creneau).Scan(&availability)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !availability {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// Créer la réservation
	result, err := db.Exec("INSERT INTO reservations (id_salon, id_coiffeur, id_creneau) VALUES (?, ?, ?)",
		newReservation.ID_salon, newReservation.ID_coiffeur, newReservation.ID_creneau)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Mettre à jour la disponibilité du créneau
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

func GetReservationsHandler(w http.ResponseWriter, r *http.Request) {
	_, _, _, _, reservationsMu := database.GetMutexes()
	reservationsMu.RLock()
	defer reservationsMu.RUnlock()

	db := database.GetDB()
	rows, err := db.Query("SELECT * FROM reservations")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reservationList []models.Reservation
	for rows.Next() {
		var reservation models.Reservation
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

func UpdateReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedReservation models.Reservation
	err := json.NewDecoder(r.Body).Decode(&updatedReservation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, _, _, _, reservationsMu := database.GetMutexes()
	reservationsMu.RLock()
	defer reservationsMu.RUnlock()

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

	_, err = db.Exec("UPDATE reservations SET id_salon=?, id_coiffeur=?, id_creneau=? WHERE id_reservation=?",
		updatedReservation.ID_salon, updatedReservation.ID_coiffeur, updatedReservation.ID_creneau, updatedReservation.ID_reservation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedReservation)
}

func DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {
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

	db := database.GetDB()
	_, _, _, _, reservationsMu := database.GetMutexes()
	reservationsMu.Lock()
	defer reservationsMu.Unlock()

	// Récupérer l'ID du créneau avant de supprimer la réservation
	var creneauID int
	err = db.QueryRow("SELECT id_creneau FROM reservations WHERE id_reservation=?", id).Scan(&creneauID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Supprimer la réservation
	_, err = db.Exec("DELETE FROM reservations WHERE id_reservation=?", id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remettre le créneau à disponible
	_, err = db.Exec("UPDATE creneaux SET availability=true WHERE id_creneau=?", creneauID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
