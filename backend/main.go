package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	fmt.Println("Server listening on 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
