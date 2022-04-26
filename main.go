package main

import (
	 "github.com/natanista/go-api/database/seeder"
	"github.com/natanista/go-api/server"
)

func main(){
	database.StartDB()

	server := server.NewServer()

	server.Run()
}