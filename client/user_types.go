// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application User Types
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.4.1

package client

import (
	"github.com/goadesign/goa"
)

// Task payload
type taskPayload struct {
	// Cognito auth user ID
	CognitoAuthUserID *string `form:"cognitoAuthUserId,omitempty" json:"cognitoAuthUserId,omitempty" yaml:"cognitoAuthUserId,omitempty" xml:"cognitoAuthUserId,omitempty"`
	// task name
	TaskName *string `form:"taskName,omitempty" json:"taskName,omitempty" yaml:"taskName,omitempty" xml:"taskName,omitempty"`
	// task state
	TaskState *string `form:"taskState,omitempty" json:"taskState,omitempty" yaml:"taskState,omitempty" xml:"taskState,omitempty"`
}

// Validate validates the taskPayload type instance.
func (ut *taskPayload) Validate() (err error) {
	if ut.CognitoAuthUserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "cognitoAuthUserId"))
	}
	if ut.TaskName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "taskName"))
	}
	if ut.TaskState == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "taskState"))
	}
	return
}

// Publicize creates TaskPayload from taskPayload
func (ut *taskPayload) Publicize() *TaskPayload {
	var pub TaskPayload
	if ut.CognitoAuthUserID != nil {
		pub.CognitoAuthUserID = *ut.CognitoAuthUserID
	}
	if ut.TaskName != nil {
		pub.TaskName = *ut.TaskName
	}
	if ut.TaskState != nil {
		pub.TaskState = *ut.TaskState
	}
	return &pub
}

// Task payload
type TaskPayload struct {
	// Cognito auth user ID
	CognitoAuthUserID string `form:"cognitoAuthUserId" json:"cognitoAuthUserId" yaml:"cognitoAuthUserId" xml:"cognitoAuthUserId"`
	// task name
	TaskName string `form:"taskName" json:"taskName" yaml:"taskName" xml:"taskName"`
	// task state
	TaskState string `form:"taskState" json:"taskState" yaml:"taskState" xml:"taskState"`
}

// Validate validates the TaskPayload type instance.
func (ut *TaskPayload) Validate() (err error) {
	if ut.CognitoAuthUserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "cognitoAuthUserId"))
	}
	if ut.TaskName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "taskName"))
	}
	if ut.TaskState == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "taskState"))
	}
	return
}

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

// Wallet payload
type walletPayload struct {
	// Cognito auth user ID
	CognitoAuthUserID *string `form:"cognitoAuthUserId,omitempty" json:"cognitoAuthUserId,omitempty" yaml:"cognitoAuthUserId,omitempty" xml:"cognitoAuthUserId,omitempty"`
	// Wallet address
	WalletAddress *string `form:"walletAddress,omitempty" json:"walletAddress,omitempty" yaml:"walletAddress,omitempty" xml:"walletAddress,omitempty"`
}

// Validate validates the walletPayload type instance.
func (ut *walletPayload) Validate() (err error) {
	if ut.CognitoAuthUserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "cognitoAuthUserId"))
	}
	if ut.WalletAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "walletAddress"))
	}
	return
}

// Publicize creates WalletPayload from walletPayload
func (ut *walletPayload) Publicize() *WalletPayload {
	var pub WalletPayload
	if ut.CognitoAuthUserID != nil {
		pub.CognitoAuthUserID = *ut.CognitoAuthUserID
	}
	if ut.WalletAddress != nil {
		pub.WalletAddress = *ut.WalletAddress
	}
	return &pub
}

// Wallet payload
type WalletPayload struct {
	// Cognito auth user ID
	CognitoAuthUserID string `form:"cognitoAuthUserId" json:"cognitoAuthUserId" yaml:"cognitoAuthUserId" xml:"cognitoAuthUserId"`
	// Wallet address
	WalletAddress string `form:"walletAddress" json:"walletAddress" yaml:"walletAddress" xml:"walletAddress"`
}

// Validate validates the WalletPayload type instance.
func (ut *WalletPayload) Validate() (err error) {
	if ut.CognitoAuthUserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "cognitoAuthUserId"))
	}
	if ut.WalletAddress == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "walletAddress"))
	}
	return
}
