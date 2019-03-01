// Code generated by goagen v1.4.1, DO NOT EDIT.
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

// Badge (default view)
//
// Identifier: application/vnd.badge+json; view=default
type Badge struct {
	// badge description
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// badge ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// badge name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// badge recipients
	Recipients int `form:"recipients" json:"recipients" yaml:"recipients" xml:"recipients"`
}

// Validate validates the Badge media type instance.
func (mt *Badge) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}

	return
}

// Health check (default view)
//
// Identifier: application/vnd.healthcheck+json; view=default
type Healthcheck struct {
	// Status
	Status string `form:"status" json:"status" yaml:"status" xml:"status"`
}

// Validate validates the Healthcheck media type instance.
func (mt *Healthcheck) Validate() (err error) {
	if mt.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "status"))
	}
	return
}

// Quiz (default view)
//
// Identifier: application/vnd.quiz+json; view=default
type Quiz struct {
	// Quiz object
	QuizObject interface{} `form:"quizObject" json:"quizObject" yaml:"quizObject" xml:"quizObject"`
}

// Quiz results (default view)
//
// Identifier: application/vnd.results+json; view=default
type Results struct {
	// quiz results list
	QuizResultsList interface{} `form:"quizResultsList" json:"quizResultsList" yaml:"quizResultsList" xml:"quizResultsList"`
}

// Task (default view)
//
// Identifier: application/vnd.task+json; view=default
type Task struct {
	// task author
	Author string `form:"author" json:"author" yaml:"author" xml:"author"`
	// task badge
	Badge *Badge `form:"badge" json:"badge" yaml:"badge" xml:"badge"`
	// task description
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// task ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// task title
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// task token
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
	// token allocation
	TokenAllocation int `form:"tokenAllocation" json:"tokenAllocation" yaml:"tokenAllocation" xml:"tokenAllocation"`
	// task type
	Type string `form:"type" json:"type" yaml:"type" xml:"type"`
}

// Validate validates the Task media type instance.
func (mt *Task) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.Author == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "author"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	if mt.Badge == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "badge"))
	}
	if mt.Badge != nil {
		if err2 := mt.Badge.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// TaskCollection is the media type for an array of Task (default view)
//
// Identifier: application/vnd.task+json; type=collection; view=default
type TaskCollection []*Task

// Validate validates the TaskCollection media type instance.
func (mt TaskCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Tasks (default view)
//
// Identifier: application/vnd.tasks+json; view=default
type Tasks struct {
	// list of tasks
	Tasks TaskCollection `form:"tasks" json:"tasks" yaml:"tasks" xml:"tasks"`
}

// Validate validates the Tasks media type instance.
func (mt *Tasks) Validate() (err error) {
	if mt.Tasks == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tasks"))
	}
	if err2 := mt.Tasks.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
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
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Wallet address
	WalletAddress *string `form:"walletAddress,omitempty" json:"walletAddress,omitempty" yaml:"walletAddress,omitempty" xml:"walletAddress,omitempty"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
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
