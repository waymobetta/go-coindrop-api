// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
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

// DecodeStandardError decodes the StandardError instance encoded in resp body.
func (c *Client) DecodeStandardError(resp *http.Response) (*StandardError, error) {
	var decoded StandardError
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeBadge decodes the Badge instance encoded in resp body.
func (c *Client) DecodeBadge(resp *http.Response) (*Badge, error) {
	var decoded Badge
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeHealthcheck decodes the Healthcheck instance encoded in resp body.
func (c *Client) DecodeHealthcheck(resp *http.Response) (*Healthcheck, error) {
	var decoded Healthcheck
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Quiz (default view)
//
// Identifier: application/vnd.quiz+json; view=default
type Quiz struct {
	// Quiz object
	QuizObject interface{} `form:"quizObject" json:"quizObject" yaml:"quizObject" xml:"quizObject"`
}

// DecodeQuiz decodes the Quiz instance encoded in resp body.
func (c *Client) DecodeQuiz(resp *http.Response) (*Quiz, error) {
	var decoded Quiz
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A Reddit User (default view)
//
// Identifier: application/vnd.reddituser+json; view=default
type Reddituser struct {
	// Comment Karma
	CommentKarma int `form:"commentKarma" json:"commentKarma" yaml:"commentKarma" xml:"commentKarma"`
	// ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Link Karma
	LinkKarma int `form:"linkKarma" json:"linkKarma" yaml:"linkKarma" xml:"linkKarma"`
	// User subreddits
	Subreddits []string `form:"subreddits" json:"subreddits" yaml:"subreddits" xml:"subreddits"`
	// User trophies
	Trophies []string `form:"trophies" json:"trophies" yaml:"trophies" xml:"trophies"`
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
	// Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
	// Social Account Verification
	Verification *Verification `form:"verification" json:"verification" yaml:"verification" xml:"verification"`
}

// Validate validates the Reddituser media type instance.
func (mt *Reddituser) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}

	if mt.Trophies == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "trophies"))
	}
	if mt.Subreddits == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subreddits"))
	}
	if mt.Verification == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "verification"))
	}
	if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
	}
	if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, mt.UserID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.userId`, mt.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
	}
	if mt.Verification != nil {
		if err2 := mt.Verification.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeReddituser decodes the Reddituser instance encoded in resp body.
func (c *Client) DecodeReddituser(resp *http.Response) (*Reddituser, error) {
	var decoded Reddituser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Quiz results (default view)
//
// Identifier: application/vnd.results+json; view=default
type Results struct {
	// quiz results list
	QuizResultsList interface{} `form:"quizResultsList" json:"quizResultsList" yaml:"quizResultsList" xml:"quizResultsList"`
}

// DecodeResults decodes the Results instance encoded in resp body.
func (c *Client) DecodeResults(resp *http.Response) (*Results, error) {
	var decoded Results
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeTask decodes the Task instance encoded in resp body.
func (c *Client) DecodeTask(resp *http.Response) (*Task, error) {
	var decoded Task
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeTaskCollection decodes the TaskCollection instance encoded in resp body.
func (c *Client) DecodeTaskCollection(resp *http.Response) (TaskCollection, error) {
	var decoded TaskCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
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

// DecodeTasks decodes the Tasks instance encoded in resp body.
func (c *Client) DecodeTasks(resp *http.Response) (*Tasks, error) {
	var decoded Tasks
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Account Verification (default view)
//
// Identifier: application/vnd.verification+json; view=default
type Verification struct {
	// Confirmed Verification Code
	ConfirmedVerificationCode string `form:"confirmedVerificationCode" json:"confirmedVerificationCode" yaml:"confirmedVerificationCode" xml:"confirmedVerificationCode"`
	// Posted Verification Code
	PostedVerificationCode string `form:"postedVerificationCode" json:"postedVerificationCode" yaml:"postedVerificationCode" xml:"postedVerificationCode"`
	// Account Verified Flag
	Verified bool `form:"verified" json:"verified" yaml:"verified" xml:"verified"`
}

// Validate validates the Verification media type instance.
func (mt *Verification) Validate() (err error) {
	if mt.PostedVerificationCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "postedVerificationCode"))
	}
	if mt.ConfirmedVerificationCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "confirmedVerificationCode"))
	}

	return
}

// DecodeVerification decodes the Verification instance encoded in resp body.
func (c *Client) DecodeVerification(resp *http.Response) (*Verification, error) {
	var decoded Verification
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeWallet decodes the Wallet instance encoded in resp body.
func (c *Client) DecodeWallet(resp *http.Response) (*Wallet, error) {
	var decoded Wallet
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
