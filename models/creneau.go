package models

type Creneau struct {
	ID_creneau   int    `json:"id_creneau"`
	ID_coiffeur  int    `json:"id_coiffeur"`
	Date         string `json:"date_creneau"`
	Availability bool   `json:"availability"`
}
