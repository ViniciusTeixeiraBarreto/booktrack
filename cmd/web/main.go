package main

import (
	"booktrack/pkg/api"
	"booktrack/pkg/database"
	"context"
)

func main() {
	ctx := context.Background()

	db := database.StartDB()
	defer database.CloseConn()

	ctx = context.WithValue(ctx, "database", db)

	server := api.NewServer()

	server.RunFiber()
}
