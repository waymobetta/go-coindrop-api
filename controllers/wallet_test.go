package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestWalletShow(t *testing.T) {
	t.Run("get wallet for an authed user", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		userID := getUserID()
		url := fmt.Sprintf("%s/v1/wallets?userId=%s", svr.URL, userID)

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
	})
}

func TestWalletUpdate(t *testing.T) {
	t.Run("update wallet for authed user", func(t *testing.T) {
		svr := createServer()
		defer svr.Close()

		newWalletAddress := "0x4bb283712e6793D99fD4E30370C14443D834BeBf"
		url := fmt.Sprintf("%s/v1/wallets", svr.URL)

		t.Logf("URL: %s", url)

		payload := []byte(fmt.Sprintf(`{"walletAddress":%q}`, newWalletAddress))

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
