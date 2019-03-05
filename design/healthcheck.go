package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("healthcheck", func() { // Resources group related API endpoints
	BasePath("/v1/health") // together. They map to REST resources for REST

	NoSecurity()

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	DefaultMedia(HealthcheckMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Returns OK if system is healthy")
		Routing(GET(""))
		Response(OK, HealthcheckMedia)
	})
})

// HealthcheckMedia ...
var HealthcheckMedia = MediaType("application/vnd.healthcheck+json", func() {
	Description("Health check")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("status", String, "Status")
		Required("status")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("status")
	})
})
