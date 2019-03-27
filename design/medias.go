package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// UserMedia defines the media type used to render users.
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", String, "Unique user ID")
		Attribute("cognitoAuthUserId", String, "Cognito auth user ID")
		Attribute("name", String, "Name of user")
		Attribute("walletAddress", String, "Wallet address")
		Required("id")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must
		Attribute("cognitoAuthUserId")
		Attribute("name")
		Attribute("walletAddress")
	})
})

// BadgeMedia ...
var BadgeMedia = MediaType("application/vnd.badge+json", func() {
	Description("Badge")
	Attributes(func() {
		Attribute("id", String, "badge ID")
		Attribute("name", String, "badge name")
		Attribute("description", String, "badge description")
		Attribute("logoURL", String, "badge logo")
		Required(
			"id",
			"name",
			"description",
			"logoURL",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("logoURL")
	})
})

// BadgesMedia ...
var BadgesMedia = MediaType("application/vnd.badges+json", func() {
	Description("Badges")
	Attributes(func() {
		Attribute("badges", CollectionOf(BadgeMedia), "list of badges")
		Required("badges")
	})
	View("default", func() {
		Attribute("badges")
	})
})

// TaskMedia ...
var TaskMedia = MediaType("application/vnd.task+json", func() {
	Description("Task")
	Attributes(func() {
		Attribute("id", String, "task ID")
		Attribute("title", String, "task title")
		Attribute("type", String, "task type")
		Attribute("author", String, "task author")
		Attribute("description", String, "task description")
		Attribute("token", String, "task token")
		Attribute("tokenAllocation", Integer, "token allocation")
		Attribute("badge", BadgeMedia, "task badge")
		Attribute("logoURL", String, "logo URL")
		Attribute("resourceId", String, "learning resource ID")
		Attribute("resourceURL", String, "learning resource URL")
		Required(
			"id",
			"title",
			"type",
			"author",
			"description",
			"token",
			"tokenAllocation",
			"badge",
			"logoURL",
			"resourceId",
			"resourceURL",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("type")
		Attribute("author")
		Attribute("description")
		Attribute("token")
		Attribute("tokenAllocation")
		Attribute("badge")
		Attribute("logoURL")
		Attribute("resourceId")
		Attribute("resourceURL")
	})
})

// TasksMedia ...
var TasksMedia = MediaType("application/vnd.tasks+json", func() {
	Description("Tasks")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("tasks", CollectionOf(TaskMedia), "list of tasks")
		Required("tasks")
	})
	View("default", func() {
		Attribute("tasks")
	})
})

// WalletMedia ...
var WalletMedia = MediaType("application/vnd.wallet+json", func() {
	Description("A wallet")
	Attributes(func() {
		Attribute("id", String, "Wallet ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("userId", String, "User ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("address", String, "wallet address")
		Attribute("walletType", String, "wallet type")
		Required(
			"id",
			"address",
			"walletType",
		)
	})
	View("default", func() {
		Attribute("address")
		Attribute("walletType")
	})
})

// WalletsMedia ...
var WalletsMedia = MediaType("application/vnd.wallets+json", func() {
	Description("Wallets")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("wallets", CollectionOf(WalletMedia), "list of wallets")
		Required("wallets")
	})
	View("default", func() {
		Attribute("wallets")
	})
})

// RedditUserMedia ...
var RedditUserMedia = MediaType("application/vnd.reddituser+json", func() {
	Description("A Reddit User")
	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("userId", String, "User ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("username", String, "Username")
		Attribute("linkKarma", Integer, "Link Karma")
		Attribute("commentKarma", Integer, "Comment Karma")
		Attribute("trophies", ArrayOf(String), "User trophies")
		Attribute("subreddits", ArrayOf(String), "User subreddits")
		Attribute("verification", VerificationMedia, "Social Account Verification")
		Required(
			"id",
			"userId",
			"username",
			"linkKarma",
			"commentKarma",
			"trophies",
			"subreddits",
			"verification",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("userId")
		Attribute("username")
		Attribute("linkKarma")
		Attribute("commentKarma")
		Attribute("trophies")
		Attribute("subreddits")
		Attribute("verification")
	})
})

// StackOverflowUserMedia ...
var StackOverflowUserMedia = MediaType("application/vnd.stackoverflowuser+json", func() {
	Description("Stack Overflow User Info")
	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("userId", String, "User ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("exchangeAccountId", Integer, "Stack Exchange Account ID")
		Attribute("stackUserId", Integer, "Stack Overflow Community-Specific Account ID")
		Attribute("displayName", String, "Display Name")
		Attribute("accounts", ArrayOf(String), "Stack Exchange Accounts")
		Attribute("verification", VerificationMedia, "Social Account Verification")
		Required(
			"id",
			"userId",
			"exchangeAccountId",
			"stackUserId",
			"displayName",
			"accounts",
			"verification",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("userId")
		Attribute("exchangeAccountId")
		Attribute("stackUserId")
		Attribute("displayName")
		Attribute("accounts")
		Attribute("verification")
	})
})

// QuizFields ...
var QuizFields = MediaType("application/vnd.quiz-fields+json", func() {
	Description("Quiz fields")
	Attributes(func() {
		Attribute("question", String, "Question")
		Attribute("answer", String, "Answer")
		Required("question", "answer")
	})
	View("default", func() {
		Attribute("question")
		Attribute("answer")
	})
})

// QuizMedia ...
var QuizMedia = MediaType("application/vnd.quiz+json", func() {
	Description("Quiz")
	Attributes(func() {
		Attribute("id", String, "Quiz ID")
		Attribute("title", String, "Quiz title")
		Attribute("userId", String, "Quiz user ID")
		Attribute("fields", CollectionOf(QuizFields), "Quiz fields")
		Required("id", "title", "userId", "fields")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("userId")
		Attribute("fields")
	})
})

// QuizResultsMedia ...
var QuizResultsMedia = MediaType("application/vnd.results+json", func() {
	Description("Quiz results")
	Attributes(func() {
		Attribute("userId", String, "User ID")
		Attribute("quizId", String, "Quiz ID")
		Attribute("questionsCorrect", Integer, "Count of correct quiz answers")
		Attribute("questionsIncorrect", Integer, "Count of incorrect quiz answers")
		Required("userId", "quizId", "questionsCorrect", "questionsIncorrect")
	})
	View("default", func() {
		Attribute("userId")
		Attribute("quizId")
		Attribute("questionsCorrect")
		Attribute("questionsIncorrect")
	})
})

// VerificationMedia ...
var VerificationMedia = MediaType("application/vnd.verification+json", func() {
	Description("Account Verification")
	Attributes(func() {
		Attribute("postedVerificationCode", String, "Posted Verification Code")
		Attribute("confirmedVerificationCode", String, "Confirmed Verification Code")
		Attribute("verified", Boolean, "Account Verified Flag")
		Required(
			"postedVerificationCode",
			"confirmedVerificationCode",
			"verified",
		)
	})
	View("default", func() {
		Attribute("postedVerificationCode")
		Attribute("confirmedVerificationCode")
		Attribute("verified")
	})
})

// ProfileMedia defines the media type used to render users.
var ProfileMedia = MediaType("application/vnd.profile+json", func() {
	Description("A user profile")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", String, "Unique user ID")
		Attribute("name", String, "Name")
		Attribute("username", String, "Username")
		Required("id", "name", "username")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must
		Attribute("name")
		Attribute("username")
	})
})
