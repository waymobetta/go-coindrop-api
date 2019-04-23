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
	// badge logo
	LogoURL string `form:"logoURL" json:"logoURL" yaml:"logoURL" xml:"logoURL"`
	// badge name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
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
	if mt.LogoURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "logoURL"))
	}
	return
}

// DecodeBadge decodes the Badge instance encoded in resp body.
func (c *Client) DecodeBadge(resp *http.Response) (*Badge, error) {
	var decoded Badge
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// BadgeCollection is the media type for an array of Badge (default view)
//
// Identifier: application/vnd.badge+json; type=collection; view=default
type BadgeCollection []*Badge

// Validate validates the BadgeCollection media type instance.
func (mt BadgeCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBadgeCollection decodes the BadgeCollection instance encoded in resp body.
func (c *Client) DecodeBadgeCollection(resp *http.Response) (BadgeCollection, error) {
	var decoded BadgeCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Badges (default view)
//
// Identifier: application/vnd.badges+json; view=default
type Badges struct {
	// list of badges
	Badges BadgeCollection `form:"badges" json:"badges" yaml:"badges" xml:"badges"`
}

// Validate validates the Badges media type instance.
func (mt *Badges) Validate() (err error) {
	if mt.Badges == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "badges"))
	}
	if err2 := mt.Badges.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeBadges decodes the Badges instance encoded in resp body.
func (c *Client) DecodeBadges(resp *http.Response) (*Badges, error) {
	var decoded Badges
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Community Object (default view)
//
// Identifier: application/vnd.community+json; view=default
type Community struct {
	// Community name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Calculated reputation/karma
	Reputation int `form:"reputation" json:"reputation" yaml:"reputation" xml:"reputation"`
}

// Validate validates the Community media type instance.
func (mt *Community) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeCommunity decodes the Community instance encoded in resp body.
func (c *Client) DecodeCommunity(resp *http.Response) (*Community, error) {
	var decoded Community
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CommunityCollection is the media type for an array of Community (default view)
//
// Identifier: application/vnd.community+json; type=collection; view=default
type CommunityCollection []*Community

// Validate validates the CommunityCollection media type instance.
func (mt CommunityCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeCommunityCollection decodes the CommunityCollection instance encoded in resp body.
func (c *Client) DecodeCommunityCollection(resp *http.Response) (CommunityCollection, error) {
	var decoded CommunityCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// ERC721 (default view)
//
// Identifier: application/vnd.erc721+json; view=default
type Erc721 struct {
	// contract address
	ContractAddress string `form:"contractAddress" json:"contractAddress" yaml:"contractAddress" xml:"contractAddress"`
	// token ID
	TokenID string `form:"tokenId" json:"tokenId" yaml:"tokenId" xml:"tokenId"`
	// total number minted
	TotalMinted int `form:"totalMinted" json:"totalMinted" yaml:"totalMinted" xml:"totalMinted"`
}

// Validate validates the Erc721 media type instance.
func (mt *Erc721) Validate() (err error) {
	if mt.TokenID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tokenId"))
	}
	if mt.ContractAddress == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "contractAddress"))
	}

	return
}

// DecodeErc721 decodes the Erc721 instance encoded in resp body.
func (c *Client) DecodeErc721(resp *http.Response) (*Erc721, error) {
	var decoded Erc721
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

// A user profile (default view)
//
// Identifier: application/vnd.profile+json; view=default
type Profile struct {
	// Unique user ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
}

// Validate validates the Profile media type instance.
func (mt *Profile) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	return
}

// DecodeProfile decodes the Profile instance encoded in resp body.
func (c *Client) DecodeProfile(resp *http.Response) (*Profile, error) {
	var decoded Profile
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Public (default view)
//
// Identifier: application/vnd.public+json; view=default
type Public struct {
	// list of badges
	Badges PublicbadgeCollection `form:"badges" json:"badges" yaml:"badges" xml:"badges"`
	// Reddit username
	RedditUsername string `form:"redditUsername" json:"redditUsername" yaml:"redditUsername" xml:"redditUsername"`
	// Stack Overflow user ID
	StackUserID int `form:"stackUserId" json:"stackUserId" yaml:"stackUserId" xml:"stackUserId"`
}

// Validate validates the Public media type instance.
func (mt *Public) Validate() (err error) {
	if mt.RedditUsername == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "redditUsername"))
	}

	if mt.Badges == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "badges"))
	}
	if err2 := mt.Badges.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodePublic decodes the Public instance encoded in resp body.
