package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
)

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
	// Get the existing entry present in the database for the given AWS cognito user_ID
	row := db.Client.QueryRow(`SELECT id FROM coindrop_auth WHERE user_id=$1`, creds.UserID)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		http.Error(w, err.Error(), 500)
		return
	}
	// Store the obtained id in `storedCreds`
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
