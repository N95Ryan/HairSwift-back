package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type ClientService struct {
	client *supa.Client
}

func NewClientService(client *supa.Client) *ClientService {
	return &ClientService{client: client}
}

func (s *ClientService) CreateClient(client models.Client) (*models.Client, error) {
	log.Printf("Tentative de création d'un client: %+v", client)

	data, err := json.Marshal(client)
	if err != nil {
		log.Printf("Erreur lors du marshalling du client: %v", err)
		return nil, fmt.Errorf("erreur lors du marshalling du client: %v", err)
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Printf("Erreur lors du unmarshalling en map: %v", err)
		return nil, fmt.Errorf("erreur lors du unmarshalling en map: %v", err)
	}

	log.Printf("Données à insérer: %+v", dataMap)

	result, _, err := s.client.From("clients").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Printf("Erreur lors de l'insertion dans Supabase: %v", err)
		return nil, fmt.Errorf("erreur lors de l'insertion dans Supabase: %v", err)
	}

	log.Printf("Résultat de l'insertion: %s", string(result))

	var insertedClient models.Client
	err = json.Unmarshal(result, &insertedClient)
	if err != nil {
		log.Printf("Erreur lors du unmarshalling du résultat: %v", err)
		return nil, fmt.Errorf("erreur lors du unmarshalling du résultat: %v", err)
	}

	log.Printf("Client créé avec succès: %+v", insertedClient)
	return &insertedClient, nil
}

func (s *ClientService) GetAllClients() ([]models.Client, error) {
	result, _, err := s.client.From("clients").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var clientList []models.Client
	err = json.Unmarshal(result, &clientList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return clientList, nil
}

func (s *ClientService) UpdateClient(client models.Client) (*models.Client, error) {
	data, err := json.Marshal(client)
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

	result, _, err := s.client.From("clients").Update(dataMap, "", "").Eq("id_client", strconv.Itoa(client.ID_client)).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var updated models.Client
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &updated, nil
}

func (s *ClientService) DeleteClient(id int) error {
	_, _, err := s.client.From("clients").Delete("", "").Eq("id_client", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
