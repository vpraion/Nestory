package api

import "net/http"

// HealthResponse represents the payload returned by the healthcheck endpoint.
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler handles HTTP GET requests to inspect server health status.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, HealthResponse{Status: "ok"})
}
