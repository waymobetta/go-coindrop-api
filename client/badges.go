// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": badges Resource Client
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

// CreateBadgesPath computes a request path to the create action of badges.
func CreateBadgesPath() string {

	return fmt.Sprintf("/v1/badges")
}

// Create a badge
func (c *Client) CreateBadges(ctx context.Context, path string, payload *CreateBadgePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateBadgesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateBadgesRequest create the request corresponding to the create action endpoint of the badges resource.
func (c *Client) NewCreateBadgesRequest(ctx context.Context, path string, payload *CreateBadgePayload, contentType string) (*http.Request, error) {
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

// ListBadgesPath computes a request path to the list action of badges.
func ListBadgesPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/v1/badges/%s", param0)
}

// Get list of user badges
func (c *Client) ListBadges(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBadgesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBadgesRequest create the request corresponding to the list action endpoint of the badges resource.
func (c *Client) NewListBadgesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowBadgesPath computes a request path to the show action of badges.
func ShowBadgesPath() string {

	return fmt.Sprintf("/v1/badges")
}

// Get all badges
func (c *Client) ShowBadges(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBadgesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBadgesRequest create the request corresponding to the show action endpoint of the badges resource.
func (c *Client) NewShowBadgesRequest(ctx context.Context, path string) (*http.Request, error) {
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
