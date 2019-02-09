// Code generated by goagen v1.3.1, DO NOT EDIT.
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
