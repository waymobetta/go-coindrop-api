package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// JWTAuth ...
var JWTAuth = BasicAuthSecurity("JWTAuth")

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
/*
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	//Scope("api:access", "API access") // Define "api:access" scope
})
*/
