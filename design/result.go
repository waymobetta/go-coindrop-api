package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("results", func() {
	BasePath("/v1/quiz/results")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("show", func() {
		Description("Get quiz results")
		Routing(GET(""))
		Params(func() {
			Param("userId", String, "User ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, QuizResultsMedia)
	})
})

// QuizResultsMedia ...
var QuizResultsMedia = MediaType("application/vnd.results+json", func() {
	Description("Quiz results")
	Attributes(func() {
		Attribute("userId", String, "user ID")
		Attribute("quizResultsList", Any, "quiz results list")
		Required("quizResultsList")
	})
	View("default", func() {
		Attribute("quizResultsList")
	})
})
