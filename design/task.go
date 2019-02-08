package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("task", func() {
	BasePath("/v1/tasks")

	Action("show", func() {
		Description("Get user task")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, TaskMedia)
		Response(NotFound, StandardErrorMedia)
	})
})

// TaskMedia ...
var TaskMedia = MediaType("application/vnd.task+json", func() {
	Description("A task")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("taskName", String, "task name")
		Required("taskName")
	})
	View("default", func() {
		Attribute("taskName")
	})
})
