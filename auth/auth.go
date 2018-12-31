package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
	"golang.org/x/crypto/bcrypt"
)

// SignUp registers a user for coindrop
func SignUp(w http.ResponseWriter, r *http.Request) {
	// initialize new Credentials struct object
	creds := &Credentials{}

	// Parse and decode the request body into a new `Credentials` instance
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		http.Error(w, err.Error(), 400)
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		// If for some reason there is an error generating the password hash from the passed in password throw 500 error
		http.Error(w, err.Error(), 500)
		return
	}

	// Next, insert the username, along with the hashed password into the database
	if _, err := db.Client.Query(`INSERT INTO coindrop_auth (email,password) VALUES ($1,$2)`, creds.Email, string(hashedPassword)); err != nil {
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

	fmt.Println("[*] Successful sign up for user:", creds.Email)
}

// SignIn logs in a user for coindrop
func SignIn(w http.ResponseWriter, r *http.Request) {
	// initialize new Credentials struct object
	creds := &Credentials{}

	// Parse and decode the request body into a new `Credentials` instance
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		http.Error(w, err.Error(), 400)
		return
	}
	// Get the existing entry present in the database for the given username
	row := db.Client.QueryRow(`SELECT password FROM coindrop_auth WHERE email=$1`, creds.Email)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		http.Error(w, err.Error(), 500)
		return
	}
	// Create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err = row.Scan(&storedCreds.Password)
	if err != nil {
		if storedCreds.Email != creds.Email {
			// If an entry with the username does not exist, send an "Unauthorized"(401) status
			http.Error(w, err.Error(), 401)
			return
		}
		http.Error(w, err.Error(), 500)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		http.Error(w, err.Error(), 401)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(creds); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("[*] Successful sign in for user:", creds.Email)
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
	row := db.Client.QueryRow(`SELECT id FROM coindrop_auth WHERE email=$1`, creds.Email)
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
