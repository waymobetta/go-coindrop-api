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
		Routing(GET("/badges/:redditUsername"))       // parameters and querystring values) and payload
		Params(func() {
			Param("redditUsername", String, "Reddit Username")
		})
		Response(OK, PublicBadgesMedia) // Responses define the shape and status code
	})

	Action("display", func() { // Actions define a single API endpoint together
		Description("Get task information by ERC721 token ID") // with its path, parameters (both path
		Routing(GET("/tokens/:erc721TokenId"))                 // parameters and querystring values) and payload
		Params(func() {
			Param("erc721TokenId", String, "ERC-721 token ID")
		})
		Response(OK, ERC721LookupMedia) // Responses define the shape and status code
	})
})
