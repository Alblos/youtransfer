package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"youtransfer/api/db/models"
)

func AddUserRoutes(r *mux.Router) {
	// r.HandleFunc("", createUser).Methods(http.MethodPost)
	r.HandleFunc("/api-key", createApiKey).Methods(http.MethodPost)
}

func createApiKey(w http.ResponseWriter, r *http.Request) {
	key := models.GenerateApiKey()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]string{"key": key})
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
