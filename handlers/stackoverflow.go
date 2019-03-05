// +build ignore

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/stackoverflow"
	"github.com/waymobetta/go-coindrop-api/utils"
	"github.com/waymobetta/go-coindrop-api/verify"
)

// StackUserAdd adds a single user listing to db
func (h *Handlers) StackUserAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
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

	// add user listing to db
	_, err = h.db.AddStackUser(user)
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

	log.Printf("[handler] successfully added stack user: %v\n", user.AuthUserID)
}

// StackUserGet returns information about a single user
func (h *Handlers) StackUserGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
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

	// get user listing by name
	userData, err := h.db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, userData)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	log.Printf("[handler] successfully returned information for user: %v\n", user.AuthUserID)
}

// GenerateStackVerificationCode creates a verifcation code for Stack Overflow
func (h *Handlers) GenerateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// declare new variable user of User struct
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

	// generate temporary verification code
	verificationCode := verify.GenerateVerificationCode()

	// promotional display code
	// TODO: store in db and read from db
	displayCode := fmt.Sprintf("[COINDROP.IO - IT PAYS TO CONTRIBUTE: %s]", verificationCode)

	// update local user object variable with generated verification code
	user.StackOverflowData.VerificationData.StoredVerificationCode = displayCode

	// marshal into JSON
	_, err = json.Marshal(&user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// store user verification data in db
	_, err = h.db.UpdateStackVerificationCode(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	log.Printf("[handler] successfully generated verification code for user: %v\n", user.AuthUserID)
}

// ValidateStackVerificationCode validates the temporary verification code
func (h *Handlers) ValidateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// declare new variable user of User struct
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

	// pull stored verification code from DB
	storedStackUser, err := h.db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// check Stack Overflow for matching verification code
	updatedStackUser, err := stackoverflow.GetProfileByUserID(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Printf("[handler] checking %s against %s\n", storedStackUser.StackOverflowData.VerificationData.PostedVerificationCode, storedStackUser.StackOverflowData.VerificationData.StoredVerificationCode)

	// secondary validation to see if codes match
	if !strings.Contains(updatedStackUser.StackOverflowData.VerificationData.PostedVerificationCode, storedStackUser.StackOverflowData.VerificationData.StoredVerificationCode) {
		log.Warnf("[handler] Verification codes do not match!")

		response = utils.Message(false, "unsuccessful")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update verification field values
	storedStackUser.StackOverflowData.VerificationData.PostedVerificationCode = updatedStackUser.StackOverflowData.VerificationData.PostedVerificationCode
	storedStackUser.StackOverflowData.VerificationData.IsVerified = true

	// update db with new info since verification codes matched
	_, err = h.db.UpdateStackVerificationCode(storedStackUser)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	log.Printf("[handler] successfully validated verification code for user: %v\n", user.AuthUserID)
}

// StackUserUpdate updates and returns profile info about the user
func (h *Handlers) StackUserUpdate(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// declare new variable of type User
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

	// pull stored verification code from DB
	storedUserInfo, err := h.db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[handler] retrieving Stack Overflow About info")
	// get general about info for user
	_, err = stackoverflow.GetProfileByUserID(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[handler] retrieving Stack Overflow associated accounts info")
	// get list of trophies user has been awarded
	_, err = stackoverflow.GetAssociatedAccounts(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	user.StackOverflowData.VerificationData.PostedVerificationCode = storedUserInfo.StackOverflowData.VerificationData.PostedVerificationCode
	user.StackOverflowData.VerificationData.StoredVerificationCode = storedUserInfo.StackOverflowData.VerificationData.StoredVerificationCode
	user.StackOverflowData.VerificationData.IsVerified = storedUserInfo.StackOverflowData.VerificationData.IsVerified

	// update db with new Reddit profile info
	_, err = h.db.UpdateStackAboutInfo(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	log.Printf("[handler] successfully updated Stack Overflow info for user: %v\n", user.AuthUserID)
}
