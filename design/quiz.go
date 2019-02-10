package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("quiz", func() {
	BasePath("/v1/quiz")

	Action("show", func() {
		Description("Get quiz")
		Routing(GET(""))
		Params(func() {
			Param("quizTitle", String, "Quiz title")
		})
		Response(OK, QuizMedia)
		Response(NotFound, StandardErrorMedia)
	})
})

// QuizMedia ...
var QuizMedia = MediaType("application/vnd.quiz+json", func() {
	Description("Quiz")
	Attributes(func() {
		Attribute("quizObject", Any, "Quiz object")
		Required("quizObject")
	})
	View("default", func() {
		Attribute("quizObject")
	})
})
