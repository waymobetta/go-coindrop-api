// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application User Types
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
)

// User payload
type userPayload struct {
	// Cognito auth user ID
	CognitoAuthUserID *string `form:"cognitoAuthUserId,omitempty" json:"cognitoAuthUserId,omitempty" yaml:"cognitoAuthUserId,omitempty" xml:"cognitoAuthUserId,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.CognitoAuthUserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "cognitoAuthUserId"))
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.CognitoAuthUserID != nil {
		pub.CognitoAuthUserID = *ut.CognitoAuthUserID
	}
	return &pub
}

// User payload
type UserPayload struct {
	// Cognito auth user ID
	CognitoAuthUserID string `form:"cognitoAuthUserId" json:"cognitoAuthUserId" yaml:"cognitoAuthUserId" xml:"cognitoAuthUserId"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.CognitoAuthUserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "cognitoAuthUserId"))
	}
	return
}
