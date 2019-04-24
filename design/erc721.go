package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("erc721", func() {
	BasePath("/v1/internal/erc721")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("assign", func() {
		Description("Assign an ERC721 to a user")
		Routing(POST("/assign"))
		Payload(AssignERC721Payload)
		Response(OK, ERC721Media)
	})
})
