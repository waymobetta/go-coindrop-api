package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("redditharvest", func() {
	BasePath("/v1/social/reddit/harvest")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("update", func() {
		Description("Update Reddit User Info")
		Routing(POST(""))
		Payload(UpdateUserPayload)
		Response(OK, RedditUserMedia)
	})
})

// UpdateUserPayload is the payload for updating a user's reddit info
var UpdateUserPayload = Type("UpdateUserPayload", func() {
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
