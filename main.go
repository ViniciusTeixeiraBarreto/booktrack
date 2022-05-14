package main

import (
	"web-api/database"
	"web-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
