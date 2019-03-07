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
	Attribute("event_id", String, "Event ID")
	Attribute("event_type", String, "Event types")
	Attribute("form_response", TypeformFormPayload, "Form response")
})

// TypeformFormPayload ...
var TypeformFormPayload = Type("TypeformFormPayload", func() {
	Description("Typeform form data")
	Attribute("form_id", String, "Form ID")
	Attribute("token", String, "Form ID")
	Attribute("landed_at", String, "Form ID")
	Attribute("submitted_at", String, "Form ID")
	Attribute("calculated", TypeformCalculatedPayload, "Calculated response")
	Attribute("hidden", TypeformHiddenPayload, "Hidden")
	Attribute("definition", Any, "Definition")
	Attribute("answers", Any, "Answers")
})

// TypeformCalculatedPayload ...
var TypeformCalculatedPayload = Type("TypeformCalculatedPayload", func() {
	Description("Typeform calculatd data")
	Attribute("score", Integer, "Score")
})

// TypeformHiddenPayload ...
var TypeformHiddenPayload = Type("TypeformHiddenPayload", func() {
	Description("Typeform hidden data")
	Attribute("user_id", String, "User ID")
})
