package main

import (
	"crudappgolang/config"
	"crudappgolang/routes"
	"log"
)

func main() {
	// Initialize configuration
	config.Initialize()

	// Start the Gin server
	router := routes.SetupRouter()
	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
