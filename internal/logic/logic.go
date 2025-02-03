package logic

import (
	"fmt"
	"patients-golang-api/internal/model"
	"regexp"
	"strings"
)

func ValidatePatient(p model.Patient) error {
	p.Name = strings.TrimSpace(p.Name)
	p.Email = strings.TrimSpace(p.Email)
	p.Address = strings.TrimSpace(p.Address)
	p.ContactNumber = strings.TrimSpace(p.ContactNumber)
	p.DocumentPhoto = strings.TrimSpace(p.DocumentPhoto)

	if p.Name == "" || p.Email == "" || p.Address == "" || p.ContactNumber == "" || p.DocumentPhoto == "" {
		return fmt.Errorf("All fields (Name, Email, Contact number, Document photo) are required")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(p.Email) {
		return fmt.Errorf("Invalid email format")
	}

	phoneRegex := `^\d{7,15}$`
	if !regexp.MustCompile(phoneRegex).MatchString(p.ContactNumber) {
		return fmt.Errorf("Invalid phone number")
	}

	return nil
}
