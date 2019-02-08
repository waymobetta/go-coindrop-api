// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// A standard error response (default view)
//
// Identifier: application/standard_error+json; view=default
type StandardError struct {
	// A code that describes the error
	Code int `form:"code" json:"code" yaml:"code" xml:"code"`
	// A message that describes the error
	Message string `form:"message" json:"message" yaml:"message" xml:"message"`
}

// Validate validates the StandardError media type instance.
func (mt *StandardError) Validate() (err error) {

	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// A task (default view)
//
// Identifier: application/vnd.task+json; view=default
type Task struct {
	// task name
	TaskName string `form:"taskName" json:"taskName" yaml:"taskName" xml:"taskName"`
}

// Validate validates the Task media type instance.
func (mt *Task) Validate() (err error) {
	if mt.TaskName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "taskName"))
	}
	return
}

// A user (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// Cognito auth user ID
	CognitoAuthUserID *string `form:"cognitoAuthUserId,omitempty" json:"cognitoAuthUserId,omitempty" yaml:"cognitoAuthUserId,omitempty" xml:"cognitoAuthUserId,omitempty"`
	// Unique user ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Wallet address
	WalletAddress *string `form:"walletAddress,omitempty" json:"walletAddress,omitempty" yaml:"walletAddress,omitempty" xml:"walletAddress,omitempty"`
}

// A wallet (default view)
//
// Identifier: application/vnd.wallet+json; view=default
type Wallet struct {
	// wallet address
	WalletAddress string `form:"walletAddress" json:"walletAddress" yaml:"walletAddress" xml:"walletAddress"`
}

// Validate validates the Wallet media type instance.
func (mt *Wallet) Validate() (err error) {
	if mt.WalletAddress == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "walletAddress"))
	}
	return
}
