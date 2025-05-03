package services

import (
	"encoding/json"
	"log"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type CoiffeurService struct {
	client *supa.Client
}

func NewCoiffeurService(client *supa.Client) *CoiffeurService {
	return &CoiffeurService{client: client}
}

func (s *CoiffeurService) CreateCoiffeur(coiffeur models.Coiffeur) (*models.Coiffeur, error) {
	data, err := json.Marshal(coiffeur)
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

	result, _, err := s.client.From("coiffeurs").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var insertedCoiffeur models.Coiffeur
	err = json.Unmarshal(result, &insertedCoiffeur)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &insertedCoiffeur, nil
}

func (s *CoiffeurService) GetAllCoiffeurs() ([]models.Coiffeur, error) {
	result, _, err := s.client.From("coiffeurs").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var coiffeurList []models.Coiffeur
	err = json.Unmarshal(result, &coiffeurList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return coiffeurList, nil
}

func (s *CoiffeurService) UpdateCoiffeur(coiffeur models.Coiffeur) (*models.Coiffeur, error) {
	data, err := json.Marshal(coiffeur)
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

	result, _, err := s.client.From("coiffeurs").Update(dataMap, "", "").Eq("id_coiffeur", strconv.Itoa(coiffeur.ID_coiffeur)).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var updated models.Coiffeur
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &updated, nil
}

func (s *CoiffeurService) DeleteCoiffeur(id int) error {
	_, _, err := s.client.From("coiffeurs").Delete("", "").Eq("id_coiffeur", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