func (c *Client) DecodePublic(resp *http.Response) (*Public, error) {
	var decoded Public
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Badge (default view)
//
// Identifier: application/vnd.publicbadge+json; view=default
type Publicbadge struct {
	// badge description
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// ERC-721
	Erc721 *Erc721 `form:"erc721" json:"erc721" yaml:"erc721" xml:"erc721"`
	// badge logo
	LogoURL string `form:"logoURL" json:"logoURL" yaml:"logoURL" xml:"logoURL"`
	// badge name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// project
	Project string `form:"project" json:"project" yaml:"project" xml:"project"`
}

// Validate validates the Publicbadge media type instance.
func (mt *Publicbadge) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.LogoURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "logoURL"))
	}
	if mt.Project == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "project"))
	}
	if mt.Erc721 == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "erc721"))
	}
	if mt.Erc721 != nil {
		if err2 := mt.Erc721.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodePublicbadge decodes the Publicbadge instance encoded in resp body.
func (c *Client) DecodePublicbadge(resp *http.Response) (*Publicbadge, error) {
	var decoded Publicbadge
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// PublicbadgeCollection is the media type for an array of Publicbadge (default view)
//
// Identifier: application/vnd.publicbadge+json; type=collection; view=default
type PublicbadgeCollection []*Publicbadge

// Validate validates the PublicbadgeCollection media type instance.
func (mt PublicbadgeCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePublicbadgeCollection decodes the PublicbadgeCollection instance encoded in resp body.
func (c *Client) DecodePublicbadgeCollection(resp *http.Response) (PublicbadgeCollection, error) {
	var decoded PublicbadgeCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Quiz (default view)
//
// Identifier: application/vnd.quiz+json; view=default
type Quiz struct {
	// Quiz fields
	Fields QuizFieldsCollection `form:"fields" json:"fields" yaml:"fields" xml:"fields"`
	// Quiz ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Quiz title
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// Quiz user ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the Quiz media type instance.
func (mt *Quiz) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}
	if mt.Fields == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fields"))
	}
	if err2 := mt.Fields.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeQuiz decodes the Quiz instance encoded in resp body.
func (c *Client) DecodeQuiz(resp *http.Response) (*Quiz, error) {
	var decoded Quiz
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Quiz fields (default view)
//
// Identifier: application/vnd.quiz-fields+json; view=default
type QuizFields struct {
	// Answer
	Answer string `form:"answer" json:"answer" yaml:"answer" xml:"answer"`
	// Question
	Question string `form:"question" json:"question" yaml:"question" xml:"question"`
}

// Validate validates the QuizFields media type instance.
func (mt *QuizFields) Validate() (err error) {
	if mt.Question == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "question"))
	}
	if mt.Answer == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "answer"))
	}
	return
}

