package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("stackoverflowharvest", func() {
	BasePath("/v1/internal/social/stackoverflow/harvest")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("updateProfile", func() {
		Description("Update Stack Overflow User Info")
		Routing(POST("/profile"))
		Payload(UpdateStackOverflowUserPayload)
		Response(OK, StackOverflowUserMedia)
	})
	Action("updateCommunities", func() {
		Description("Update Stack Overflow User Communities Info")
		Routing(POST("/communities"))
		Payload(UpdateStackOverflowUserPayload)
		Response(OK, StackOverflowUserMedia)
	})
})
