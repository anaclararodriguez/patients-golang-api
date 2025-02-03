package handler

import (
	"fmt"
	"log"
	"patients-golang-api/internal/config"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailNotifier struct {
	FromAddress string
	SMTPServer  string
	SMTPPort    int
	SMTPUser    string
	SMTPPass    string
}

type Notifier interface {
	SetConfig(configs ...interface{}) error
	Send(recipient string, message string) error
}

func (e *EmailNotifier) SetConfig(configs ...interface{}) error {
	conf := config.GetConfig()
	e.FromAddress = conf.FromAddress
	e.SMTPServer = conf.SmtpServer

	var err error
	e.SMTPPort, err = strconv.Atoi(conf.SmtpPort)
	if err != nil {
		log.Fatalf("Error converting SMTP Port: %v", err)
	}

	e.SMTPUser = conf.SmtpUser

	e.SMTPPass = conf.SmtpPass

	return nil
}

func (e *EmailNotifier) Send(recipient string, name string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", e.FromAddress)
	mail.SetHeader("To", recipient)
	mail.SetHeader("Subject", "Patient Registration Confirmation")
	mail.SetBody("text/plain", fmt.Sprintf("Hello %s,\n\nYour registration was successful!", name))

	dialer := gomail.NewDialer(e.SMTPServer, e.SMTPPort, e.SMTPUser, e.SMTPPass)
	if err := dialer.DialAndSend(mail); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}
