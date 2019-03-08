// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop": users Resource Client
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

// CreateUsersPath computes a request path to the create action of users.
func CreateUsersPath() string {

	return fmt.Sprintf("/v1/users")
}

// Create a new user
func (c *Client) CreateUsers(ctx context.Context, path string, payload *UserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUsersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUsersRequest create the request corresponding to the create action endpoint of the users resource.
func (c *Client) NewCreateUsersRequest(ctx context.Context, path string, payload *UserPayload, contentType string) (*http.Request, error) {
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
	return req, nil
}

// ListUsersPath computes a request path to the list action of users.
func ListUsersPath() string {

	return fmt.Sprintf("/v1/users")
}

// Get user ID mapped to Cognito auth user ID
func (c *Client) ListUsers(ctx context.Context, path string, cognitoAuthUserID *string) (*http.Response, error) {
	req, err := c.NewListUsersRequest(ctx, path, cognitoAuthUserID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUsersRequest create the request corresponding to the list action endpoint of the users resource.
func (c *Client) NewListUsersRequest(ctx context.Context, path string, cognitoAuthUserID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if cognitoAuthUserID != nil {
		values.Set("cognitoAuthUserId", *cognitoAuthUserID)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowUsersPath computes a request path to the show action of users.
func ShowUsersPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/users/%s", param0)
}

// Get user by id
func (c *Client) ShowUsers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUsersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUsersRequest create the request corresponding to the show action endpoint of the users resource.
func (c *Client) NewShowUsersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
