package main

// go run cmd/api-server/main.go
import (
	"mangahub/internal"
	"mangahub/pkg/database"
)

func main() {
	database.Connect()

	r := internal.SetupRouter(database.DB)
	r.Run(":8080")
}

