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
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// add user listing to db
	userData, err := db.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully added user: %v\n\n", user.Info.ID)
}

// UsersGet handles queries to return all stored users
// TODO:
// return all users across multiple tables
// currently only returns reddit users
func UsersGet(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	users := new(db.Users)

	// return slice of structs of all user listings
	_, err := db.GetUsers(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// marshall into JSON
	_, err = json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&users); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully returned information for %d users\n\n", len(users.Users))
}

// UserGet returns information about a single user
func UserGet(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// get user listing by name
	userData, err := db.GetUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully returned information for user: %v\n\n", user.Info.ID)
}

// UserRemove removes a single user listing from db
func UserRemove(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// remove user listing from db
	userData, err := db.RemoveUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully deleted user: %v\n\n", user.Info.ID)
}

// WalletUpdate handles updates to the wallet address for a user
func WalletUpdate(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update the user listing in db
	updatedUserData, err := db.UpdateWallet(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&updatedUserData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully updated wallet address for user: %v\n\n", user.Info.ID)
}

// REDDIT

// UpdateRedditVerificationCode handles updates to the verification data for a user
func UpdateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update the user listing in db
	updatedUserData, err := db.UpdateRedditVerificationCode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&updatedUserData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully updated verification info for user: %v\n\n", user.Info.ID)
}

// GenerateRedditVerificationCode generates a temporary verification code
func GenerateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// generate temporary verification code
	verificationCode := verify.GenerateVerificationCode()

	// update local user object variable with generated verification code
	user.Info.RedditData.VerificationData.StoredVerificationCode = verificationCode

	// marshal into JSON
	_, err = json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// store user verification data in db
	userData, err := db.UpdateRedditVerificationCode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully generated verification code for user: %v\n\n", user.Info.ID)
}

// ValidateRedditVerificationCode validates the temporary verification code
func ValidateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := new(db.User)

	// initialize struct for reddit auth sessions
	authSession := new(reddit.AuthSessions)

	// initializes reddit OAuth sessions
	authSession.InitRedditAuth()

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// pull stored verification code + reddit username from DB
	storedUserInfo, err := db.GetUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// check reddit for matching verification code
	updatedUserObj, err := authSession.GetRecentPostsFromSubreddit(storedUserInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Checking %s against %s\n\n", updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode, storedUserInfo.Info.RedditData.VerificationData.StoredVerificationCode)

	// secondary validation of verification code
	if updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode != storedUserInfo.Info.RedditData.VerificationData.StoredVerificationCode {
		http.Error(w, err.Error(), 500)
		return
	}

	// update verification field values
	storedUserInfo.Info.RedditData.VerificationData.PostedVerificationCode = updatedUserObj.Info.RedditData.VerificationData.PostedVerificationCode
	storedUserInfo.Info.RedditData.VerificationData.IsVerified = true

	// update db with new info since verification codes matched
	userData, err := db.UpdateRedditVerificationCode(storedUserInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully validated verification code for user: %v\n\n", user.Info.ID)
}

// RedditUpdate returns Reddit profile info about the user
func RedditUpdate(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := new(db.User)

	// initialize struct for reddit auth sessions
	authSession := new(reddit.AuthSessions)

	// initializes reddit OAuth sessions
	authSession.InitRedditAuth()

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// pull stored reddit username from DB
	user, err = db.GetUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Reddit About info\n")
	// get general about info for user
	user, err = authSession.GetAboutInfo(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Reddit Trophy info\n")
	// get list of trophies user has been awarded
	if err = authSession.GetUserTrophies(user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Reddit Submitted info\n")
	// get slice of subreddits user is subscribed to based on activity
	user, err = authSession.GetSubmittedInfo(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update db with new Reddit profile info
	userData, err := db.UpdateRedditInfo(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully updated Reddit info for user: %v\n\n", user.Info.ID)
}

// STACK OVERFLOW

// StackUserAdd adds a single user listing to db
func StackUserAdd(w http.ResponseWriter, r *http.Request) {
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// add user listing to db
	userData, err := db.AddStackUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully added user: %v\n\n", user.Info.ID)
}

// StackUserGet returns information about a single user
func StackUserGet(w http.ResponseWriter, r *http.Request) {
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// get user listing by name
	userData, err := db.GetStackUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully returned information for user: %v\n\n", user.Info.ID)
}

// GenerateStackVerificationCode creates a verifcation code for Stack Overflow
func GenerateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// generate temporary verification code
	verificationCode := verify.GenerateVerificationCode()

	// promotional display code
	displayCode := fmt.Sprintf("[COINDROP.IO - IT PAYS TO CONTRIBUTE: %s]", verificationCode)

	// update local user object variable with generated verification code
	user.Info.StackOverflowData.VerificationData.StoredVerificationCode = displayCode

	// marshal into JSON
	_, err = json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// store user verification data in db
	stackUserData, err := db.UpdateStackVerificationCode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// if err := json.NewEncoder(w).Encode(&stackUserData); err != nil {
	if err := json.NewEncoder(w).Encode(&stackUserData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully generated verification code for user: %v\n\n", user.Info.ID)
}

// ValidateStackVerificationCode validates the temporary verification code
func ValidateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable user of User struct
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// pull stored verification code from DB
	storedStackUser, err := db.GetStackUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// check Stack Overflow for matching verification code
	updatedStackUser, err := stackoverflow.GetProfileByUserID(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Checking %s against %s\n\n", storedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode, storedStackUser.Info.StackOverflowData.VerificationData.StoredVerificationCode)

	// secondary validation to see if codes match
	if !strings.Contains(updatedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode, storedStackUser.Info.StackOverflowData.VerificationData.StoredVerificationCode) {
		log.Println("[!] Verification codes do not match!\n")
		http.Error(w, err.Error(), 500)
		return
	}

	// update verification field values
	storedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode = updatedStackUser.Info.StackOverflowData.VerificationData.PostedVerificationCode
	storedStackUser.Info.StackOverflowData.VerificationData.IsVerified = true

	// update db with new info since verification codes matched
	userData, err := db.UpdateStackVerificationCode(storedStackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully validated verification code for user: %v\n\n", user.Info.ID)
}

// StackUserUpdate updates and returns profile info about the user
func StackUserUpdate(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := new(db.User)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// pull stored verification code from DB
	user, err = db.GetStackUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Stack Overflow About info\n")
	// get general about info for user
	user, err = stackoverflow.GetProfileByUserID(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Stack Overflow associated accounts info\n")
	// get list of trophies user has been awarded
	user, err = stackoverflow.GetAssociatedAccounts(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update db with new Reddit profile info
	userData, err := db.UpdateStackAboutInfo(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully updated Stack Overflow info for user: %v\n\n", user.Info.ID)
}
