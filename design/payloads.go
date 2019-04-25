package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// UserPayload is the payload for creating a user
var UserPayload = Type("UserPayload", func() {
	Description("User payload")
	Attribute("cognitoAuthUserId", String, "Cognito auth user ID")
	Required("cognitoAuthUserId")
})

// CreateBadgePayload is the payload for creating a badge
var CreateBadgePayload = Type("CreateBadgePayload", func() {
	Description("Create Badge payload")
	Attribute("name", String, "Badge name")
	Attribute("description", String, "Badge description")
	Attribute("logoURL", String, "Badge logo URL")
	Required(
		"name",
		"description",
		"logoURL",
	)
})

// TODO:
// Need to add more data fields
// CreateTaskPayload is the payload for creating a task
var CreateTaskPayload = Type("CreateTaskPayload", func() {
	Description("Create Task payload")
	Attribute("userId", String, "User ID")
	Attribute("taskId", String, "Task ID")
	Required("userId", "taskId")
})

// TaskPayload is the payload for updating a task
var TaskPayload = Type("TaskPayload", func() {
	Description("Task payload")
	Attribute("completed", Boolean, "Task completed")
	Required("completed")
})

// SetTargetingPayload is for assigning a task to eligible users
var SetTargetingPayload = Type("SetTargetingPayload", func() {
	Description("Set Targeting payload")
	Attribute("taskId", String, "Task ID")
	Attribute("users", String, "List of users")
	Required(
		"taskId",
		"users",
	)
})

// AssignERC721Payload is the payload for assigning an ERC721 to a user
var AssignERC721Payload = Type("AssignERC721Payload", func() {
	Description("AssignERC721 payload")
	Attribute("badgeId", String, "Badge ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("walletAddress", String, "Wallet address", func() {
		Pattern("^0x[0-9a-fA-F]{40}$")
		Example("0x845fdD93Cca3aE9e380d5556818e6d0b902B977c")
	})
	Required(
		"badgeId",
		"userId",
		"walletAddress",
	)
})

// WalletPayload is the payload for updating a user's wallet
var WalletPayload = Type("WalletPayload", func() {
	Description("Wallet payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("walletAddress", String, "Wallet address", func() {
		Pattern("^0x[0-9a-fA-F]{40}$")
		Example("0x845fdD93Cca3aE9e380d5556818e6d0b902B977c")
	})
	Attribute("walletType", String, "wallet type")
	Required(
		"userId",
		"walletAddress",
		"walletType",
	)
})

// CreateRedditUserPayload is the payload for creating a listing for a user's reddit info
var CreateRedditUserPayload = Type("CreateUserPayload", func() {
	Description("Create Reddit User payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("username", String, "Username")
	Required(
		"userId",
		"username",
	)
})

// CreateStackOverflowUserPayload is the payload for creating a listing for a user's reddit info
var CreateStackOverflowUserPayload = Type("CreateStackOverflowUserPayload", func() {
	Description("Create Stack Overflow User payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("stackUserId", Integer, "Stack Overflow Community-Specific Account ID")
	Required(
		"userId",
		"stackUserId",
	)
})

// QuizPayload is the payload for creating a quiz
var QuizPayload = Type("QuizPayload", func() {
	Description("Quiz payload")
	Attribute("title", String, "Title")
	Required("title")
})

// QuizResultsPayload is the payload for creating a user
var QuizResultsPayload = Type("QuizResultsPayload", func() {
	Description("Quiz results payload")
	Attribute("quizId", String, "Quiz ID")
	Attribute("userId", String, "User ID")
	Attribute("questionsCorrect", Integer, "Number of questions that were answered correct")
	Attribute("questionsIncorrect", Integer, "Number of questions that were answered incorrect")
	Required("quizId", "userId", "questionsCorrect", "questionsIncorrect")
})

// VerificationPayload is the payload for updating verification data of a social account
var VerificationPayload = Type("VerificationPayload", func() {
	Description("Social Account Verification Payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Required(
		"userId",
	)
})

// UpdateRedditUserPayload is the payload for updating a user's reddit info
var UpdateRedditUserPayload = Type("UpdateRedditUserPayload", func() {
	Description("Update Reddit User payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("username", String, "Reddit Username")
	Required(
		"userId",
		"username",
	)
})

// UpdateStackOverflowUserPayload is the payload for updating a user's reddit info
var UpdateStackOverflowUserPayload = Type("UpdateStackOverflowUserPayload", func() {
	Description("Update Stack Overflow User payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("stackUserId", Integer, "Stack Overflow Community-Specific Account ID")
	Required(
		"userId",
		"stackUserId",
	)
})

// TypeformPayload is the payload for webhook.
var TypeformPayload = Type("TypeformPayload", func() {
	Description("Typeform payload")
	Attribute("event_id", String, "Event ID")
	Attribute("event_type", String, "Event types")
	Attribute("form_response", TypeformFormPayload, "Form response")
})

// TypeformFormPayload ...
var TypeformFormPayload = Type("TypeformFormPayload", func() {
	Description("Typeform form data")
	Attribute("form_id", String, "Form ID")
	Attribute("token", String, "Form ID")
	Attribute("landed_at", String, "Form ID")
	Attribute("submitted_at", String, "Form ID")
	Attribute("calculated", TypeformCalculatedPayload, "Calculated response")
	Attribute("hidden", TypeformHiddenPayload, "Hidden")
	Attribute("definition", Any, "Definition")
	Attribute("answers", Any, "Answers")
})

// TypeformCalculatedPayload ...
var TypeformCalculatedPayload = Type("TypeformCalculatedPayload", func() {
	Description("Typeform calculatd data")
	Attribute("score", Integer, "Score")
})

// TypeformHiddenPayload ...
var TypeformHiddenPayload = Type("TypeformHiddenPayload", func() {
	Description("Typeform hidden data")
	Attribute("user_id", String, "User ID")
})

// ProfilePayload is the payload for creating a user
var ProfilePayload = Type("ProfilePayload", func() {
	Description("Profile payload")
	Attribute("name", String, "Name")
	Attribute("username", String, "Username")
	Required("name", "username")
})

// ClaimPayload
var ClaimPayload = Type("ClaimPayload", func() {
	Description("Claim payload")
	Attribute("userId", String, "User ID")
	Attribute("taskId", String, "Task ID")
	Required(
		"userId",
		"taskId",
	)
})

// WalletVerificationPayload
var WalletVerificationPayload = Type("WalletVerificationPayload", func() {
	Description("Wallet verification payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("taskId", String, "Task ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("address", String, "Wallet address")
	Attribute("message", String, "Message")
	Attribute("signature", String, "Signature")
	Attribute("version", String, "Version")
	Required(
		"userId",
		"taskId",
		"address",
		"message",
		"signature",
		"version",
	)
})
