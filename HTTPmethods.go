package main

import (
	"encoding/json"
	"net/http"
)

func userAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := DBloginUser(r.Header.Get("name"), r.Header.Get("pass"))

	if user != nil {
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode("No user found")
	}
}

func userRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name, password := r.Header.Get("name"), r.Header.Get("pass")

	user := DBfindUser(name)

	if user != nil {
		json.NewEncoder(w).Encode("User exists!")
	} else {
		json.NewEncoder(w).Encode(DBregisterUser(name, password))
	}
}
