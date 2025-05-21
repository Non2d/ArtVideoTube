package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Non2d/ArtVideoTube/backend/internal/presentation/http/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Setup router
	r := router.New()

	log.Printf("Starting server on :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
