package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	//Scope("api:access", "API access") // Define "api:access" scope
})

var _ = Resource("tasks", func() {
	BasePath("/v1/tasks")

	Security(JWT)

	Action("show", func() {
		Description("Get user tasks")
		Security(JWTAuth)
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

var BadgeMedia = MediaType("application/vnd.badge+json", func() {
	Description("Badge")
	Attributes(func() {
		Attribute("id", Integer, "badge ID")
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
		Attribute("id", Integer, "task ID")
		Attribute("title", String, "task title")
		Attribute("type", String, "task type")
		Attribute("author", String, "task author")
		Attribute("description", String, "task description")
		Attribute("token", String, "task token")
		Attribute("tokenAllocation", Integer, "token allocation")
		Attribute("badge", BadgeMedia, "task badge")
		Attribute("isAssigned", Boolean, "task assigned flag")
		Attribute("isCompleted", Boolean, "task completed flag")
		Required(
			"id",
			"title",
			"type",
			"author",
			"description",
			"token",
			"tokenAllocation",
			"badge",
			"isAssigned",
			"isCompleted",
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
		Attribute("isAssigned")
		Attribute("isCompleted")
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

// TaskPayload is the payload for creating a task
var TaskPayload = Type("TaskPayload", func() {
	Description("Task payload")
	Attribute("cognitoAuthUserId", String, "Cognito auth user ID")
	Attribute("taskName", String, "task name")
	Attribute("taskState", String, "task state")
	Required("cognitoAuthUserId", "taskName", "taskState")
})