// DecodeQuizFields decodes the QuizFields instance encoded in resp body.
func (c *Client) DecodeQuizFields(resp *http.Response) (*QuizFields, error) {
	var decoded QuizFields
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Quiz-FieldsCollection is the media type for an array of Quiz-Fields (default view)
//
// Identifier: application/vnd.quiz-fields+json; type=collection; view=default
type QuizFieldsCollection []*QuizFields

// Validate validates the QuizFieldsCollection media type instance.
func (mt QuizFieldsCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeQuizFieldsCollection decodes the QuizFieldsCollection instance encoded in resp body.
func (c *Client) DecodeQuizFieldsCollection(resp *http.Response) (QuizFieldsCollection, error) {
	var decoded QuizFieldsCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// QuizCollection is the media type for an array of Quiz (default view)
//
// Identifier: application/vnd.quiz+json; type=collection; view=default
type QuizCollection []*Quiz

// Validate validates the QuizCollection media type instance.
func (mt QuizCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeQuizCollection decodes the QuizCollection instance encoded in resp body.
func (c *Client) DecodeQuizCollection(resp *http.Response) (QuizCollection, error) {
	var decoded QuizCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Reddit Targeting (default view)
//
// Identifier: application/vnd.reddittargeting+json; view=default
type Reddittargeting struct {
	// Users
	Users ReddituserCollection `form:"users" json:"users" yaml:"users" xml:"users"`
}

// Validate validates the Reddittargeting media type instance.
func (mt *Reddittargeting) Validate() (err error) {
	if mt.Users == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	if err2 := mt.Users.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeReddittargeting decodes the Reddittargeting instance encoded in resp body.
func (c *Client) DecodeReddittargeting(resp *http.Response) (*Reddittargeting, error) {
	var decoded Reddittargeting
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
	Subreddits CommunityCollection `form:"subreddits" json:"subreddits" yaml:"subreddits" xml:"subreddits"`
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
	if err2 := mt.Subreddits.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
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

// ReddituserCollection is the media type for an array of Reddituser (default view)
//
// Identifier: application/vnd.reddituser+json; type=collection; view=default
type ReddituserCollection []*Reddituser

// Validate validates the ReddituserCollection media type instance.
func (mt ReddituserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeReddituserCollection decodes the ReddituserCollection instance encoded in resp body.
func (c *Client) DecodeReddituserCollection(resp *http.Response) (ReddituserCollection, error) {
	var decoded ReddituserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Quiz results (default view)
//
// Identifier: application/vnd.results+json; view=default
type Results struct {
	// Count of correct quiz answers
	QuestionsCorrect int `form:"questionsCorrect" json:"questionsCorrect" yaml:"questionsCorrect" xml:"questionsCorrect"`
	// Count of incorrect quiz answers
	QuestionsIncorrect int `form:"questionsIncorrect" json:"questionsIncorrect" yaml:"questionsIncorrect" xml:"questionsIncorrect"`
	// Quiz ID
	QuizID string `form:"quizId" json:"quizId" yaml:"quizId" xml:"quizId"`
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the Results media type instance.
func (mt *Results) Validate() (err error) {
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}
	if mt.QuizID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "quizId"))
	}

	return
}

// DecodeResults decodes the Results instance encoded in resp body.
func (c *Client) DecodeResults(resp *http.Response) (*Results, error) {
	var decoded Results
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ResultsCollection is the media type for an array of Results (default view)
//
// Identifier: application/vnd.results+json; type=collection; view=default
type ResultsCollection []*Results

// Validate validates the ResultsCollection media type instance.
func (mt ResultsCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeResultsCollection decodes the ResultsCollection instance encoded in resp body.
func (c *Client) DecodeResultsCollection(resp *http.Response) (ResultsCollection, error) {
	var decoded ResultsCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Stack Overflow User Info (default view)
//
// Identifier: application/vnd.stackoverflowuser+json; view=default
type Stackoverflowuser struct {
	// Stack Exchange Accounts
	Accounts CommunityCollection `form:"accounts" json:"accounts" yaml:"accounts" xml:"accounts"`
	// Display Name
	DisplayName string `form:"displayName" json:"displayName" yaml:"displayName" xml:"displayName"`
	// Stack Exchange Account ID
	ExchangeAccountID int `form:"exchangeAccountId" json:"exchangeAccountId" yaml:"exchangeAccountId" xml:"exchangeAccountId"`
	// ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Stack Overflow Community-Specific Account ID
	StackUserID int `form:"stackUserId" json:"stackUserId" yaml:"stackUserId" xml:"stackUserId"`
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
	// Social Account Verification
	Verification *Verification `form:"verification" json:"verification" yaml:"verification" xml:"verification"`
}

// Validate validates the Stackoverflowuser media type instance.
func (mt *Stackoverflowuser) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}

	if mt.DisplayName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "displayName"))
	}
	if mt.Accounts == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "accounts"))
	}
	if mt.Verification == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "verification"))
	}
	if err2 := mt.Accounts.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
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

