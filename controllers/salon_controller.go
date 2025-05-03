package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type SalonController struct {
	client *supa.Client
}

func NewSalonController(client *supa.Client) *SalonController {
	return &SalonController{client: client}
}

func (c *SalonController) AddSalon(w http.ResponseWriter, r *http.Request) {
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

	data, err := json.Marshal(newSalon)
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

	result, _, err := c.client.From("salons").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var insertedSalon models.Salon
	err = json.Unmarshal(result, &insertedSalon)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedSalon)
}

func (c *SalonController) GetSalons(w http.ResponseWriter, r *http.Request) {
	result, _, err := c.client.From("salons").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var salonList []models.Salon
	err = json.Unmarshal(result, &salonList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salonList)
}

func (c *SalonController) UpdateSalon(w http.ResponseWriter, r *http.Request) {
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

	data, err := json.Marshal(updatedSalon)
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

	result, _, err := c.client.From("salons").Update(dataMap, "", "").Eq("id_salon", strconv.Itoa(updatedSalon.ID_salon)).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updated models.Salon
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (c *SalonController) DeleteSalon(w http.ResponseWriter, r *http.Request) {
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

	_, _, err = c.client.From("salons").Delete("", "").Eq("id_salon", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
