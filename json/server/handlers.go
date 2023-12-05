package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User added successfully")
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}
