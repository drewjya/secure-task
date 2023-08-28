package main

import (
	"log"

	app "github.com/drewjya/secure-task"
	postgresql "github.com/drewjya/secure-task/libs/postgresql"
	"github.com/drewjya/secure-task/presentation/route"
	"github.com/drewjya/secure-task/repository"
)

func main() {
	app := app.SetupApp()

	// Get the desired port from environment variable or use a default value
	port := postgresql.GoDotEnvVariable("PORT")
	if port == "" {
		port = "3001"
	}
	db, err := postgresql.Connect()
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
	defer db.Close()

	queries := repository.New(db)

	route.Setup(app, queries)

	// Start the server
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
