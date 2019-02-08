package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

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
