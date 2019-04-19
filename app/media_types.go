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
