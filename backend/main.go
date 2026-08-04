package main

import (
	"Nestory/internal/api"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	router := api.NewRouter()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
