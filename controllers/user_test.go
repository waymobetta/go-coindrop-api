package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/waymobetta/go-coindrop-api/app"
)

func TestUserShow(t *testing.T) {
	t.Run("show user", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		userID := getUserID()
		url := fmt.Sprintf("%s/v1/users/%s", svr.URL, userID)

		t.Logf("URL: %s", url)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Error(err)
		}

		setAuth(req)
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		user := &app.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Error(err)
		}

		if user.ID == "" {
			t.Error("expected user ID")
		}
	})
}

func TestUserList(t *testing.T) {
	t.Run("show user by cognito auth user ID", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		cognitoUserID := getCognitoAuthUserID()
		url := fmt.Sprintf("%s/v1/users/?cognitoAuthUserId=%s", svr.URL, cognitoUserID)

		t.Logf("URL: %s", url)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Error(err)
		}

		setAuth(req)
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		user := &app.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Error(err)
		}

		if user.ID == "" {
			t.Error("expected user ID")
		}
	})
}
