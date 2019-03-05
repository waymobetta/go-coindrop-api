package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("quizzes", func() {
	BasePath("/v1/quizzes")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("create", func() {
		Description("Create quiz")
		Routing(POST(""))
		Payload(QuizPayload)
		Response(OK, QuizMedia)
	})

	Action("list", func() {
		Description("Get quizzes")
		Routing(GET(""))
		Response(OK, CollectionOf(QuizMedia))
	})

	Action("show", func() {
		Description("Get quizzes")
		Routing(GET("/:quizId"))
		Params(func() {
			Param("quizId", String, "Quiz ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, QuizMedia)
	})
})

// QuizMedia ...
var QuizMedia = MediaType("application/vnd.quiz+json", func() {
	Description("Quiz")
	Attributes(func() {
		Attribute("id", String, "Quiz ID")
		Attribute("title", String, "Quiz title")
		Attribute("userId", String, "Quiz user ID")
		Attribute("fields", CollectionOf(QuizFields), "Quiz fields")
		Required("id", "title", "userId", "fields")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("userId")
		Attribute("fields")
	})
})

// QuizFields ...
var QuizFields = MediaType("application/vnd.quiz-fields+json", func() {
	Description("Quiz fields")
	Attributes(func() {
		Attribute("question", String, "Question")
		Attribute("answer", String, "Answer")
		Required("question", "answer")
	})
	View("default", func() {
		Attribute("question")
		Attribute("answer")
	})
})

// QuizPayload is the payload for creating a quiz
var QuizPayload = Type("QuizPayload", func() {
	Description("Quiz payload")
	Attribute("title", String, "Title")
	Required("title")
})
