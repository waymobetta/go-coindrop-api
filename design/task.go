package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("tasks", func() {
	BasePath("/v1/tasks")

	Action("show", func() {
		Description("Get user tasks")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, TasksMedia)
		Response(NotFound, StandardErrorMedia)
	})

	Action("update", func() {
		Description("Update user task state")
		Routing(POST(""))
		Payload(TaskPayload)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
	})
})

// TasksMedia ...
var TasksMedia = MediaType("application/vnd.tasks+json", func() {
	Description("Tasks")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("taskList", Any, "list of tasks")
		Required("taskList")
	})
	View("default", func() {
		Attribute("taskList")
	})
})

// TaskPayload is the payload for creating a task
var TaskPayload = Type("TaskPayload", func() {
	Description("Task payload")
	Attribute("cognitoAuthUserId", String, "Cognito auth user ID")
	Attribute("taskName", String, "task name")
	Attribute("taskState", String, "task state")
	Required("cognitoAuthUserId", "taskName", "taskState")
})
