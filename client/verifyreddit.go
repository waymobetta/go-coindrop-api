// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": verifyreddit Resource Client
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.4.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowVerifyredditPath computes a request path to the show action of verifyreddit.
func ShowVerifyredditPath() string {

	return fmt.Sprintf("/v1/social/reddit/userid/verify")
}

// Get
func (c *Client) ShowVerifyreddit(ctx context.Context, path string, userID *string) (*http.Response, error) {
	req, err := c.NewShowVerifyredditRequest(ctx, path, userID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowVerifyredditRequest create the request corresponding to the show action endpoint of the verifyreddit resource.
func (c *Client) NewShowVerifyredditRequest(ctx context.Context, path string, userID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if userID != nil {
		values.Set("userId", *userID)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTAuthSigner != nil {
		if err := c.JWTAuthSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateVerifyredditPath computes a request path to the update action of verifyreddit.
func UpdateVerifyredditPath() string {

	return fmt.Sprintf("/v1/social/reddit/userid/verify")
}

// Update Reddit Verification Code
func (c *Client) UpdateVerifyreddit(ctx context.Context, path string, payload *VerificationPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateVerifyredditRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateVerifyredditRequest create the request corresponding to the update action endpoint of the verifyreddit resource.
func (c *Client) NewUpdateVerifyredditRequest(ctx context.Context, path string, payload *VerificationPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTAuthSigner != nil {
		if err := c.JWTAuthSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
