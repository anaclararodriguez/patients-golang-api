package handler

import (
	"encoding/json"
	"net/http"
	"patients-golang-api/internal/db"
	"patients-golang-api/internal/logic"
	"patients-golang-api/internal/model"
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

	if err := logic.ValidatePatient(patient); err != nil {
		sendResponse(w, StatusError, err.Error(), nil, nil)
		return
	}

	err = db.InsertPatient(patient)
	if err != nil {
		sendResponse(w, StatusError, "Failed to create patient", nil, err)
		return
	}
	go logic.SendConfirmationEmail(patient.Email, patient.Name)
	sendResponse(w, StatusSuccess, "Patient created successfully", patient, nil)
}
