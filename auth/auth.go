package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
)

// AddUserID adds an AWS cognito user ID to the coindrop_auth table
func AddUserID(w http.ResponseWriter, r *http.Request) {
	// initialize new Credentials struct object
	creds := &Credentials{}

	// Parse and decode the request body into a new `Credentials` instance
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		http.Error(w, err.Error(), 400)
		return
	}

	// Next, insert the AWS cognito user ID into the coindrop_auth table
	if _, err := db.Client.Query(`INSERT INTO coindrop_auth (auth_user_id) VALUES ($1)`, creds.AuthUserID); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&creds); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("[*] Successfully added user:", creds.AuthUserID)
}

// GetID gets a user id from their email
func GetID(w http.ResponseWriter, r *http.Request) {
	// initialize new Credentials struct object
	creds := &Credentials{}

	user := new(db.User)

	// Parse and decode the request body into a new `Credentials` instance
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		http.Error(w, err.Error(), 400)
		return
	}
	// Get the existing entry present in the database for the given username
	row := db.Client.QueryRow(`SELECT id FROM coindrop_auth WHERE auth_user_id=$1`, creds.AuthUserID)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		http.Error(w, err.Error(), 500)
		return
	}
	// Store the obtained password in `storedCreds`
	err = row.Scan(&user.Info.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user.Info.ID); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("[*] Returned user id: ", user.Info.ID)
}
