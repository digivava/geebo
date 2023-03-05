package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// In-memory data store in the form of an array.
// Normally we would store our data in a proper database
// so that it would persist even after the app is shut down,
// but we're keeping things simple today.
var users []User

func main() {
	// Seed the in-memory store with some sample data
	users = []User{
		{ID: 1, Name: "Nikita", Email: "nikita@example.com"},
		{ID: 2, Name: "Kevin", Email: "kevin@example.com"},
		{ID: 3, Name: "Juan", Email: "juan@example.com"},
	}

	// Define the HTTP endpoints for our API
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Start the HTTP server
	log.Printf("Serving on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
