package main

import (
	"log"
	"os"

	"mangahub/internal"
	"mangahub/pkg/database"
	
)

func main() {
	database.Init()

	r := internal.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)
	r.Run(":" + port)
}