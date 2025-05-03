package models

type Reservation struct {
	ID_reservation int `json:"id_reservation"`
	ID_salon       int `json:"id_salon"`
	ID_coiffeur    int `json:"id_coiffeur"`
	ID_creneau     int `json:"id_creneau"`
}
