package services

import (
	"encoding/json"
	"log"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type SalonService struct {
	client *supa.Client
}

func NewSalonService(client *supa.Client) *SalonService {
	return &SalonService{client: client}
}

func (s *SalonService) CreateSalon(salon models.Salon) (*models.Salon, error) {
	data, err := json.Marshal(salon)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, _, err := s.client.From("salons").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var insertedSalon models.Salon
	err = json.Unmarshal(result, &insertedSalon)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &insertedSalon, nil
}

func (s *SalonService) GetAllSalons() ([]models.Salon, error) {
	result, _, err := s.client.From("salons").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var salonList []models.Salon
	err = json.Unmarshal(result, &salonList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return salonList, nil
}

func (s *SalonService) UpdateSalon(salon models.Salon) (*models.Salon, error) {
	data, err := json.Marshal(salon)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, _, err := s.client.From("salons").Update(dataMap, "", "").Eq("id_salon", strconv.Itoa(salon.ID_salon)).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var updated models.Salon
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &updated, nil
}

func (s *SalonService) DeleteSalon(id int) error {
	_, _, err := s.client.From("salons").Delete("", "").Eq("id_salon", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
