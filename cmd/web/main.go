package main

import (
	"context"
	"web-api/pkg/api"
	"web-api/pkg/database"
)

func main() {
	ctx := context.Background()

	db := database.StartDB()
	defer database.CloseConn()

	ctx = context.WithValue(ctx, "database", db)

	server := api.NewServer()

	server.Run()
}
