package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("wallet", func() {
	BasePath("/v1/wallets")

	Action("show", func() {
		Description("Get user wallet")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, WalletMedia)
		Response(NotFound, StandardErrorMedia)
	})
})

// WalletMedia ...
var WalletMedia = MediaType("application/vnd.wallet+json", func() {
	Description("A wallet")
	Attributes(func() {
		Attribute("userId", String, "User ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("walletAddress", String, "wallet address")
		Required("walletAddress")
	})
	View("default", func() {
		Attribute("walletAddress")
	})
})
