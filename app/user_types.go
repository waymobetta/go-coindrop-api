// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "coindrop": Application User Types
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-api
// --version=v1.4.1

package app

import (
	"github.com/goadesign/goa"
)

// Create Task payload
type createTaskPayload struct {
	// Task ID
	TaskID *string `form:"taskId,omitempty" json:"taskId,omitempty" yaml:"taskId,omitempty" xml:"taskId,omitempty"`
	// User ID
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the createTaskPayload type instance.
func (ut *createTaskPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	if ut.TaskID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "taskId"))
	}
	return
}

// Publicize creates CreateTaskPayload from createTaskPayload
func (ut *createTaskPayload) Publicize() *CreateTaskPayload {
	var pub CreateTaskPayload
	if ut.TaskID != nil {
		pub.TaskID = *ut.TaskID
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// Create Task payload
type CreateTaskPayload struct {
	// Task ID
	TaskID string `form:"taskId" json:"taskId" yaml:"taskId" xml:"taskId"`
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the CreateTaskPayload type instance.
func (ut *CreateTaskPayload) Validate() (err error) {
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "userId"))
	}
	if ut.TaskID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "taskId"))
	}
	return
}

// Create Reddit User payload
type createUserPayload struct {
	// User ID
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the createUserPayload type instance.
func (ut *createUserPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "username"))
	}
	if ut.UserID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, *ut.UserID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.userId`, *ut.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
		}
	}
	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (ut *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// Create Reddit User payload
type CreateUserPayload struct {
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
	// Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
}

// Validate validates the CreateUserPayload type instance.
func (ut *CreateUserPayload) Validate() (err error) {
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "userId"))
	}
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "username"))
	}
	if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, ut.UserID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.userId`, ut.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
	}
	return
}

// Quiz payload
type quizPayload struct {
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the quizPayload type instance.
func (ut *quizPayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	return
}

// Publicize creates QuizPayload from quizPayload
func (ut *quizPayload) Publicize() *QuizPayload {
	var pub QuizPayload
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	return &pub
}

// Quiz payload
type QuizPayload struct {
	// Title
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// Validate validates the QuizPayload type instance.
func (ut *QuizPayload) Validate() (err error) {
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	return
}

// Quiz results payload
type quizResultsPayload struct {
	// Number of questions that were answered correct
	QuestionsCorrect *int `form:"questionsCorrect,omitempty" json:"questionsCorrect,omitempty" yaml:"questionsCorrect,omitempty" xml:"questionsCorrect,omitempty"`
	// Number of questions that were answered incorrect
	QuestionsIncorrect *int `form:"questionsIncorrect,omitempty" json:"questionsIncorrect,omitempty" yaml:"questionsIncorrect,omitempty" xml:"questionsIncorrect,omitempty"`
	// Quiz ID
	QuizID *string `form:"quizId,omitempty" json:"quizId,omitempty" yaml:"quizId,omitempty" xml:"quizId,omitempty"`
	// User ID
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the quizResultsPayload type instance.
func (ut *quizResultsPayload) Validate() (err error) {
	if ut.QuizID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "quizId"))
	}
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	if ut.QuestionsCorrect == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "questionsCorrect"))
	}
	if ut.QuestionsIncorrect == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "questionsIncorrect"))
	}
	return
}

// Publicize creates QuizResultsPayload from quizResultsPayload
func (ut *quizResultsPayload) Publicize() *QuizResultsPayload {
	var pub QuizResultsPayload
	if ut.QuestionsCorrect != nil {
		pub.QuestionsCorrect = *ut.QuestionsCorrect
	}
	if ut.QuestionsIncorrect != nil {
		pub.QuestionsIncorrect = *ut.QuestionsIncorrect
	}
	if ut.QuizID != nil {
		pub.QuizID = *ut.QuizID
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// Quiz results payload
type QuizResultsPayload struct {
	// Number of questions that were answered correct
	QuestionsCorrect int `form:"questionsCorrect" json:"questionsCorrect" yaml:"questionsCorrect" xml:"questionsCorrect"`
	// Number of questions that were answered incorrect
	QuestionsIncorrect int `form:"questionsIncorrect" json:"questionsIncorrect" yaml:"questionsIncorrect" xml:"questionsIncorrect"`
	// Quiz ID
	QuizID string `form:"quizId" json:"quizId" yaml:"quizId" xml:"quizId"`
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the QuizResultsPayload type instance.
func (ut *QuizResultsPayload) Validate() (err error) {
	if ut.QuizID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "quizId"))
	}
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "userId"))
	}

	return
}

