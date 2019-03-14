package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("profiles", func() { // Resources group related API endpoints
	BasePath("/v1/profiles") // together. They map to REST resources for REST

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	DefaultMedia(ProfileMedia)

	Action("create", func() {
		Description("Upsert a new profile")
		Routing(POST(""))
		Payload(ProfilePayload)
		Response(OK)
	})

	Action("update", func() { // Actions define a single API endpoint together
		Description("Update profile") // with its path, parameters (both path
		Routing(POST("/:userId"))     // parameters and querystring values) and payload
		Payload(ProfilePayload)
		Response(OK) // Responses define the shape and status code
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get profile by user id") // with its path, parameters (both path
		Routing(GET("/:userId"))              // parameters and querystring values) and payload
		Response(OK)                          // Responses define the shape and status code
	})

	Action("list", func() { // Actions define a single API endpoint together
		Description("Get user profile") // with its path, parameters (both path

		Routing(GET(""))
		Response(OK) // Responses define the shape and status code
	})
})
