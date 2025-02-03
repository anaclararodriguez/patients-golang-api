package main

import (
	"fmt"
	"patients-golang-api/internal/config"
	"patients-golang-api/internal/db"
	"patients-golang-api/internal/handler"
)

func main() {
	conf := config.GetConfig()
	err := db.ConnectDB(*conf)
	if err != nil {
		fmt.Println("Error when trying to connect to DB: ", err.Error())
	}
	server := handler.NewServer(conf.ServerPort)
	server.Start()
}
