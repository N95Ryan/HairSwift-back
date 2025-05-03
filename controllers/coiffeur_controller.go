package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type CoiffeurController struct {
	client *supa.Client
}

func NewCoiffeurController(client *supa.Client) *CoiffeurController {
	return &CoiffeurController{client: client}
}

func (c *CoiffeurController) AddCoiffeur(w http.ResponseWriter, r *http.Request) {
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

	data, err := json.Marshal(newCoiffeur)
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

	result, _, err := c.client.From("coiffeurs").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var insertedCoiffeur models.Coiffeur
	err = json.Unmarshal(result, &insertedCoiffeur)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedCoiffeur)
}

func (c *CoiffeurController) GetCoiffeurs(w http.ResponseWriter, r *http.Request) {
	result, _, err := c.client.From("coiffeurs").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var coiffeurList []models.Coiffeur
	err = json.Unmarshal(result, &coiffeurList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coiffeurList)
}

func (c *CoiffeurController) UpdateCoiffeur(w http.ResponseWriter, r *http.Request) {
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

	data, err := json.Marshal(updatedCoiffeur)
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

	result, _, err := c.client.From("coiffeurs").Update(dataMap, "", "").Eq("id_coiffeur", strconv.Itoa(updatedCoiffeur.ID_coiffeur)).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updated models.Coiffeur
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (c *CoiffeurController) DeleteCoiffeur(w http.ResponseWriter, r *http.Request) {
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

	_, _, err = c.client.From("coiffeurs").Delete("", "").Eq("id_coiffeur", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
