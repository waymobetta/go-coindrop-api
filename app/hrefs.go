// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// HealthcheckHref returns the resource href.
func HealthcheckHref() string {
	return "/v1/health"
}

// QuizHref returns the resource href.
func QuizHref() string {
	return "/v1/quiz"
}

// RedditHref returns the resource href.
func RedditHref() string {
	return "/v1/social/reddit"
}

// ResultsHref returns the resource href.
func ResultsHref() string {
	return "/v1/quiz/results"
}

// TasksHref returns the resource href.
func TasksHref() string {
	return "/v1/tasks"
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/users/%v", paramuserID)
}

// WalletHref returns the resource href.
func WalletHref() string {
	return "/v1/wallets"
}
