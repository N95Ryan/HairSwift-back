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

func AddCoiffeurHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newCoiffeur models.Coiffeur
	err := json.NewDecoder(r.Body).Decode(&newCoiffeur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, _, coiffeursMu, _, _ := database.GetMutexes()

	result, err := db.Exec("INSERT INTO coiffeurs (id_salon, firstname, lastname) VALUES (?, ?, ?)",
		newCoiffeur.ID_salon, newCoiffeur.Firstname, newCoiffeur.Lastname)
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

func GetCoiffeursHandler(w http.ResponseWriter, r *http.Request) {
	_, _, coiffeursMu, _, _ := database.GetMutexes()
	coiffeursMu.RLock()
	defer coiffeursMu.RUnlock()

	db := database.GetDB()
	rows, err := db.Query("SELECT * FROM coiffeurs")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coiffeurList []models.Coiffeur
	for rows.Next() {
		var coiffeur models.Coiffeur
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

func UpdateCoiffeurHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedCoiffeur models.Coiffeur
	err := json.NewDecoder(r.Body).Decode(&updatedCoiffeur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, _, coiffeursMu, _, _ := database.GetMutexes()
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

	_, err = db.Exec("UPDATE coiffeurs SET id_salon=?, firstname=?, lastname=? WHERE id_coiffeur=?",
		updatedCoiffeur.ID_salon, updatedCoiffeur.Firstname, updatedCoiffeur.Lastname, updatedCoiffeur.ID_coiffeur)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCoiffeur)
}

func DeleteCoiffeurHandler(w http.ResponseWriter, r *http.Request) {
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

	db := database.GetDB()
	_, _, coiffeursMu, _, _ := database.GetMutexes()
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
