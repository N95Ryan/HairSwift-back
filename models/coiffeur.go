package models

type Coiffeur struct {
	ID_coiffeur int    `json:"id_coiffeur"`
	ID_salon    int    `json:"id_salon"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
}
