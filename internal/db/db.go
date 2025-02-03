package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"patients-golang-api/internal/model"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(config model.Config) error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDb)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		return fmt.Errorf("error verifying database connection: %w", err)
	}

	fmt.Println("Connected to MySQL successfully")

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

		<-signalChan
		log.Println("Received shutdown signal, closing database connection...")

		if err := DB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}

		os.Exit(0)
	}()
	return nil
}

func GetPatients() ([]model.Patient, error) {
	rows, err := DB.Query("SELECT id, name, email, address, contact_number, document_photo FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []model.Patient
	for rows.Next() {
		var patient model.Patient

		if err := rows.Scan(&patient.ID, &patient.Name, &patient.Email, &patient.Address, &patient.ContactNumber, &patient.DocumentPhoto); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return patients, nil
}

func InsertPatient(patient model.Patient) error {
	query := "INSERT INTO patients (name, email, address, contact_number, document_photo) VALUES (?, ?, ?, ?, ?)"
	_, err := DB.Exec(query, patient.Name, patient.Email, patient.Address, patient.ContactNumber, patient.DocumentPhoto)
	return err
}
