package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received getUsers request.")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received getUser request.")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range users {
		if strconv.Itoa(user.ID) == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received createUser request.")

	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	if user.Username == "" {
		log.Printf("ERROR: Missing Username")
		json.NewEncoder(w).Encode("ERROR: Missing Username")
		return
	}
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received updateUser request.")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if strconv.Itoa(user.ID) == params["id"] {
			var updatedUser User
			json.NewDecoder(r.Body).Decode(&updatedUser)
			updatedUser.ID = user.ID
			users[index] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received deleteUser request.")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if strconv.Itoa(user.ID) == params["id"] {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
