package database

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db             *sql.DB
	clientsMu      sync.RWMutex
	salonsMu       sync.RWMutex
	coiffeursMu    sync.RWMutex
	creneauxMu     sync.RWMutex
	reservationsMu sync.RWMutex
)

func InitDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", "goteam:root@tcp(localhost:3306)/golang")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			id_client INT AUTO_INCREMENT PRIMARY KEY,
			firstname VARCHAR(150),
			lastname VARCHAR(150),
			email VARCHAR(150),
			password VARCHAR(255)
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS salons (
			id_salon INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(150)
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS coiffeurs (
			id_coiffeur INT AUTO_INCREMENT PRIMARY KEY,
			id_salon INT,
			firstname VARCHAR(150),
			lastname VARCHAR(150)
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS creneaux (
			id_creneau INT AUTO_INCREMENT PRIMARY KEY,
			id_coiffeur INT,
			date_creneau VARCHAR(150),
			availability BOOLEAN
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS reservations (
			id_reservation INT AUTO_INCREMENT PRIMARY KEY,
			id_salon INT,
			id_coiffeur INT,
			id_creneau INT
		);
	`)
	return err
}

func GetDB() *sql.DB {
	return db
}

func GetMutexes() (*sync.RWMutex, *sync.RWMutex, *sync.RWMutex, *sync.RWMutex, *sync.RWMutex) {
	return &clientsMu, &salonsMu, &coiffeursMu, &creneauxMu, &reservationsMu
}
