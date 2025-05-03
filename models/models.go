package models

type Client struct {
	ID_client int    `json:"id_client"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Salon struct {
	ID_salon   int    `json:"id_salon"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Phone      string `json:"phone"`
}

type Coiffeur struct {
	ID_coiffeur int    `json:"id_coiffeur"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ID_salon    int    `json:"id_salon"`
}

type Creneau struct {
	ID_creneau  int    `json:"id_creneau"`
	Date        string `json:"date"`
	Heure_debut string `json:"heure_debut"`
	Heure_fin   string `json:"heure_fin"`
	ID_coiffeur int    `json:"id_coiffeur"`
}

type Reservation struct {
	ID_reservation int    `json:"id_reservation"`
	Date           string `json:"date"`
	ID_client      int    `json:"id_client"`
	ID_creneau     int    `json:"id_creneau"`
}
