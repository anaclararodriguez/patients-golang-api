package main

import (
	"patients-golang-api/internal/db"
	"patients-golang-api/internal/handler"
)

func main() {
	db.ConnectDB()
	server := handler.NewServer("8080")
	server.Start()
}
