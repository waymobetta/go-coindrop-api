package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("public", func() { // Resources group related API endpoints
	BasePath("/v1/public") // together. They map to REST resources for REST

	// Security(JWTAuth)
	NoSecurity()

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	// DefaultMedia(ProfileMedia)

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get profile by Reddit username") // with its path, parameters (both path
		Routing(GET("/:redditUsername"))              // parameters and querystring values) and payload
		Params(func() {
			Param("redditUsername", String, "Reddit Username")
		})
		Response(OK, PublicMedia) // Responses define the shape and status code
	})
})