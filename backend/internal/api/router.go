package api

import "net/http"

func NewRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", HealthHandler)

	return mux
}
