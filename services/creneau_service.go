package services

import (
	"encoding/json"
	"log"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type CreneauService struct {
	client *supa.Client
}

func NewCreneauService(client *supa.Client) *CreneauService {
	return &CreneauService{client: client}
}

func (s *CreneauService) CreateCreneau(creneau models.Creneau) (*models.Creneau, error) {
	data, err := json.Marshal(creneau)
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

	result, _, err := s.client.From("creneaux").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var insertedCreneau models.Creneau
	err = json.Unmarshal(result, &insertedCreneau)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &insertedCreneau, nil
}

func (s *CreneauService) GetAllCreneaux() ([]models.Creneau, error) {
	result, _, err := s.client.From("creneaux").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var creneauList []models.Creneau
	err = json.Unmarshal(result, &creneauList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return creneauList, nil
}

func (s *CreneauService) UpdateCreneau(creneau models.Creneau) (*models.Creneau, error) {
	data, err := json.Marshal(creneau)
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

	result, _, err := s.client.From("creneaux").Update(dataMap, "", "").Eq("id_creneau", strconv.Itoa(creneau.ID_creneau)).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var updated models.Creneau
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &updated, nil
}

func (s *CreneauService) DeleteCreneau(id int) error {
	_, _, err := s.client.From("creneaux").Delete("", "").Eq("id_creneau", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
