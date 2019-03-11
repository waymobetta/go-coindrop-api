package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("stackoverflowharvest", func() {
	BasePath("/v1/social/stackoverflow/harvest")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("update", func() {
		Description("Update Stack Overflow User Info")
		Routing(POST(""))
		Payload(UpdateStackOverflowUserPayload)
		Response(OK, StackOverflowUserMedia)
	})
})
