package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("redditharvest", func() {
	BasePath("/v1/social/reddit/harvest")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("updateAbout", func() {
		Description("Update Reddit User About Info")
		Routing(POST("/about"))
		Payload(UpdateRedditUserPayload)
		Response(OK, RedditUserMedia)
	})

	Action("updateTrophies", func() {
		Description("Update Reddit User Trophy Info")
		Routing(POST("/trophies"))
		Payload(UpdateRedditUserPayload)
		Response(OK, RedditUserMedia)
	})

	Action("updateSubmittedInfo", func() {
		Description("Update Reddit User Submitted Info")
		Routing(POST("/submitted"))
		Payload(UpdateRedditUserPayload)
		Response(OK, RedditUserMedia)
	})
})
