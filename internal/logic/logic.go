package logic

import (
	"fmt"
	"log"
	"patients-golang-api/internal/model"
	"regexp"
	"strings"

	"gopkg.in/gomail.v2"
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

func SendConfirmationEmail(email, name string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "justanemail726@gmail.com")
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Patient Registration Confirmation")
	mail.SetBody("text/plain", fmt.Sprintf("Hello %s,\n\nYour registration was successful!", name))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "justanemail726@gmail.com", "vfpp fmpw vehi zamd")

	if err := dialer.DialAndSend(mail); err != nil {
		log.Println("Failed to send email:", err)
	}
}
