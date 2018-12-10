package coindrophandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/waymobetta/go-coindrop-api/coindropverification"
	"github.com/waymobetta/go-coindrop-api/goreddit"
	"github.com/waymobetta/go-coindrop-api/gostackoverflow"
)

// PROFILE

// UserAdd adds a single user listing to db
func UserAdd(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(goreddit.User)

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
	userData, err := goreddit.AddUser(user)
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

	fmt.Printf("Successfully added user: %s\n\n", user.Info.RedditData.Username)
}

// UsersGet handles queries to return all stored users
func UsersGet(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	users := new(goreddit.Users)

	// return slice of structs of all user listings
	_, err := goreddit.GetUsers(users)
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
	user := new(goreddit.User)

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
	userData, err := goreddit.GetUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully returned information for user: %s\n\n", user.Info.RedditData.Username)
}

// UserRemove removes a single user listing from db
func UserRemove(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(goreddit.User)

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
	userData, err := goreddit.RemoveUser(user)
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

	fmt.Printf("Successfully deleted user: %s\n\n", user.Info.RedditData.Username)
}

// WalletUpdate handles updates to the wallet address for a user
func WalletUpdate(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(goreddit.User)

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
	updatedUserData, err := goreddit.UpdateWallet(user)
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

	fmt.Printf("Successfully updated wallet address for user: %s\n\n", user.Info.RedditData.Username)
}

// REDDIT

// UpdateRedditVerificationCode handles updates to the 2FA data for a user
func UpdateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(goreddit.User)

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
	updatedUserData, err := coindropverification.UpdateRedditVerificationCode(user)
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

	fmt.Printf("Successfully updated 2FA info for user: %s\n\n", user.Info.RedditData.Username)
}

// GenerateRedditVerificationCode generates a temporary 2FA code
func GenerateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	// initialize new variable user of User struct
	user := new(goreddit.User)

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
	twoFACode := coindropverification.GenerateVerificationCode()

	// update local user object variable with generated 2FA code
	user.Info.TwoFAData.StoredTwoFACode = twoFACode

	// marshal into JSON
	_, err = json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// store user 2FA data in db
	userData, err := coindropverification.UpdateRedditVerificationCode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully generated 2FA code for user: %s\n\n", user.Info.RedditData.Username)
}

// ValidateRedditVerificationCode validates the temporary 2FA code
func ValidateRedditVerificationCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := new(goreddit.User)

	// initialize struct for reddit auth sessions
	authSession := new(goreddit.AuthSessions)

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

	// pull stored 2FA code from DB
	storedUserInfo, err := goreddit.GetUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// check reddit for matching 2FA code
	updatedUserObj, err := authSession.GetRecentPostsFromSubreddit(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Checking %s against %s\n\n", updatedUserObj.Info.TwoFAData.PostedTwoFACode, storedUserInfo.Info.TwoFAData.StoredTwoFACode)

	// secondary validation of 2FA code
	if updatedUserObj.Info.TwoFAData.PostedTwoFACode != storedUserInfo.Info.TwoFAData.StoredTwoFACode {
		http.Error(w, err.Error(), 500)
		return
	}

	// update 2FA field values
	storedUserInfo.Info.TwoFAData.PostedTwoFACode = updatedUserObj.Info.TwoFAData.PostedTwoFACode
	storedUserInfo.Info.TwoFAData.IsValidated = true

	// update db with new info since verification codes matched
	userData, err := coindropverification.UpdateRedditVerificationCode(storedUserInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully validated 2FA code for user: %s\n\n", user.Info.RedditData.Username)
}

// RedditUpdate returns Reddit profile info about the user
func RedditUpdate(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := new(goreddit.User)

	// initialize struct for reddit auth sessions
	authSession := new(goreddit.AuthSessions)

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
	userData, err := goreddit.UpdateRedditInfo(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully updated Reddit info for user: %s\n\n", user.Info.RedditData.Username)
}

// STACK OVERFLOW

// StackUserAdd adds a single user listing to db
func StackUserAdd(w http.ResponseWriter, r *http.Request) {
	stackUser := new(gostackoverflow.StackOverflowData)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// add user listing to db
	userData, err := gostackoverflow.AddStackUser(stackUser)
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

	fmt.Printf("Successfully added user: %v\n\n", stackUser.UserID)
}

// StackUserGet returns information about a single user
func StackUserGet(w http.ResponseWriter, r *http.Request) {
	stackUser := new(gostackoverflow.StackOverflowData)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// get user listing by name
	userData, err := gostackoverflow.GetStackUser(stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully returned information for user: %v\n\n", stackUser.UserID)
}

// GenerateStackVerificationCode creates a verifcation code for Stack Overflow
func GenerateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable user of StackUser struct
	stackUser := new(gostackoverflow.StackOverflowData)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// generate temporary 2FA code
	verificationCode := coindropverification.GenerateVerificationCode()

	// promotional display code
	displayCode := fmt.Sprintf("[COINDROP.IO - IT PAYS TO CONTRIBUTE: %s]", verificationCode)

	// update local user object variable with generated verification code
	stackUser.VerificationData.StoredVerificationCode = displayCode

	// marshal into JSON
	_, err = json.Marshal(&stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// store user verification data in db
	stackUserData, err := gostackoverflow.UpdateVerificationCode(stackUser)
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

	fmt.Printf("Successfully generated verification code for user: %v\n\n", stackUser.UserID)
}

// ValidateStackVerificationCode validates the temporary verification code
func ValidateStackVerificationCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable user of StackUser struct
	stackUser := new(gostackoverflow.StackOverflowData)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// pull stored verification code from DB
	storedStackUser, err := gostackoverflow.GetStackUser(stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// check Stack Overflow for matching verification code
	updatedStackUser, err := stackUser.GetProfileByUserID()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Checking %s against %s\n\n", updatedStackUser.VerificationData.PostedVerificationCode, storedStackUser.VerificationData.StoredVerificationCode)

	// secondary validation to see if codes match
	if !strings.Contains(updatedStackUser.VerificationData.PostedVerificationCode, storedStackUser.VerificationData.StoredVerificationCode) {
		log.Println("[!] Verification codes do not match!\n")
		http.Error(w, err.Error(), 500)
		return
	}

	// update 2FA field values
	storedStackUser.VerificationData.PostedVerificationCode = updatedStackUser.VerificationData.PostedVerificationCode
	storedStackUser.VerificationData.IsVerified = true

	// update db with new info since 2FA codes matched
	userData, err := gostackoverflow.UpdateVerificationCode(storedStackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully validated verification code for user: %v\n\n", stackUser.UserID)
}

// StackUserUpdate updates and returns profile info about the user
func StackUserUpdate(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	stackUser := new(gostackoverflow.StackOverflowData)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, &stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Stack Overflow About info\n")
	// get general about info for user
	stackUser, err = stackUser.GetProfileByUserID()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println("[+] Retrieving Stack Overflow associated accounts info\n")
	// get list of trophies user has been awarded
	stackUser, err = stackUser.GetAssociatedAccounts()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update db with new Reddit profile info
	userData, err := gostackoverflow.UpdateStackAboutInfo(stackUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&userData); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Successfully updated Reddit info for user: %v\n\n", stackUser.UserID)
}
