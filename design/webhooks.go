package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("webhooks", func() {
	BasePath("/v1/webhooks")

	NoSecurity()

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("typeform", func() {
		Description("Typeform webhook")
		Routing(POST("/typeform"))
		Payload(TypeformPayload)
		Response(OK)
	})
})

// TypeformPayload is the payload for webhook.
var TypeformPayload = Type("TypeformPayload", func() {
	Description("Typeform payload")
	Attribute("data", Any, "Data", func() {
	})
	Required("data")
})
