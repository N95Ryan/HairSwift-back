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

func AddCreneauHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newCreneau models.Creneau
	err := json.NewDecoder(r.Body).Decode(&newCreneau)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, _, _, creneauxMu, _ := database.GetMutexes()

	result, err := db.Exec("INSERT INTO creneaux (id_coiffeur, date_creneau, availability) VALUES (?, ?, ?)",
		newCreneau.ID_coiffeur, newCreneau.Date, newCreneau.Availability)
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

func GetCreneauxHandler(w http.ResponseWriter, r *http.Request) {
	_, _, _, creneauxMu, _ := database.GetMutexes()
	creneauxMu.RLock()
	defer creneauxMu.RUnlock()

	db := database.GetDB()
	rows, err := db.Query("SELECT * FROM creneaux")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var creneauList []models.Creneau
	for rows.Next() {
		var creneau models.Creneau
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

func UpdateCreneauHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedCreneau models.Creneau
	err := json.NewDecoder(r.Body).Decode(&updatedCreneau)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, _, _, creneauxMu, _ := database.GetMutexes()
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

	_, err = db.Exec("UPDATE creneaux SET id_coiffeur=?, date_creneau=?, availability=? WHERE id_creneau=?",
		updatedCreneau.ID_coiffeur, updatedCreneau.Date, updatedCreneau.Availability, updatedCreneau.ID_creneau)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCreneau)
}

func DeleteCreneauHandler(w http.ResponseWriter, r *http.Request) {
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

	db := database.GetDB()
	_, _, _, creneauxMu, _ := database.GetMutexes()
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
