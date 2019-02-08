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