// DecodeStackoverflowuser decodes the Stackoverflowuser instance encoded in resp body.
func (c *Client) DecodeStackoverflowuser(resp *http.Response) (*Stackoverflowuser, error) {
	var decoded Stackoverflowuser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Targeting (default view)
//
// Identifier: application/vnd.targeting+json; view=default
type Targeting struct {
	// List of users
	Users string `form:"users" json:"users" yaml:"users" xml:"users"`
}

// Validate validates the Targeting media type instance.
func (mt *Targeting) Validate() (err error) {
	if mt.Users == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	return
}

// DecodeTargeting decodes the Targeting instance encoded in resp body.
func (c *Client) DecodeTargeting(resp *http.Response) (*Targeting, error) {
	var decoded Targeting
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
	// task completed flag
	Completed bool `form:"completed" json:"completed" yaml:"completed" xml:"completed"`
	// task description
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// task ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// logo URL
	LogoURL string `form:"logoURL" json:"logoURL" yaml:"logoURL" xml:"logoURL"`
	// learning resource ID
	ResourceID string `form:"resourceId" json:"resourceId" yaml:"resourceId" xml:"resourceId"`
	// learning resource URL
	ResourceURL string `form:"resourceURL" json:"resourceURL" yaml:"resourceURL" xml:"resourceURL"`
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
	if mt.LogoURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "logoURL"))
	}
	if mt.ResourceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "resourceId"))
	}
	if mt.ResourceURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "resourceURL"))
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

// Transaction (default view)
//
// Identifier: application/vnd.transaction+json; view=default
type Transaction struct {
	// transaction hash
	Hash string `form:"hash" json:"hash" yaml:"hash" xml:"hash"`
	// transaction ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// task ID
	TaskID string `form:"taskId" json:"taskId" yaml:"taskId" xml:"taskId"`
	// user ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the Transaction media type instance.
func (mt *Transaction) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}
	if mt.TaskID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "taskId"))
	}
	if mt.Hash == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "hash"))
	}
	return
}

// DecodeTransaction decodes the Transaction instance encoded in resp body.
func (c *Client) DecodeTransaction(resp *http.Response) (*Transaction, error) {
	var decoded Transaction
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TransactionCollection is the media type for an array of Transaction (default view)
//
// Identifier: application/vnd.transaction+json; type=collection; view=default
type TransactionCollection []*Transaction

// Validate validates the TransactionCollection media type instance.
func (mt TransactionCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTransactionCollection decodes the TransactionCollection instance encoded in resp body.
func (c *Client) DecodeTransactionCollection(resp *http.Response) (TransactionCollection, error) {
	var decoded TransactionCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Transactions (default view)
//
// Identifier: application/vnd.transactions+json; view=default
type Transactions struct {
	// list of transactions
	Transactions TransactionCollection `form:"transactions" json:"transactions" yaml:"transactions" xml:"transactions"`
}

// Validate validates the Transactions media type instance.
func (mt *Transactions) Validate() (err error) {
	if mt.Transactions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "transactions"))
	}
	if err2 := mt.Transactions.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTransactions decodes the Transactions instance encoded in resp body.
func (c *Client) DecodeTransactions(resp *http.Response) (*Transactions, error) {
	var decoded Transactions
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
	Address string `form:"address" json:"address" yaml:"address" xml:"address"`
	// wallet verified flag
	Verified bool `form:"verified" json:"verified" yaml:"verified" xml:"verified"`
	// wallet type
	WalletType string `form:"walletType" json:"walletType" yaml:"walletType" xml:"walletType"`
}

// Validate validates the Wallet media type instance.
func (mt *Wallet) Validate() (err error) {
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}
	if mt.WalletType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "walletType"))
	}

	return
}

// DecodeWallet decodes the Wallet instance encoded in resp body.
func (c *Client) DecodeWallet(resp *http.Response) (*Wallet, error) {
	var decoded Wallet
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// WalletCollection is the media type for an array of Wallet (default view)
//
// Identifier: application/vnd.wallet+json; type=collection; view=default
type WalletCollection []*Wallet

// Validate validates the WalletCollection media type instance.
func (mt WalletCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeWalletCollection decodes the WalletCollection instance encoded in resp body.
func (c *Client) DecodeWalletCollection(resp *http.Response) (WalletCollection, error) {
	var decoded WalletCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Wallets (default view)
//
// Identifier: application/vnd.wallets+json; view=default
type Wallets struct {
	// list of wallets
	Wallets WalletCollection `form:"wallets" json:"wallets" yaml:"wallets" xml:"wallets"`
}

// Validate validates the Wallets media type instance.
func (mt *Wallets) Validate() (err error) {
	if mt.Wallets == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "wallets"))
	}
	if err2 := mt.Wallets.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeWallets decodes the Wallets instance encoded in resp body.
func (c *Client) DecodeWallets(resp *http.Response) (*Wallets, error) {
	var decoded Wallets
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
