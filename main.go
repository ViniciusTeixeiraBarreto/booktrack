package main

import (
	"web-api/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
