// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop": stackoverflowharvest Resource Client
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

// UpdateStackoverflowharvestPath computes a request path to the update action of stackoverflowharvest.
func UpdateStackoverflowharvestPath() string {

	return fmt.Sprintf("/v1/social/stackoverflow/harvest")
}

// Update Stack Overflow User Info
func (c *Client) UpdateStackoverflowharvest(ctx context.Context, path string, payload *UpdateStackOverflowUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateStackoverflowharvestRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateStackoverflowharvestRequest create the request corresponding to the update action endpoint of the stackoverflowharvest resource.
func (c *Client) NewUpdateStackoverflowharvestRequest(ctx context.Context, path string, payload *UpdateStackOverflowUserPayload, contentType string) (*http.Request, error) {
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
