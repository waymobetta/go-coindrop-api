package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("stackoverflow", func() {
	BasePath("/v1/social/stackoverflow")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("show", func() {
		Description("Get Stack Overflow User")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, StackOverflowUserMedia)
	})

	Action("update", func() {
		Description("Update stack overflow user info")
		Routing(POST(""))
		Payload(CreateStackOverflowUserPayload)
		Response(OK, StackOverflowUserMedia)
	})

	Action("verify", func() {
		Description("Update Stack Overflow Verification")
		Routing(POST("/:userId/verify"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Payload(VerificationPayload)
		Response(OK, StackOverflowUserMedia)
	})

	Action("display", func() {
		Description("Get Stack Overflow Verification")
		Routing(GET("/:userId/verify"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, StackOverflowUserMedia)
	})
})
