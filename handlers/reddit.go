package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/reddit"
	"github.com/waymobetta/go-coindrop-api/utils"
	"github.com/waymobetta/go-coindrop-api/verify"
)

// REDDIT

// RedditUserAdd adds a single user listing to db
func RedditUserAdd(w http.ResponseWriter, r *http.Request) {
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
	err = json.Unmarshal(body, user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// add user listing to db
	_, err = db.AddRedditUser(user)
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

	fmt.Printf("Successfully added reddit user: %v\n\n", user.Info.AuthUserID)
}

// UsersGet handles queries to return all stored users
func UsersGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	users := new(db.Users)

	// return slice of structs of all user listings
	_, err := db.GetUsers(users)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// marshall into JSON
	_, err = json.Marshal(users)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, users)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned information for %d users\n\n", len(users.Users))
}

// RedditUserGet returns information about a single user
func RedditUserGet(w http.ResponseWriter, r *http.Request) {
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

	// get user listing by name
	userData, err := db.GetRedditUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, &userData)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned information for user: %v\n\n", user.Info.AuthUserID)
}

// RedditUserRemove removes a single user listing from db
func RedditUserRemove(w http.ResponseWriter, r *http.Request) {
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

	// remove user listing from db
	_, err = db.RemoveRedditUser(user)
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

	fmt.Printf("Successfully deleted user: %v\n\n", user.Info.AuthUserID)
}

// REDDIT

// UpdateRedditVerificationCode handles updates to the verification data for a user
func UpdateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
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
	_, err = db.UpdateRedditVerificationCode(user)
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

	fmt.Printf("Successfully updated verification info for user: %v\n\n", user.Info.AuthUserID)
}

// GenerateRedditVerificationCode generates a temporary verification code
func GenerateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
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

	// generate temporary verification code
	verificationCode := verify.GenerateVerificationCode()

	// update local user object variable with generated verification code
	user.Info.RedditData.VerificationData.StoredVerificationCode = verificationCode

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
	_, err = db.UpdateRedditVerificationCode(user)
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

// ValidateRedditVerificationCode validates the temporary verification code
func ValidateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// declare new variable of type User
	user := new(db.User)

	// initialize struct for reddit auth sessions
	authSession := new(reddit.AuthSessions)

	// initializes reddit OAuth sessions
	authSession.InitRedditAuth()

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

	// pull stored verification code + reddit username from DB
	storedUserInfo, err := db.GetRedditUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// check reddit for matching verification code
	updatedUserObj, err := authSession.GetRecentPostsFromSubreddit(storedUserInfo)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	fmt.Printf("Checking %s against %s\n\n", updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode, storedUserInfo.Info.RedditData.VerificationData.StoredVerificationCode)

	// secondary validation of verification code
	if updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode != storedUserInfo.Info.RedditData.VerificationData.StoredVerificationCode {
		response = utils.Message(false, "unsuccessful")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update verification field values
	storedUserInfo.Info.RedditData.VerificationData.PostedVerificationCode = updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode
	storedUserInfo.Info.RedditData.VerificationData.IsVerified = true

	// update db with new info since verification codes matched
	_, err = db.UpdateRedditVerificationCode(storedUserInfo)
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

// RedditUpdate returns Reddit profile info about the user
func RedditUpdate(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// declare new variable of type User
	user := new(db.User)

	// initialize struct for reddit auth sessions
	authSession := new(reddit.AuthSessions)

	// initializes reddit OAuth sessions
	authSession.InitRedditAuth()

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

	// pull stored reddit username from DB
	user, err = db.GetRedditUser(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Reddit About info\n")
	// get general about info for user
	user, err = authSession.GetAboutInfo(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Reddit Trophy info\n")
	// get list of trophies user has been awarded
	if err = authSession.GetRedditUserTrophies(user); err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Reddit Submitted info\n")
	// get slice of subreddits user is subscribed to based on activity
	user, err = authSession.GetSubmittedInfo(user)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update db with new Reddit profile info
	_, err = db.UpdateRedditInfo(user)
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

	fmt.Printf("Successfully updated Reddit info for user: %v\n\n", user.Info.AuthUserID)
}
