package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// userAdd adds a single user listing to db
func userAdd(w http.ResponseWriter, r *http.Request) {
	user := &User{}

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
	userData, err := addUser(user)
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

// walletUpdate handles updates to the wallet address for a user
func walletUpdate(w http.ResponseWriter, r *http.Request) {
	user := &User{}

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
	updatedUserData, err := updateWallet(user)
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

// twoFAUpdate handles updates to the 2FA data for a user
func twoFAUpdate(w http.ResponseWriter, r *http.Request) {
	user := &User{}

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
	updatedUserData, err := updateTwoFA(user)
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

// usersGet handles queries to return all stored users
func usersGet(w http.ResponseWriter, r *http.Request) {
	users := &Users{}

	// return slice of structs of all user listings
	_, err := getUsers(users)
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

// userGet returns information about a single user
func userGet(w http.ResponseWriter, r *http.Request) {
	user := &User{}

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
	userData, err := getUser(user)
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

// userRemove removes a single user listing from db
func userRemove(w http.ResponseWriter, r *http.Request) {
	user := &User{}

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
	userData, err := removeUser(user)
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

// generateTwoEffEhCode generates a temporary 2FA code
func generateTwoEffEhCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable user of User struct
	user := &User{}

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

	// generate temporary 2FA code
	twoFACode := generateTwoFACode()

	// update local user object variable with generated 2FA code
	user.Info.TwoFAData.StoredTwoFACode = twoFACode

	// marshal into JSON
	_, err = json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// store user 2FA data in db
	userData, err := updateTwoFA(user)
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

// ValidateTwoEffEhCode validates the temporary 2FA code
func validateTwoEffEhCode(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := &User{}

	// initialize struct for reddit auth sessions
	authSession := new(AuthSessions)

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
	storedUserInfo, err := getUser(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update local user variable with stored info from db
	// user.Info.TwoFAData.StoredTwoFACode = storedUserInfo.Info.TwoFAData.StoredTwoFACode

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

	// update db with new info since 2FA codes matched
	userData, err := updateTwoFA(storedUserInfo)
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

// getRedditInfo returns Reddit profile info about the user
func redditUpdate(w http.ResponseWriter, r *http.Request) {
	// declare new variable of type User
	user := &User{}

	// initialize struct for reddit auth sessions
	authSession := new(AuthSessions)

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
	if err = authSession.GetAboutInfo(user); err != nil {
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
	if err = authSession.GetSubmittedInfo(user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// update db with new Reddit profile info
	userData, err := updateRedditInfo(user)
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
