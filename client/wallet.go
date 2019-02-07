// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": wallet Resource Client
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowWalletPath computes a request path to the show action of wallet.
func ShowWalletPath() string {

	return fmt.Sprintf("/v1/wallets")
}

// Get user wallet
func (c *Client) ShowWallet(ctx context.Context, path string, userID *string) (*http.Response, error) {
	req, err := c.NewShowWalletRequest(ctx, path, userID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowWalletRequest create the request corresponding to the show action endpoint of the wallet resource.
func (c *Client) NewShowWalletRequest(ctx context.Context, path string, userID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if userID != nil {
		values.Set("userID", *userID)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
