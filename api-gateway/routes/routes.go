package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-gateway/handlers"
)

func HandleRoutes(r *mux.Router, endpoint *map[string]interface{}) {
	r.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	r.HandleFunc("/api/listBusinesses", handlers.ListBusinesses).Methods("GET")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
