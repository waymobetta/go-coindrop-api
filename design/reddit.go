package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("reddit", func() {
	BasePath("/v1/social/reddit")

	Security(JWTAuth)

	Action("show", func() {
		Description("Get ")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, RedditUserMedia)
		Response(NotFound, StandardErrorMedia)
	})

	Action("update", func() {
		Description("Update Reddit User")
		Routing(POST(""))
		Payload(RedditUserPayload)
		Response(OK, RedditUserMedia)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
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
		Attribute("userId", String, "User ID")
		Attribute("username", String, "Username")
		Attribute("linkKarma", Integer, "Link Karma")
		Attribute("commentKarma", Integer, "Comment Karma")
		Attribute("accountCreatedUTC", Number, "Account Created Timestamp") // Number used in lieu of Float32/64
		Attribute("trophies", ArrayOf(String), "User trophies")
		Attribute("subreddits", ArrayOf(String), "User subreddits")
		Attribute("verification", VerificationMedia, "Social Account Verification")
		Required(
			"id",
			"userId",
			"username",
			"linkKarma",
			"commentKarma",
			"accountCreatedUTC",
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
		Attribute("accountCreatedUTC")
		Attribute("trophies")
		Attribute("subreddits")
		Attribute("verification")
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

// RedditUserPayload is the payload for updating a user's wallet
var RedditUserPayload = Type("RedditUserPayload", func() {
	Description("Reddit User payload")
	Attribute("id", String, "User ID", func() {
		Pattern("^0x[0-9a-fA-F]{40}$")
		Example("0x845fdD93Cca3aE9e380d5556818e6d0b902B977c")
	})
	Required(
		"id",
	)
})
