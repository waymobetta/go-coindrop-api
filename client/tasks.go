// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": tasks Resource Client
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

// CreateTasksPath computes a request path to the create action of tasks.
func CreateTasksPath() string {

	return fmt.Sprintf("/v1/tasks")
}

// Create a user task
func (c *Client) CreateTasks(ctx context.Context, path string, payload *CreateTaskPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateTasksRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTasksRequest create the request corresponding to the create action endpoint of the tasks resource.
func (c *Client) NewCreateTasksRequest(ctx context.Context, path string, payload *CreateTaskPayload, contentType string) (*http.Request, error) {
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

// ListTasksPath computes a request path to the list action of tasks.
func ListTasksPath() string {

	return fmt.Sprintf("/v1/tasks")
}

// Get list of user tasks
func (c *Client) ListTasks(ctx context.Context, path string, userID *string) (*http.Response, error) {
	req, err := c.NewListTasksRequest(ctx, path, userID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTasksRequest create the request corresponding to the list action endpoint of the tasks resource.
func (c *Client) NewListTasksRequest(ctx context.Context, path string, userID *string) (*http.Request, error) {
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

// ShowTasksPath computes a request path to the show action of tasks.
func ShowTasksPath(taskID string) string {
	param0 := taskID

	return fmt.Sprintf("/v1/tasks/%s", param0)
}

// Get single task
func (c *Client) ShowTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTasksRequest create the request corresponding to the show action endpoint of the tasks resource.
func (c *Client) NewShowTasksRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateTasksPath computes a request path to the update action of tasks.
func UpdateTasksPath(taskID string) string {
	param0 := taskID

	return fmt.Sprintf("/v1/tasks/%s", param0)
}

// Update user task state
func (c *Client) UpdateTasks(ctx context.Context, path string, payload *TaskPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateTasksRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTasksRequest create the request corresponding to the update action endpoint of the tasks resource.
func (c *Client) NewUpdateTasksRequest(ctx context.Context, path string, payload *TaskPayload, contentType string) (*http.Request, error) {
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
