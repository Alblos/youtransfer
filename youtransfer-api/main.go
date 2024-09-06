package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	apiv1 := r.PathPrefix("/api/v1").Subrouter()

	apiv1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}).Methods(http.MethodGet)

	printRoutes(r) // Print the routes

	c := cors.Default().Handler(r) // CORS

	// Start the server and print server started
	err := http.ListenAndServe(":8080", c)
	if err != nil {
		log.Panicf("Error starting server: %v", err)
	} else {
		println("Server started")
	}
}

func printRoutes(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		m, _ := route.GetMethods()
		println(t, m)
		return nil
	})
	if err != nil {
		log.Panicf("Error walking routes: %v", err)
		return
	}
}
