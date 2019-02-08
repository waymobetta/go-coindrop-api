package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = API("coindrop", func() { // API defines the microservice endpoint and
	Title("The Coindrop API")           // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:5000")
})

var _ = Resource("user", func() { // Resources group related API endpoints
	BasePath("/v1/users")   // together. They map to REST resources for REST
	DefaultMedia(UserMedia) // services.

	Action("create", func() {
		Description("Create a new user")
		Routing(POST(""))
		Payload(UserPayload)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
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
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

var _ = Resource("wallet", func() {
	BasePath("/v1/wallets")

	Action("show", func() {
		Description("Get user wallet")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, WalletMedia)
		Response(NotFound, StandardErrorMedia)
	})
})

var _ = Resource("task", func() {
	BasePath("/v1/tasks")

	Action("show", func() {
		Description("Get user task")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, TaskMedia)
		Response(NotFound, StandardErrorMedia)
	})
})

// UserMedia defines the media type used to render users.
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", String, "Unique user ID")
		Attribute("cognitoAuthUserId", String, "Cognito auth user ID")
		Attribute("name", String, "Name of user")
		Attribute("walletAddress", String, "Wallet address")
		Required("id")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must
		Attribute("cognitoAuthUserId")
		Attribute("name")
		Attribute("walletAddress")
	})
})

// WalletMedia ...
var WalletMedia = MediaType("application/vnd.wallet+json", func() {
	Description("A wallet")
	Attributes(func() {
		Attribute("userId", String, "User ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("walletAddress", String, "wallet address")
		Required("walletAddress")
	})
	View("default", func() {
		Attribute("walletAddress")
	})
})

// TaskMedia ...
var TaskMedia = MediaType("application/vnd.task+json", func() {
	Description("A task")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("taskName", String, "task name")
		Required("taskName")
	})
	View("default", func() {
		Attribute("taskName")
	})
})

// UserPayload is the payload for creating a user
var UserPayload = Type("UserPayload", func() {
	Description("User payload")
	Attribute("cognitoAuthUserId", String, "Cognito auth user ID")
	Required("cognitoAuthUserId")
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
