package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/stackoverflow"
	"github.com/waymobetta/go-coindrop-api/utils"
	"github.com/waymobetta/go-coindrop-api/verify"
)

// STACK OVERFLOW

// StackUserAdd adds a single user listing to db
func StackUserAdd(w http.ResponseWriter, r *http.Request) {
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
	_, err = db.AddStackUser(user)
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

	fmt.Printf("Successfully added stack user: %v\n\n", user.Info.AuthUserID)
}

// StackUserGet returns information about a single user
func StackUserGet(w http.ResponseWriter, r *http.Request) {
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
	userData, err := db.GetStackUser(user)
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

	fmt.Printf("Successfully returned information for user: %v\n\n", user.Info.AuthUserID)
}

// GenerateStackVerificationCode creates a verifcation code for Stack Overflow
func GenerateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
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
	displayCode := fmt.Sprintf("[COINDROP.IO - IT PAYS TO CONTRIBUTE: %s]", verificationCode)

	// update local user object variable with generated verification code
	user.Info.StackOverflowData.VerificationData.StoredVerificationCode = displayCode
	user.Info.StackOverflowData.UserID = user.Info.StackOverflowData.UserID
	user.Info.AuthUserID = user.Info.AuthUserID

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
	_, err = db.UpdateStackVerificationCode(user)
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

	fmt.Printf("Successfully generated verification code for user: %v\n\n", user.Info.AuthUserID)
}

// ValidateStackVerificationCode validates the temporary verification code
func ValidateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
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
	storedStackUser, err := db.GetStackUser(user)
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

	fmt.Printf("Checking %s against %s\n\n", storedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode, storedStackUser.Info.StackOverflowData.VerificationData.StoredVerificationCode)

	// secondary validation to see if codes match
	if !strings.Contains(updatedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode, storedStackUser.Info.StackOverflowData.VerificationData.StoredVerificationCode) {
		log.Println("[!] Verification codes do not match!\n")

		response = utils.Message(false, "unsuccessful")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update verification field values
	storedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode = updatedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode
	storedStackUser.Info.StackOverflowData.VerificationData.IsVerified = true

	// update db with new info since verification codes matched
	_, err = db.UpdateStackVerificationCode(storedStackUser)
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

	fmt.Printf("Successfully validated verification code for user: %v\n\n", user.Info.AuthUserID)
}

// StackUserUpdate updates and returns profile info about the user
func StackUserUpdate(w http.ResponseWriter, r *http.Request) {
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
	storedUserInfo, err := db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Stack Overflow About info\n")
	// get general about info for user
	_, err = stackoverflow.GetProfileByUserID(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Stack Overflow associated accounts info\n")
	// get list of trophies user has been awarded
	_, err = stackoverflow.GetAssociatedAccounts(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	user.Info.StackOverflowData.VerificationData.PostedVerificationCode = storedUserInfo.Info.StackOverflowData.VerificationData.PostedVerificationCode
	user.Info.StackOverflowData.VerificationData.StoredVerificationCode = storedUserInfo.Info.StackOverflowData.VerificationData.StoredVerificationCode
	user.Info.StackOverflowData.VerificationData.IsVerified = storedUserInfo.Info.StackOverflowData.VerificationData.IsVerified

	// update db with new Reddit profile info
	_, err = db.UpdateStackAboutInfo(user)
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

	fmt.Printf("Successfully updated Stack Overflow info for user: %v\n\n", user.Info.AuthUserID)
}
