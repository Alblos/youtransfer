package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"youtransfer/api/db"
	"youtransfer/api/db/models"
	"youtransfer/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// DB
	db.Connect()
	migrateModels()

	// Router setup
	r := mux.NewRouter()
	apiv1 := r.PathPrefix("/api/v1").Subrouter()

	routes.AddUserRoutes(apiv1.PathPrefix("/users").Subrouter())

	apiv1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}).Methods(http.MethodGet)

	c := cors.Default().Handler(r) // CORS

	// Start the server and print server started
	err = http.ListenAndServe(":8080", c)
	if err != nil {
		log.Panicf("Error starting server: %v", err)
	} else {
		println("Server started")
	}
}

func migrateModels() {
	err := db.GetConnection().AutoMigrate(&models.ApiKey{})

	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	} else {
		log.Println("Models migrated successfully")
	}
}
