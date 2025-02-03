package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"patients-golang-api/internal/model"
)

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

func sendResponse(w http.ResponseWriter, status string, message string, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, fmt.Sprintf("%s: %v", message, err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(model.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
