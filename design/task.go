package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("tasks", func() {
	BasePath("/v1/tasks")

	Security(JWTAuth)

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

	Action("create", func() {
		Description("Create a user task")
		Routing(POST(""))
		Payload(CreateTaskPayload)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
	})

	Action("update", func() {
		Description("Update user task state")
		Routing(POST("/:taskId"))
		Params(func() {
			Param("taskId", String, "Task ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})

		Payload(TaskPayload)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
	})
})

var BadgeMedia = MediaType("application/vnd.badge+json", func() {
	Description("Badge")
	Attributes(func() {
		Attribute("id", String, "badge ID")
		Attribute("name", String, "badge name")
		Attribute("description", String, "badge description")
		Attribute("recipients", Integer, "badge recipients")
		Required(
			"id",
			"name",
			"description",
			"recipients",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("recipients")
	})
})

// TaskMedia ...
var TaskMedia = MediaType("application/vnd.task+json", func() {
	Description("Task")
	Attributes(func() {
		Attribute("id", String, "task ID")
		Attribute("title", String, "task title")
		Attribute("type", String, "task type")
		Attribute("author", String, "task author")
		Attribute("description", String, "task description")
		Attribute("token", String, "task token")
		Attribute("tokenAllocation", Integer, "token allocation")
		Attribute("badge", BadgeMedia, "task badge")
		Required(
			"id",
			"title",
			"type",
			"author",
			"description",
			"token",
			"tokenAllocation",
			"badge",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("type")
		Attribute("author")
		Attribute("description")
		Attribute("token")
		Attribute("tokenAllocation")
		Attribute("badge")
	})
})

// TasksMedia ...
var TasksMedia = MediaType("application/vnd.tasks+json", func() {
	Description("Tasks")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("tasks", CollectionOf(TaskMedia), "list of tasks")
		Required("tasks")
	})
	View("default", func() {
		Attribute("tasks")
	})
})

// CreateTaskPayload is the payload for creating a task
var CreateTaskPayload = Type("CreateTaskPayload", func() {
	Description("Create Task payload")
	Attribute("userId", String, "User ID")
	Attribute("taskId", String, "Task ID")
	Required("userId", "taskId")
})

// TaskPayload is the payload for updating a task
var TaskPayload = Type("TaskPayload", func() {
	Description("Task payload")
	Attribute("completed", Boolean, "Task completed")
	Required("completed")
})
