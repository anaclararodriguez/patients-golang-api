package main

import (
	"fmt"
	"log"
	"patients-golang-api/internal/config"
	"patients-golang-api/internal/db"
	"patients-golang-api/internal/handler"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	err = db.ConnectDB(*config)
	if err != nil {
		fmt.Println("Error when trying to connect to DB: ", err.Error())
	}
	server := handler.NewServer("8080")
	server.Start()
}
