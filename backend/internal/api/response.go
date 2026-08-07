package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Sends a standardized JSON response with the given HTTP status code.
func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Erreur d'encodage JSON de la réponse : %v", err)
	}
}
