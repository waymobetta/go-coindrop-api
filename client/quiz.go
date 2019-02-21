// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop": quiz Resource Client
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

// ShowQuizPath computes a request path to the show action of quiz.
func ShowQuizPath() string {

	return fmt.Sprintf("/v1/quiz")
}

// Get quiz
func (c *Client) ShowQuiz(ctx context.Context, path string, quizTitle *string) (*http.Response, error) {
	req, err := c.NewShowQuizRequest(ctx, path, quizTitle)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowQuizRequest create the request corresponding to the show action endpoint of the quiz resource.
func (c *Client) NewShowQuizRequest(ctx context.Context, path string, quizTitle *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if quizTitle != nil {
		values.Set("quizTitle", *quizTitle)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
