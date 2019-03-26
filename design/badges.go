package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("badges", func() {
	BasePath("/v1/badges")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("list", func() {
		Description("Get list of user badges")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, BadgesMedia)
	})

	Action("show", func() {
		Description("Get all badges")
		Routing(GET(""))
		Response(OK, CollectionOf(BadgeMedia))
	})

	Action("create", func() {
		Description("Create a badge")
		Routing(POST(""))
		Payload(CreateBadgePayload)
		Response(OK, BadgeMedia)
	})
})