// Task payload
type taskPayload struct {
	// Task completed
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" yaml:"completed,omitempty" xml:"completed,omitempty"`
}

// Validate validates the taskPayload type instance.
func (ut *taskPayload) Validate() (err error) {
	if ut.Completed == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "completed"))
	}
	return
}

// Publicize creates TaskPayload from taskPayload
func (ut *taskPayload) Publicize() *TaskPayload {
	var pub TaskPayload
	if ut.Completed != nil {
		pub.Completed = *ut.Completed
	}
	return &pub
}

// Task payload
type TaskPayload struct {
	// Task completed
	Completed bool `form:"completed" json:"completed" yaml:"completed" xml:"completed"`
}

// Update Reddit User payload
type updateUserPayload struct {
	// User ID
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
	// Reddit Username
	Username *string `form:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the updateUserPayload type instance.
func (ut *updateUserPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "username"))
	}
	if ut.UserID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, *ut.UserID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.userId`, *ut.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
		}
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (ut *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// Update Reddit User payload
type UpdateUserPayload struct {
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
	// Reddit Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
}

// Validate validates the UpdateUserPayload type instance.
func (ut *UpdateUserPayload) Validate() (err error) {
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "userId"))
	}
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "username"))
	}
	if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, ut.UserID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.userId`, ut.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
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

// Social Account Verification Payload
type verificationPayload struct {
	// Verification Code Posted In Social Forum
	ConfirmedVerificationCode *string `form:"confirmedVerificationCode,omitempty" json:"confirmedVerificationCode,omitempty" yaml:"confirmedVerificationCode,omitempty" xml:"confirmedVerificationCode,omitempty"`
	// User ID
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the verificationPayload type instance.
func (ut *verificationPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	if ut.ConfirmedVerificationCode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "confirmedVerificationCode"))
	}
	if ut.UserID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, *ut.UserID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.userId`, *ut.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
		}
	}
	return
}

// Publicize creates VerificationPayload from verificationPayload
func (ut *verificationPayload) Publicize() *VerificationPayload {
	var pub VerificationPayload
	if ut.ConfirmedVerificationCode != nil {
		pub.ConfirmedVerificationCode = *ut.ConfirmedVerificationCode
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// Social Account Verification Payload
type VerificationPayload struct {
	// Verification Code Posted In Social Forum
	ConfirmedVerificationCode string `form:"confirmedVerificationCode" json:"confirmedVerificationCode" yaml:"confirmedVerificationCode" xml:"confirmedVerificationCode"`
	// User ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the VerificationPayload type instance.
func (ut *VerificationPayload) Validate() (err error) {
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "userId"))
	}
	if ut.ConfirmedVerificationCode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "confirmedVerificationCode"))
	}
	if ok := goa.ValidatePattern(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`, ut.UserID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.userId`, ut.UserID, `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`))
	}
	return
}

// Wallet payload
type walletPayload struct {
	// Wallet address
	WalletAddress *string `form:"walletAddress,omitempty" json:"walletAddress,omitempty" yaml:"walletAddress,omitempty" xml:"walletAddress,omitempty"`
}

// Validate validates the walletPayload type instance.
func (ut *walletPayload) Validate() (err error) {
	if ut.WalletAddress == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "walletAddress"))
	}
	if ut.WalletAddress != nil {
		if ok := goa.ValidatePattern(`^0x[0-9a-fA-F]{40}$`, *ut.WalletAddress); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.walletAddress`, *ut.WalletAddress, `^0x[0-9a-fA-F]{40}$`))
		}
	}
	return
}

// Publicize creates WalletPayload from walletPayload
func (ut *walletPayload) Publicize() *WalletPayload {
	var pub WalletPayload
	if ut.WalletAddress != nil {
		pub.WalletAddress = *ut.WalletAddress
	}
	return &pub
}

// Wallet payload
type WalletPayload struct {
	// Wallet address
	WalletAddress string `form:"walletAddress" json:"walletAddress" yaml:"walletAddress" xml:"walletAddress"`
}

// Validate validates the WalletPayload type instance.
func (ut *WalletPayload) Validate() (err error) {
	if ut.WalletAddress == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "walletAddress"))
	}
	if ok := goa.ValidatePattern(`^0x[0-9a-fA-F]{40}$`, ut.WalletAddress); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.walletAddress`, ut.WalletAddress, `^0x[0-9a-fA-F]{40}$`))
	}
	return
}
