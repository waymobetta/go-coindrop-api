package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("reddit", func() {
	BasePath("/v1/social/reddit")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("show", func() {
		Description("Get Reddit User")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, RedditUserMedia)
	})

	Action("update", func() {
		Description("Update Reddit User")
		Routing(POST(""))
		Payload(CreateUserPayload)
		Response(OK, RedditUserMedia)
	})

	Action("verify", func() {
		Description("Update Reddit Verification")
		Routing(POST("/:userId/verify"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Payload(VerificationPayload)
		Response(OK, RedditUserMedia)
	})

	Action("display", func() {
		Description("Get Reddit Verification")
		Routing(GET("/:userId/verify"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, RedditUserMedia)
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

// CreateUserPayload is the payload for creating a listing for a user's reddit info
var CreateUserPayload = Type("CreateUserPayload", func() {
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
