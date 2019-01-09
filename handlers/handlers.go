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
	"github.com/waymobetta/go-coindrop-api/services/reddit"
	"github.com/waymobetta/go-coindrop-api/services/stackoverflow"
	"github.com/waymobetta/go-coindrop-api/utils"
	"github.com/waymobetta/go-coindrop-api/verify"
)

// TEST

// HandleIndex prints test to screen if successful
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}

// PROFILE

// UserAdd adds a single user listing to db
func UserAdd(w http.ResponseWriter, r *http.Request) {
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
	err = json.Unmarshal(body, user)
	if err != nil {
		response = utils.Message(false, "JSON Unmarshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// add user listing to db
	_, err = db.AddRedditUser(user)
	if err != nil {
		response = utils.Message(false, "Could not add Reddit user to db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	fmt.Printf("Successfully added user: %v\n\n", user.Info.AuthUserID)
}

// UsersGet handles queries to return all stored users
func UsersGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	users := new(db.Users)

	// return slice of structs of all user listings
	_, err := db.GetUsers(users)
	if err != nil {
		response = utils.Message(false, "Could not retrieve users from db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// marshall into JSON
	_, err = json.Marshal(users)
	if err != nil {
		response = utils.Message(false, "JSON Marshal error")
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

// UserGet returns information about a single user
func UserGet(w http.ResponseWriter, r *http.Request) {
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

	// get user listing by name
	userData, err := db.GetRedditUser(user)
	if err != nil {
		response = utils.Message(false, "Could not get Reddit user")
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

// UserRemove removes a single user listing from db
func UserRemove(w http.ResponseWriter, r *http.Request) {
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

	// remove user listing from db
	_, err = db.RemoveRedditUser(user)
	if err != nil {
		response = utils.Message(false, "Could not remove Reddit user")
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

// REDDIT

// UpdateRedditVerificationCode handles updates to the verification data for a user
func UpdateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
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
	_, err = db.UpdateRedditVerificationCode(user)
	if err != nil {
		response = utils.Message(false, "Could not update Reddit verification code")
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

	// generate temporary verification code
	verificationCode := verify.GenerateVerificationCode()

	// update local user object variable with generated verification code
	user.Info.RedditData.VerificationData.StoredVerificationCode = verificationCode

	// marshal into JSON
	_, err = json.Marshal(&user)
	if err != nil {
		response = utils.Message(false, "JSON Marshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// store user verification data in db
	_, err = db.UpdateRedditVerificationCode(user)
	if err != nil {
		response = utils.Message(false, "Could not update Reddit verification code")
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

	// pull stored verification code + reddit username from DB
	storedUserInfo, err := db.GetRedditUser(user)
	if err != nil {
		response = utils.Message(false, "Could not get Reddit user")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// check reddit for matching verification code
	updatedUserObj, err := authSession.GetRecentPostsFromSubreddit(storedUserInfo)
	if err != nil {
		response = utils.Message(false, "Could not get recent subreddit posts")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	fmt.Printf("Checking %s against %s\n\n", updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode, storedUserInfo.Info.RedditData.VerificationData.StoredVerificationCode)

	// secondary validation of verification code
	if updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode != storedUserInfo.Info.RedditData.VerificationData.StoredVerificationCode {
		response = utils.Message(false, "Verification codes do not match")
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
		response = utils.Message(false, "Could not update Reddit verification code")
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

	// pull stored reddit username from DB
	user, err = db.GetRedditUser(user)
	if err != nil {
		response = utils.Message(false, "Could not get Reddit user")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Reddit About info\n")
	// get general about info for user
	user, err = authSession.GetAboutInfo(user)
	if err != nil {
		response = utils.Message(false, "Could not get Reddit user about info")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Reddit Trophy info\n")
	// get list of trophies user has been awarded
	if err = authSession.GetRedditUserTrophies(user); err != nil {
		response = utils.Message(false, "Could not get Reddit user trophies")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Reddit Submitted info\n")
	// get slice of subreddits user is subscribed to based on activity
	user, err = authSession.GetSubmittedInfo(user)
	if err != nil {
		response = utils.Message(false, "Could not get Reddit user submitted info")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// update db with new Reddit profile info
	_, err = db.UpdateRedditInfo(user)
	if err != nil {
		response = utils.Message(false, "Could not update db with Reddit info")
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

// STACK OVERFLOW

// StackUserAdd adds a single user listing to db
func StackUserAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
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

	// add user listing to db
	_, err = db.AddStackUser(user)
	if err != nil {
		response = utils.Message(false, "Could not add Stack Overflow user")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	fmt.Printf("Successfully added user: %v\n\n", user.Info.AuthUserID)
}

// StackUserGet returns information about a single user
func StackUserGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
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

	// get user listing by name
	userData, err := db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, "Could not get Stack Overflow user")
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
		response = utils.Message(false, "JSON Marshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// store user verification data in db
	_, err = db.UpdateStackVerificationCode(user)
	if err != nil {
		response = utils.Message(false, "Could not update Stack Overflow verification code")
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

	// pull stored verification code from DB
	storedStackUser, err := db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, "Could not get user from db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// check Stack Overflow for matching verification code
	updatedStackUser, err := stackoverflow.GetProfileByUserID(user)
	if err != nil {
		response = utils.Message(false, "Could not return profile information")
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
		response = utils.Message(false, "Could not update db with verification code")
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

	// pull stored verification code from DB
	storedUserInfo, err := db.GetStackUser(user)
	if err != nil {
		response = utils.Message(false, "Could not get Stack Overflow user info")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Stack Overflow About info\n")
	// get general about info for user
	_, err = stackoverflow.GetProfileByUserID(user)
	if err != nil {
		response = utils.Message(false, "Could not get Stack Overflow user about me info")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	log.Println("[+] Retrieving Stack Overflow associated accounts info\n")
	// get list of trophies user has been awarded
	_, err = stackoverflow.GetAssociatedAccounts(user)
	if err != nil {
		response = utils.Message(false, "Could not get Stack Overflow user associated accounts info")
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
		response = utils.Message(false, "Could not update db with Stack Overflow profile info")
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

// TasksGet handles queries to return all stored tasks
func TasksGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	tasks := new(db.Tasks)

	// return slice of structs of all task listings
	_, err := db.GetTasks(tasks)
	if err != nil {
		response = utils.Message(false, "Could not retrieve tasks from db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// marshall into JSON
	_, err = json.Marshal(tasks)
	if err != nil {
		response = utils.Message(false, "JSON Marshal error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, tasks)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned information for %d tasks\n\n", len(tasks.Tasks))
}
