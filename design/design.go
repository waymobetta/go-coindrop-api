package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("coindrop", func() { // API defines the microservice endpoint and
	Title("The Coindrop API")           // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})

var _ = Resource("user", func() { // Resources group related API endpoints
	BasePath("/v1/users")   // together. They map to REST resources for REST
	DefaultMedia(UserMedia) // services.

	Action("create", func() {
		Description("Create a new user")
		Routing(POST(""))
		Params(func() {
			Param("authUserID", String, "Cognito Auth User ID")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get user by id") // with its path, parameters (both path
		Routing(GET("/:userID"))      // parameters and querystring values) and payload
		Params(func() {               // (shape of the request body).
			Param("userID", Integer, "User ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// UserMedia defines the media type used to render users.
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique user ID")
		Attribute("name", String, "Name of user")
		Required("id", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must
		Attribute("name")
	})
})

// StandardErrorMedia defines the standard error type.
var StandardErrorMedia = MediaType("application/standard_error+json", func() {
	Description("A standard error response")

	Attributes(func() {
		Attribute("code", Integer, "A code that describes the error", func() {
			Example(400)
		})
		Attribute("message", String, "A message that describes the error", func() {
			Example("Bad Request")

		})

		Required("code", "message")
	})

	View("default", func() {
		Attribute("code")
		Attribute("message")
	})
})
