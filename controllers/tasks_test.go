package controllers

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShow(t *testing.T) {
	t.Run("get tasks for a authed user", func(t *testing.T) {
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
