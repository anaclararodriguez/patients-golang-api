package main

import (
	"fmt"
	"patients-golang-api/internal/db"
	"patients-golang-api/internal/handler"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error when trying to connect to DB: ", err.Error())
	}
	server := handler.NewServer("8080")
	server.Start()
}
