package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/utils"
)

// UserIDAdd adds an AWS cognito user ID to the coindrop_auth table
func (h *Handlers) UserIDAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new user struct object
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("[db] %v\n", err)
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Next, insert the AWS cognito user ID into the coindrop_auth table
	_, err = h.db.AddUserID(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	utils.Respond(w, response)

	fmt.Printf("[db] successfully added coindrop user: %v\n", user.AuthUserID)
}

// WalletUpdate handles updates to the wallet address for a user
func (h *Handlers) WalletUpdate(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update the user listing in db
	_, err = h.db.UpdateWallet(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	log.Printf("[db] successfully updated wallet address for user: %v\n", user.AuthUserID)
}

// WalletGet gets a user's wallet address from their auth_user_id
func (h *Handlers) WalletGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new user struct object
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("[db] %v\n", err)
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Get the existing entry present in the database for the given username
	user, err = h.db.GetWallet(user)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, user.WalletAddress)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	utils.Respond(w, response)

	log.Printf("[db] successfully returned wallet address for user: %v\n", user.AuthUserID)
}
