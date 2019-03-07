package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestWebhooksTypeform(t *testing.T) {
	t.Run("typeorm webhook post", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		url := fmt.Sprintf("%s/v1/webhooks/typeform", svr.URL)

		t.Logf("URL: %s", url)

		payload, err := ioutil.ReadFile(".webhook_sample_response")
		if err != nil {
			t.Error(err)
		}

		client := &http.Client{}
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

		spew.Dump(string(body))
	})
}
