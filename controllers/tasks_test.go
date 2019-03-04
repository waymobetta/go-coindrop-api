package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestTaskCreate(t *testing.T) {
	t.Run("create task", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		taskID := "6bc25651-c46d-448b-a88e-ff2e2ed3b54c"
		userID := getUserID()
		url := fmt.Sprintf("%s/v1/tasks", svr.URL)

		t.Logf("URL: %s", url)

		client := &http.Client{}
		payload := []byte(fmt.Sprintf(`{"taskId": %q, "userId": %q}`, taskID, userID))

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
			t.Fatalf("got %d; want %d\n", resp.StatusCode, http.StatusOK)
		}
	})
}

func TestTaskUpdate(t *testing.T) {
	t.Run("update task", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		taskID := "6bc25651-c46d-448b-a88e-ff2e2ed3b54c"
		completed := true
		url := fmt.Sprintf("%s/v1/tasks/%s", svr.URL, taskID)

		t.Logf("URL: %s", url)

		client := &http.Client{}
		payload := []byte(fmt.Sprintf(`{"completed": %v}`, completed))

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

func TestTaskShow(t *testing.T) {
	t.Run("get tasks for an authed user", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		userID := getUserID()
		url := fmt.Sprintf("%s/v1/tasks?userId=%s", svr.URL, userID)

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
	})
}
