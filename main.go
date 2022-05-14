package main

import (
	"web-api/database"
	"web-api/server"
)

func main() {
	database.StartDB()
	defer database.CloseConn()

	server := server.NewServer()

	server.Run()
}
