package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"patients-golang-api/internal/db"
	"patients-golang-api/internal/model"
	"regexp"
	"strings"
)

func GetPatients(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetPatients()
	if err != nil {
		sendResponse(w, StatusError, "Error fetching users", nil, err)
		return
	}
	sendResponse(w, StatusSuccess, "Users retrieved susccessfully", users, nil)

}

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient model.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		sendResponse(w, StatusError, "Invalid input format", nil, err)
		return
	}

	if err := validatePatient(patient); err != nil {
		sendResponse(w, StatusError, err.Error(), nil, nil)
		return
	}

	err = db.InsertPatient(patient)
	if err != nil {
		sendResponse(w, StatusError, "Failed to create patient", nil, err)
		return
	}

	sendResponse(w, StatusSuccess, "Patient created successfully", patient, nil)
}

func validatePatient(p model.Patient) error {
	p.Name = strings.TrimSpace(p.Name)
	p.Email = strings.TrimSpace(p.Email)
	p.PhoneNumber = strings.TrimSpace(p.PhoneNumber)
	p.DocumentPhoto = strings.TrimSpace(p.DocumentPhoto)

	if p.Name == "" || p.Email == "" || p.PhoneNumber == "" {
		return fmt.Errorf("All fields (FirstName, LastName, Email, Document photo) are required")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(p.Email) {
		return fmt.Errorf("Invalid email format")
	}

	phoneRegex := `^\d{7,15}$`
	if !regexp.MustCompile(phoneRegex).MatchString(p.PhoneNumber) {
		return fmt.Errorf("Invalid phone number")
	}

	// TODO: validate photo

	return nil
}
