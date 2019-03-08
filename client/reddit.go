// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop": reddit Resource Client
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DisplayRedditPath computes a request path to the display action of reddit.
func DisplayRedditPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/social/reddit/%s/verify", param0)
}

// Get Reddit Verification
func (c *Client) DisplayReddit(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDisplayRedditRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDisplayRedditRequest create the request corresponding to the display action endpoint of the reddit resource.
func (c *Client) NewDisplayRedditRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
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

// ShowRedditPath computes a request path to the show action of reddit.
func ShowRedditPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/social/reddit/%s", param0)
}

// Get Reddit User
func (c *Client) ShowReddit(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowRedditRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowRedditRequest create the request corresponding to the show action endpoint of the reddit resource.
func (c *Client) NewShowRedditRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
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

// UpdateRedditPath computes a request path to the update action of reddit.
func UpdateRedditPath() string {

	return fmt.Sprintf("/v1/social/reddit")
}

// Update Reddit User
func (c *Client) UpdateReddit(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateRedditRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateRedditRequest create the request corresponding to the update action endpoint of the reddit resource.
func (c *Client) NewUpdateRedditRequest(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Request, error) {
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

// VerifyRedditPath computes a request path to the verify action of reddit.
func VerifyRedditPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/social/reddit/%s/verify", param0)
}

// Update Reddit Verification
func (c *Client) VerifyReddit(ctx context.Context, path string, payload *VerificationPayload, contentType string) (*http.Response, error) {
	req, err := c.NewVerifyRedditRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewVerifyRedditRequest create the request corresponding to the verify action endpoint of the reddit resource.
func (c *Client) NewVerifyRedditRequest(ctx context.Context, path string, payload *VerificationPayload, contentType string) (*http.Request, error) {
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
