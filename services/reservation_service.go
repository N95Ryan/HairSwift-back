package services

import (
	"encoding/json"
	"log"
	"strconv"

	"hairswift-back/models"

	supa "github.com/supabase-community/supabase-go"
)

type ReservationService struct {
	client *supa.Client
}

func NewReservationService(client *supa.Client) *ReservationService {
	return &ReservationService{client: client}
}

func (s *ReservationService) CreateReservation(reservation models.Reservation) (*models.Reservation, error) {
	data, err := json.Marshal(reservation)
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

	result, _, err := s.client.From("reservations").Insert(dataMap, false, "", "", "").Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var insertedReservation models.Reservation
	err = json.Unmarshal(result, &insertedReservation)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &insertedReservation, nil
}

func (s *ReservationService) GetAllReservations() ([]models.Reservation, error) {
	result, _, err := s.client.From("reservations").Select("*", "", false).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var reservationList []models.Reservation
	err = json.Unmarshal(result, &reservationList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return reservationList, nil
}

func (s *ReservationService) UpdateReservation(reservation models.Reservation) (*models.Reservation, error) {
	data, err := json.Marshal(reservation)
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

	result, _, err := s.client.From("reservations").Update(dataMap, "", "").Eq("id_reservation", strconv.Itoa(reservation.ID_reservation)).Execute()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var updated models.Reservation
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &updated, nil
}

func (s *ReservationService) DeleteReservation(id int) error {
	_, _, err := s.client.From("reservations").Delete("", "").Eq("id_reservation", strconv.Itoa(id)).Execute()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
