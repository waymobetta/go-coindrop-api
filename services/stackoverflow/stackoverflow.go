package stackoverflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/verify"
)

const (
	// baseAPI is the base URL for the Stack Overflow API
	baseAPI = "https://api.stackexchange.com/2.2"
)

var (
	// TODO:
	// remove global client
	client = &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}
)

// GetProfileByUserID fetches basic user profile info by unique user ID
func GetProfileByUserID(u *db.User) (*db.User, error) {
	log.Printf("[stackoverflow] collecting profile information for user ID: %v\n", u.StackOverflow.UserID)

	profileEndpoint := fmt.Sprintf("/users/%v?order=desc&sort=reputation&site=stackoverflow&filter=!-*jbN*IioeFP", u.StackOverflow.UserID)

	// concatenate endpoint with base URL of db
	profileURL := fmt.Sprintf("%s%s", baseAPI, profileEndpoint)

	// prepare GET request
	req, err := http.NewRequest("GET", profileURL, nil)
	if err != nil {
		log.Errorf("[stackoverflow] Error preparing GET request for user profile info; %v\n", err)
		return u, err
	}
	req.Header.Set("Content-Type", "application/json")

	// execute GET request
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("[stackoverflow] Error fetching user profile info; %v\n", err)
		return u, err
	}
	defer res.Body.Close()

	// read result of GET request
	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[stackoverflow] Error reading response body; %v\n", err)
		return u, err
	}

	// initialize new struct to contain AboutProfileResponse
	aboutProfResStruct := new(AboutProfileResponse)

	// unmarshal JSON into AboutProfileResponse struct
	if err := json.Unmarshal(byteArr, &aboutProfResStruct); err != nil {
		log.Errorf("[stackoverflow] Error unmarshalling JSON; %v\n", err)
		return u, err
	}

	log.Printf("[stackoverflow] found profile info for user: %s!\n", aboutProfResStruct.Items[0].DisplayName)

	// initialize empty accounts slice
	accounts := []string{}

	// iterate over number of items in the response
	// NOTE: there should only be a single item
	for index := range aboutProfResStruct.Items {
		u.StackOverflow = db.StackOverflow{
			DisplayName:       aboutProfResStruct.Items[0].DisplayName,
			ExchangeAccountID: aboutProfResStruct.Items[index].AccountID,
			UserID:            aboutProfResStruct.Items[index].UserID,
			Accounts:          accounts,
			Communities: []db.Community{
				db.Community{
					Name:          "stackoverflow",
					Reputation:    aboutProfResStruct.Items[index].Reputation,
					QuestionCount: aboutProfResStruct.Items[index].QuestionCount,
					AnswerCount:   aboutProfResStruct.Items[index].AnswerCount,
					BadgeCounts: map[string]int{
						"Bronze": aboutProfResStruct.Items[index].BadgeCounts.Bronze,
						"Silver": aboutProfResStruct.Items[index].BadgeCounts.Silver,
						"Gold":   aboutProfResStruct.Items[index].BadgeCounts.Gold,
					},
				},
			},
			Verification: &verify.Verification2{
				PostedVerificationCode:    aboutProfResStruct.Items[index].AboutMe,
				ConfirmedVerificationCode: u.StackOverflow.Verification.ConfirmedVerificationCode,
				Verified:                  false,
			},
		}
		return u, nil
	}
	return u, nil
}

// GetAssociatedAccounts method fetches associated communities of user
func GetAssociatedAccounts(u *db.User) (*db.User, error) {
	log.Printf("[stackoverflow] collecting associated account information for user: %s\n", u.StackOverflow.DisplayName)

	associatedAccountsEndpoint := fmt.Sprintf("/users/%v/associated", u.StackOverflow.ExchangeAccountID)

	// concatenate endpoint with base URL of db
	associatedAccountsURL := fmt.Sprintf("%s%s", baseAPI, associatedAccountsEndpoint)

	// prepare GET request
	req, err := http.NewRequest("GET", associatedAccountsURL, nil)
	if err != nil {
		log.Errorf("[stackoverflow] Error preparing GET request for user associated accounts info; %v\n", err)
		return u, err
	}
	req.Header.Set("Content-Type", "application/json")

	// execute GET request
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("[stackoverflow] Error fetching user profile info; %v\n", err)
		return u, err
	}
	defer res.Body.Close()

	// return if not 200
	if res.StatusCode != http.StatusOK {
		return u, err
	}

	// read result of GET request
	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[stackoverflow] Error reading response body\n%v", err)
		return u, err
	}

	// initialize new struct to contain AssociatedCommunitiesResponse
	associatedCommunitiesStruct := new(AssociatedCommunitiesResponse)

	// unmarshal JSON into AssociatedCommunitiesResponse struct
	if err := json.Unmarshal(byteArr, &associatedCommunitiesStruct); err != nil {
		log.Errorf("[stackoverflow] Error unmarshalling JSON; %v\n", err)
		return u, err
	}

	log.Printf("[stackoverflow] Found associated account info for user: %s!\n", u.StackOverflow.DisplayName)

	// initialize new struct object to hold Community data
	communityObj := db.Community{}

	// iterate over number of items in the response
	// NOTE: there could be multiple items
	for index := range associatedCommunitiesStruct.Items {
		// for each item, overwrite struct object to hold updated data
		communityObj = db.Community{
			Name:          associatedCommunitiesStruct.Items[index].SiteName,
			Reputation:    associatedCommunitiesStruct.Items[index].Reputation,
			QuestionCount: associatedCommunitiesStruct.Items[index].QuestionCount,
			AnswerCount:   associatedCommunitiesStruct.Items[index].AnswerCount,
			BadgeCounts: map[string]int{
				"Bronze": associatedCommunitiesStruct.Items[index].BadgeCounts.Bronze,
				"Silver": associatedCommunitiesStruct.Items[index].BadgeCounts.Silver,
				"Gold":   associatedCommunitiesStruct.Items[index].BadgeCounts.Gold,
			},
		}

		// append community name to account slice
		u.StackOverflow.Accounts = append(u.StackOverflow.Accounts, associatedCommunitiesStruct.Items[index].SiteName)

		// append updated struct data to slice of objects
		u.StackOverflow.Communities = append(u.StackOverflow.Communities, communityObj)
	}

	return u, nil
}
