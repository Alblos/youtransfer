package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	apiv1 := r.PathPrefix("/api/v1").Subrouter()

	apiv1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}).Methods("GET")

	// Print the routes
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

	// Start the server and print server started
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		println("Error starting server")
	} else {
		println("Server started")
	}
}
