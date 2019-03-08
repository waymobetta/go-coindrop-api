package stackoverflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/waymobetta/go-coindrop-api/types"
)

const (
	// baseAPI is the base URL for the Stack Overflow API
	baseAPI = "https://api.stackexchange.com/2.2"
)

var (
	// noVerifError is generated if the verification code does not match
	noVerifError = errors.New("verification code does not match")
)

// GetProfileByUserID fetches basic user profile info by unique user ID
func GetProfileByUserID(u *types.User) error {
	log.Printf("[stackoverflow] collecting profile information for user ID: %v\n", u.Social.StackOverflow.StackUserID)

	profileEndpoint := fmt.Sprintf("/users/%v?order=desc&sort=reputation&site=stackoverflow&filter=!-*jbN*IioeFP", u.Social.StackOverflow.StackUserID)

	// concatenate endpoint with base URL of db
	profileURL := fmt.Sprintf("%s%s", baseAPI, profileEndpoint)

	// prepare GET request
	req, err := http.NewRequest("GET", profileURL, nil)
	if err != nil {
		log.Errorf("[services/stackoverflow] Error preparing GET request for user profile info; %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	// execute GET request
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("[services/stackoverflow] Error fetching user profile info; %v\n", err)
		return err
	}
	defer res.Body.Close()

	// read result of GET request
	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[services/stackoverflow] Error reading response body; %v\n", err)
		return err
	}

	// initialize new struct to contain AboutProfileResponse
	aboutProfResStruct := new(AboutProfileResponse)

	// unmarshal JSON into AboutProfileResponse struct
	if err := json.Unmarshal(byteArr, &aboutProfResStruct); err != nil {
		log.Errorf("[services/stackoverflow] Error unmarshalling JSON; %v\n", err)
		return err
	}

	log.Printf("[services/stackoverflow] found profile info for user: %s!\n", aboutProfResStruct.Items[0].DisplayName)

	// initialize empty accounts slice
	accounts := []string{}

	// iterate over number of items in the response
	// NOTE: there should only be a single item
	for index := range aboutProfResStruct.Items {
		u = &types.User{
			Social: &types.Social{
				StackOverflow: &types.StackOverflow{
					DisplayName:       aboutProfResStruct.Items[0].DisplayName,
					ExchangeAccountID: aboutProfResStruct.Items[index].AccountID,
					StackUserID:       aboutProfResStruct.Items[index].UserID,
					Accounts:          accounts,
					Communities: []types.Community{
						types.Community{
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
					Verification: &types.Verification{
						PostedVerificationCode:    aboutProfResStruct.Items[index].AboutMe,
						ConfirmedVerificationCode: u.Social.StackOverflow.Verification.ConfirmedVerificationCode,
						Verified:                  u.Social.StackOverflow.Verification.Verified,
					},
				},
			},
		}
		return nil
	}
	return nil
}

// GetAssociatedAccounts method fetches associated communities of user
func GetAssociatedAccounts(u *types.User) error {
	log.Printf("[services/stackoverflow] collecting associated account information for user: %s\n", u.Social.StackOverflow.DisplayName)

	associatedAccountsEndpoint := fmt.Sprintf("/users/%v/associated", u.Social.StackOverflow.ExchangeAccountID)

	// concatenate endpoint with base URL of db
	associatedAccountsURL := fmt.Sprintf("%s%s", baseAPI, associatedAccountsEndpoint)

	// prepare GET request
	req, err := http.NewRequest("GET", associatedAccountsURL, nil)
	if err != nil {
		log.Errorf("[services/stackoverflow] Error preparing GET request for user associated accounts info; %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	// execute GET request
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("[services/stackoverflow] Error fetching user profile info; %v\n", err)
		return err
	}
	defer res.Body.Close()

	// return if not 200
	if res.StatusCode != http.StatusOK {
		return err
	}

	// read result of GET request
	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[services/stackoverflow] Error reading response body\n%v", err)
		return err
	}

	// initialize new struct to contain AssociatedCommunitiesResponse
	associatedCommunitiesStruct := new(AssociatedCommunitiesResponse)

	// unmarshal JSON into AssociatedCommunitiesResponse struct
	if err := json.Unmarshal(byteArr, &associatedCommunitiesStruct); err != nil {
		log.Errorf("[services/stackoverflow] Error unmarshalling JSON; %v\n", err)
		return err
	}

	log.Printf("[services/stackoverflow] Found associated account info for user: %s!\n", u.Social.StackOverflow.DisplayName)

	// initialize new struct object to hold Community data
	communityObj := types.Community{}

	// iterate over number of items in the response
	// NOTE: there could be multiple items
	for index := range associatedCommunitiesStruct.Items {
		// for each item, overwrite struct object to hold updated data
		communityObj = types.Community{
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
		u.Social.StackOverflow.Accounts = append(u.Social.StackOverflow.Accounts, associatedCommunitiesStruct.Items[index].SiteName)

		// append updated struct data to slice of objects
		u.Social.StackOverflow.Communities = append(u.Social.StackOverflow.Communities, communityObj)
	}

	return nil
}

// VerificationCheck checks posted verif. code against that which is stored
func VerificationCheck(u *types.User) error {
	// secondary validation to see if codes match
	if !strings.Contains(
		u.Social.StackOverflow.Verification.PostedVerificationCode,
		u.Social.StackOverflow.Verification.ConfirmedVerificationCode,
	) {
		return noVerifError
	}

	// if no error, update verification field values
	u.Social.StackOverflow.Verification.Verified = true

	return nil
}
