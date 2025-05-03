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

func AddSalonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newSalon models.Salon
	err := json.NewDecoder(r.Body).Decode(&newSalon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, salonsMu, _, _, _ := database.GetMutexes()

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

func GetSalonsHandler(w http.ResponseWriter, r *http.Request) {
	_, salonsMu, _, _, _ := database.GetMutexes()
	salonsMu.RLock()
	defer salonsMu.RUnlock()

	db := database.GetDB()
	rows, err := db.Query("SELECT * FROM salons")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var salonList []models.Salon
	for rows.Next() {
		var salon models.Salon
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

func UpdateSalonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedSalon models.Salon
	err := json.NewDecoder(r.Body).Decode(&updatedSalon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	_, salonsMu, _, _, _ := database.GetMutexes()
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

func DeleteSalonHandler(w http.ResponseWriter, r *http.Request) {
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

	db := database.GetDB()
	_, salonsMu, _, _, _ := database.GetMutexes()
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
