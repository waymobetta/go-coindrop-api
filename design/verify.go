package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("verifyreddit", func() {
	BasePath("/v1/social/reddit/userid/verify")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

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
	})

	Action("update", func() {
		Description("Update Reddit Verification Code")
		Routing(POST(""))
		Payload(UpdateUserPayload)
		Response(OK, RedditUserMedia)
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

// VerificationPayload is the payload for updating verification data of a social account
var VerificationPayload = Type("VerificationPayload", func() {
	Description("Social Account Verification Payload")
	Attribute("userId", String, "User ID", func() {
		Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
		Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
	})
	Attribute("postedVerificationCode", String, "Verification Code Posted In Social Forum")
	Required(
		"userId",
		"postedVerificationCode",
	)
})
