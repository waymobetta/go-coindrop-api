package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/utils"
)

// AddUserID adds an AWS cognito user ID to the coindrop_auth table
func AddUserID(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new user struct object
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, "Error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		response = utils.Message(false, "JSON Unmarshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Next, insert the AWS cognito user ID into the coindrop_auth table
	if _, err := db.Client.Query(`INSERT INTO coindrop_auth (auth_user_id) VALUES ($1)`, user.Info.AuthUserID); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		response = utils.Message(false, "Could not add coindrop user to db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	utils.Respond(w, response)

	fmt.Printf("Successfully added coindrop user: %v\n\n", user.Info.AuthUserID)
}

// WalletUpdate handles updates to the wallet address for a user
func WalletUpdate(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, "Error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		response = utils.Message(false, "JSON Unmarshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update the user listing in db
	_, err = db.UpdateWallet(user)
	if err != nil {
		response = utils.Message(false, "Could not update user wallet")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	fmt.Printf("Successfully updated wallet address for user: %v\n\n", user.Info.AuthUserID)
}

// GetWalletAddress gets a user's wallet address from their auth_user_id
func GetWalletAddress(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new user struct object
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, "Error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		response = utils.Message(false, "JSON Unmarshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Get the existing entry present in the database for the given username
	row := db.Client.QueryRow(`SELECT wallet_address FROM coindrop_auth WHERE auth_user_id=$1`, user.Info.AuthUserID)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		response = utils.Message(false, "Could not get user info")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	err = row.Scan(&user.Info.WalletAddress)
	if err != nil {
		response = utils.Message(false, "Could not read rows from user info")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	utils.Respond(w, response)

	fmt.Printf("Successfully returned wallet address for user: %v\n\n", user.Info.AuthUserID)
}
