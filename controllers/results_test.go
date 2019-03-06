package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/waymobetta/go-coindrop-api/app"
)

func TestQuizResultsCreate(t *testing.T) {
	t.Run("add quiz results", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		quizID := "8295c792-5d65-4ba0-98cc-7f319967e628"
		userID := getUserID()
		questionsCorrect := 2
		questionsIncorrect := 1
		url := fmt.Sprintf("%s/v1/quizzes/results", svr.URL)

		t.Logf("URL: %s", url)

		client := &http.Client{}
		payload := []byte(fmt.Sprintf(`{"questionsCorrect": %v, "questionsIncorrect": %v, "userId": %q, "quizId": %q}`, questionsCorrect, questionsIncorrect, userID, quizID))

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Content-Type", "application/json")

		setAuth(req)
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp == nil {
			t.Fail()
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}
	})
}

func TestQuizResultsShow(t *testing.T) {
	t.Run("get quiz", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		userID := getUserID()
		quizID := "8295c792-5d65-4ba0-98cc-7f319967e628"
		url := fmt.Sprintf("%s/v1/quizzes/%s/results?userId=%s", svr.URL, quizID, userID)

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
			t.Errorf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		results := &app.Results{}
		err = json.Unmarshal(body, &results)
		if err != nil {
			t.Error(err)
		}

		if results.QuestionsCorrect == 0 {
			t.Error("expected correct > 0")
		}
	})
}

func TestQuizResultsList(t *testing.T) {
	t.Run("get quiz", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		userID := getUserID()
		url := fmt.Sprintf("%s/v1/quizzes/results?userId=%s", svr.URL, userID)

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
			t.Errorf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var results app.ResultsCollection
		err = json.Unmarshal(body, &results)
		if err != nil {
			t.Error(err)
		}

		if len(results) == 0 {
			t.Error("expected count > 0")
		}
	})
}
