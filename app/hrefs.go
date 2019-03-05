// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.4.1

package app

import (
	"fmt"
	"strings"
)

// HealthcheckHref returns the resource href.
func HealthcheckHref() string {
	return "/v1/health"
}

// QuizzesHref returns the resource href.
func QuizzesHref(quizID interface{}) string {
	paramquizID := strings.TrimLeftFunc(fmt.Sprintf("%v", quizID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/quizzes/%v", paramquizID)
}

// RedditHref returns the resource href.
func RedditHref() string {
	return "/v1/social/reddit/userid"
}

// ResultsHref returns the resource href.
func ResultsHref() string {
	return "/v1/quiz/results"
}

// TasksHref returns the resource href.
func TasksHref(taskID interface{}) string {
	paramtaskID := strings.TrimLeftFunc(fmt.Sprintf("%v", taskID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/tasks/%v", paramtaskID)
}

// UsersHref returns the resource href.
func UsersHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/users/%v", paramuserID)
}

// VerifyredditHref returns the resource href.
func VerifyredditHref() string {
	return "/v1/social/reddit/userid/verify"
}

// WalletsHref returns the resource href.
func WalletsHref() string {
	return "/v1/wallets"
}
