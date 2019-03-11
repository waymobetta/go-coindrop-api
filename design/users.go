package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("users", func() { // Resources group related API endpoints
	BasePath("/v1/users") // together. They map to REST resources for REST

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	DefaultMedia(UserMedia) // services.

	Action("create", func() {
		Description("Create a new user")
		Routing(POST(""))
		Payload(UserPayload)
		Response(OK)
		NoSecurity()
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get user by id") // with its path, parameters (both path
		Routing(GET("/:userId"))      // parameters and querystring values) and payload
		Params(func() {               // (shape of the request body).
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK) // Responses define the shape and status code
		NoSecurity()
	})

	Action("list", func() { // Actions define a single API endpoint together
		Description("Get user ID mapped to Cognito auth user ID") // with its path, parameters (both path

		Routing(GET(""))

		Params(func() { // (shape of the request body).
			Param("cognitoAuthUserId", String, "Cognito Auth User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})

		Response(OK) // Responses define the shape and status code
		NoSecurity()
	})
})
