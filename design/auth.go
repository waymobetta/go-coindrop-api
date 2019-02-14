package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// JWTAuth ...
var JWTAuth = BasicAuthSecurity("JWTAuth")
