package db

import (
	"database/sql"
	"fmt"
	"patients-golang-api/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {
	var err error
	dsn := ""
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error verifying database connection: %w", err)
	}

	fmt.Println("Connected to MySQL successfully")
	return nil
}

func GetPatients() ([]model.Patient, error) {
	return nil, nil
}

func InsertPatient(patient model.Patient) error {
	return nil
}
