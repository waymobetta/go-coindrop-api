// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": profiles Resource Client
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

// CreateProfilesPath computes a request path to the create action of profiles.
func CreateProfilesPath() string {

	return fmt.Sprintf("/v1/profiles")
}

// Upsert a new profile
func (c *Client) CreateProfiles(ctx context.Context, path string, payload *ProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateProfilesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateProfilesRequest create the request corresponding to the create action endpoint of the profiles resource.
func (c *Client) NewCreateProfilesRequest(ctx context.Context, path string, payload *ProfilePayload, contentType string) (*http.Request, error) {
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

// ListProfilesPath computes a request path to the list action of profiles.
func ListProfilesPath() string {

	return fmt.Sprintf("/v1/profiles")
}

// Get user profile
func (c *Client) ListProfiles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListProfilesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListProfilesRequest create the request corresponding to the list action endpoint of the profiles resource.
func (c *Client) NewListProfilesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowProfilesPath computes a request path to the show action of profiles.
func ShowProfilesPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/profiles/%s", param0)
}

// Get profile by user id
func (c *Client) ShowProfiles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowProfilesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowProfilesRequest create the request corresponding to the show action endpoint of the profiles resource.
func (c *Client) NewShowProfilesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateProfilesPath computes a request path to the update action of profiles.
func UpdateProfilesPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/profiles/%s", param0)
}

// Update profile
func (c *Client) UpdateProfiles(ctx context.Context, path string, payload *ProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateProfilesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateProfilesRequest create the request corresponding to the update action endpoint of the profiles resource.
func (c *Client) NewUpdateProfilesRequest(ctx context.Context, path string, payload *ProfilePayload, contentType string) (*http.Request, error) {
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
