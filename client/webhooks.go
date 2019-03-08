// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop": webhooks Resource Client
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

// TypeformWebhooksPath computes a request path to the typeform action of webhooks.
func TypeformWebhooksPath() string {

	return fmt.Sprintf("/v1/webhooks/typeform")
}

// Typeform webhook
func (c *Client) TypeformWebhooks(ctx context.Context, path string, payload *TypeformPayload, contentType string) (*http.Response, error) {
	req, err := c.NewTypeformWebhooksRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTypeformWebhooksRequest create the request corresponding to the typeform action endpoint of the webhooks resource.
func (c *Client) NewTypeformWebhooksRequest(ctx context.Context, path string, payload *TypeformPayload, contentType string) (*http.Request, error) {
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
