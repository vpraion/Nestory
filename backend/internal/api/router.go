package api

import "net/http"

// NewRouter initializes and configures the HTTP request multiplexer with API routes.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", healthHandler)

	return mux
}
