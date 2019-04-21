package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("targeting", func() {
	BasePath("/v1/targeting")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("display", func() {
		Description("Get list of eligible users")
		Routing(GET("/users/:project"))
		Params(func() {
			Param("project", String, "Project name")
		})
		Response(OK, TargetingMedia)
	})

	Action("list", func() {
		Description("Get list of all reddit users and their subreddits")
		Routing(GET("/users/reddit"))
		Response(OK, RedditTargetingMedia)
	})

	Action("set", func() {
		Description("Set users as eligible")
		Routing(POST("/tasks/set"))
		Payload(SetTargetingPayload)
		Response(OK)
	})
})
