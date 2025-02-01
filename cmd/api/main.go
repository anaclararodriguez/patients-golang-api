package main

import (
	"patients-golang-api/internal/handler"
)

func main() {

	server := handler.NewServer("8080")
	server.Start()
}
