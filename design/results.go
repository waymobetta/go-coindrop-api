package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("results", func() {
	BasePath("/v1/quizzes")

	Security(JWTAuth)

	Response(NotFound, StandardErrorMedia)
	Response(BadRequest, StandardErrorMedia)
	Response(Gone, StandardErrorMedia)
	Response(InternalServerError, StandardErrorMedia)

	Action("create", func() {
		Description("Add quiz results")
		Routing(POST("/results"))
		Payload(QuizResultsPayload)
		Response(OK, QuizResultsMedia)
	})

	Action("show", func() {
		Description("Get quiz results")
		Routing(GET("/:quizId/results"))
		Params(func() {
			Param("quizId", String, "Quiz ID", func() {
				Pattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
				Example("9302608f-f6a4-4004-b088-63e5fb43cc26")
			})
		})
		Response(OK, QuizResultsMedia)
	})

	Action("list", func() {
		Description("Get all quiz results")
		Routing(GET("/results"))
		Response(OK, CollectionOf(QuizResultsMedia))
	})
})
