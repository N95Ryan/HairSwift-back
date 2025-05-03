package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type CreneauController struct {
	client *supa.Client
}

func NewCreneauController(client *supa.Client) *CreneauController {
	return &CreneauController{client: client}
}

func (c *CreneauController) AddCreneau(w http.ResponseWriter, r *http.Request) {
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

	data, err := json.Marshal(newCreneau)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, _, err := c.client.From("creneaux").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var insertedCreneau models.Creneau
	err = json.Unmarshal(result, &insertedCreneau)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedCreneau)
}

func (c *CreneauController) GetCreneaux(w http.ResponseWriter, r *http.Request) {
	result, _, err := c.client.From("creneaux").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var creneauList []models.Creneau
	err = json.Unmarshal(result, &creneauList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creneauList)
}

func (c *CreneauController) UpdateCreneau(w http.ResponseWriter, r *http.Request) {
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

	data, err := json.Marshal(updatedCreneau)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, _, err := c.client.From("creneaux").Update(dataMap, "", "").Eq("id_creneau", strconv.Itoa(updatedCreneau.ID_creneau)).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updated models.Creneau
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (c *CreneauController) DeleteCreneau(w http.ResponseWriter, r *http.Request) {
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

	_, _, err = c.client.From("creneaux").Delete("", "").Eq("id_creneau", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
