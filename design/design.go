package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = API("coindrop", func() { // API defines the microservice endpoint and
	Title("The Coindrop API")           // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:5000")

	// Sets CORS response headers for requests with any Origin header
	Origin("*", func() {
		Methods("OPTIONS", "HEAD", "POST", "GET", "UPDATE", "DELETE", "PATCH")
		Credentials()
	})
})
