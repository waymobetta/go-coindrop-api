package design

/*

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("stackoverflow", func() {
	BasePath("/v1/stackoverflow")

	Security(JWTAuth)

	Action("show", func() {
		Description("Get stack overflow user info")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, StackOverflowUserMedia)
		Response(NotFound, StandardErrorMedia)
	})

	// Action("update", func() {
	// 	Description("Update stack overflow user info")
	// 	Routing(POST(""))
	// 	Payload(WalletPayload)
	// 	Response(OK)
	// 	Response(NotFound)
	// 	Response(BadRequest, StandardErrorMedia)
	// 	Response(Gone, StandardErrorMedia)
	// 	Response(InternalServerError, StandardErrorMedia)
	// })
})

// StackOverflowUserMedia ...
var StackOverflowUserMedia = MediaType("application/vnd.stackoverflowuser+json", func() {
	Description("Stack Overflow User Info")
	Attributes(func() {
		Attribute("id", String, "ID")
		Attribute("userId", String, "User ID", func() {
			Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
			Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
		})
		Attribute("exchangeAccountId", String, "Stack Exchange Account ID")
		Attribute("stackUserId", String, "Stack Overflow Community-Specific Account ID")
		Attribute("displayName", String, "Display Name")
		Attribute("accounts", CollectionOf(String), "Stack Exchange Accounts")
		Required("accounts")
	})
	View("default", func() {
		Attribute("displayName")
	})
})

// // StackOverflowUserPayload is the payload for updating a user's wallet
// var WalletPayload = Type("WalletPayload", func() {
// 	Description("Wallet payload")
// 	Attribute("walletAddress", String, "Wallet address", func() {
// 		Pattern("^0x[0-9a-fA-F]{40}$")
// 		Example("0x845fdD93Cca3aE9e380d5556818e6d0b902B977c")
// 	})
// 	Required("walletAddress")
// })

*/
