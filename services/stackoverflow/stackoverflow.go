package stackoverflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	fmt.Printf("Collecting profile information for user ID: %v\n\n", u.Info.StackOverflowData.UserID)

	profileEndpoint := fmt.Sprintf("/users/%v?order=desc&sort=reputation&site=stackoverflow&filter=!-*jbN*IioeFP", u.Info.StackOverflowData.UserID)

	// concatenate endpoint with base URL of db
	profileURL := fmt.Sprintf("%s%s", baseAPI, profileEndpoint)

	// prepare GET request
	req, err := http.NewRequest("GET", profileURL, nil)
	if err != nil {
		err = fmt.Errorf("[!] Error preparing GET request for user profile info\n%v", err)
		return u, err
	}
	req.Header.Set("Content-Type", "application/json")

	// execute GET request
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("[!] Error fetching user profile info\n%v", err)
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
		err = fmt.Errorf("[!] Error reading response body\n%v", err)
		return u, err
	}

	// initialize new struct to contain AboutProfileResponse
	aboutProfResStruct := new(AboutProfileResponse)

	// unmarshal JSON into AboutProfileResponse struct
	if err := json.Unmarshal(byteArr, &aboutProfResStruct); err != nil {
		err = fmt.Errorf("[!] Error unmarshalling JSON\n%v", err)
		return u, err
	}

	fmt.Printf("[+] Found profile info for user: %s!\n", aboutProfResStruct.Items[0].DisplayName)

	// initialize empty accounts slice
	accounts := []string{}

	// iterate over number of items in the response
	// NOTE: there should only be a single item
	for index := range aboutProfResStruct.Items {
		u.Info.StackOverflowData = db.StackOverflowData{
			DisplayName:       aboutProfResStruct.Items[index].DisplayName,
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
			VerificationData: verify.VerificationData{
				PostedVerificationCode: aboutProfResStruct.Items[index].AboutMe,
				StoredVerificationCode: u.Info.StackOverflowData.VerificationData.StoredVerificationCode,
				IsVerified:             false,
			},
		}
		return u, nil
	}
	return u, nil
}

// GetAssociatedAccounts method fetches associated communities of user
func GetAssociatedAccounts(u *db.User) (*db.User, error) {
	fmt.Printf("Collecting associated account information for user: %s\n", u.Info.StackOverflowData.DisplayName)

	associatedAccountsEndpoint := fmt.Sprintf("/users/%v/associated", u.Info.StackOverflowData.ExchangeAccountID)

	// concatenate endpoint with base URL of db
	associatedAccountsURL := fmt.Sprintf("%s%s", baseAPI, associatedAccountsEndpoint)

	// prepare GET request
	req, err := http.NewRequest("GET", associatedAccountsURL, nil)
	if err != nil {
		err = fmt.Errorf("[!] Error preparing GET request for user associated accounts info\n%v", err)
		return u, err
	}
	req.Header.Set("Content-Type", "application/json")

	// execute GET request
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("[!] Error fetching user profile info\n%v", err)
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
		err = fmt.Errorf("[!] Error reading response body\n%v", err)
		return u, err
	}

	// initialize new struct to contain AssociatedCommunitiesResponse
	associatedCommunitiesStruct := new(AssociatedCommunitiesResponse)

	// unmarshal JSON into AssociatedCommunitiesResponse struct
	if err := json.Unmarshal(byteArr, &associatedCommunitiesStruct); err != nil {
		err = fmt.Errorf("[!] Error unmarshalling JSON\n%v", err)
		return u, err
	}

	fmt.Printf("[+] Found associated account info for user: %s!\n", u.Info.StackOverflowData.DisplayName)

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
		u.Info.StackOverflowData.Accounts = append(u.Info.StackOverflowData.Accounts, associatedCommunitiesStruct.Items[index].SiteName)

		// append updated struct data to slice of objects
		u.Info.StackOverflowData.Communities = append(u.Info.StackOverflowData.Communities, communityObj)
	}

	return u, nil
}
