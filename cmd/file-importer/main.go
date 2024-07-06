package main

import (
	"log"
	"os"

	"github.com/Reaper1994/go-file-storage-service/db"
	"github.com/Reaper1994/go-file-storage-service/pkg/server"
)

func main() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	db.InitDB(uri)

	s := server.NewServer()
	s.ConfigureRouter()

	port := "8080"
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(s.Start(port))
}
