package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/waymobetta/go-coindrop-api/app"
)

func TestQuizCreate(t *testing.T) {
	t.Run("get quiz", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		url := fmt.Sprintf("%s/v1/quizzes", svr.URL)

		t.Logf("URL: %s", url)

		client := &http.Client{}
		title := "hello world"
		payload := []byte(fmt.Sprintf(`{"title": %q}`, title))

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

		if resp.StatusCode != http.StatusOK {
			t.Errorf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		quiz := &app.Quiz{}
		err = json.Unmarshal(body, &quiz)
		if err != nil {
			t.Error(err)
		}

		if quiz.Title != title {
			t.Error("expected quiz title")
		}
	})
}

func TestQuizShow(t *testing.T) {
	t.Run("get quiz", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		quizID := "8295c792-5d65-4ba0-98cc-7f319967e628"
		url := fmt.Sprintf("%s/v1/quizzes/%s", svr.URL, quizID)

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

		quiz := &app.Quiz{}
		err = json.Unmarshal(body, &quiz)
		if err != nil {
			t.Error(err)
		}

		if quiz.Title == "" {
			t.Error("expected quiz title")
		}
	})
}

func TestQuizList(t *testing.T) {
	t.Run("get list of quizzes", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		url := fmt.Sprintf("%s/v1/quizzes", svr.URL)

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

		spew.Dump(string(body))

		var quizzes []*app.Quiz
		err = json.Unmarshal(body, &quizzes)
		if err != nil {
			t.Error(err)
		}

		if len(quizzes) == 0 {
			t.Error("expected len(quizzes) > 0")
		}
	})
}
