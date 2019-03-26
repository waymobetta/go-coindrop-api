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

// BadgesHref returns the resource href.
func BadgesHref() string {
	return "/v1/badges"
}

// HealthcheckHref returns the resource href.
func HealthcheckHref() string {
	return "/v1/health"
}

// ProfilesHref returns the resource href.
func ProfilesHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/profiles/%v", paramuserID)
}

// QuizzesHref returns the resource href.
func QuizzesHref(quizID interface{}) string {
	paramquizID := strings.TrimLeftFunc(fmt.Sprintf("%v", quizID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/quizzes/%v", paramquizID)
}

// RedditHref returns the resource href.
func RedditHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/social/reddit/%v", paramuserID)
}

// ResultsHref returns the resource href.
func ResultsHref(quizID interface{}) string {
	paramquizID := strings.TrimLeftFunc(fmt.Sprintf("%v", quizID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/quizzes/%v/results", paramquizID)
}

// StackoverflowHref returns the resource href.
func StackoverflowHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/social/stackoverflow/%v", paramuserID)
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

// WalletsHref returns the resource href.
func WalletsHref() string {
	return "/v1/wallets"
}
