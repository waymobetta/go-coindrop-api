package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// QuizController implements the quiz resource.
type QuizController struct {
	*goa.Controller
	db *db.DB
}

// NewQuizController creates a quiz controller.
func NewQuizController(service *goa.Service, dbs *db.DB) *QuizController {
	return &QuizController{
		Controller: service.NewController("QuizController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *QuizController) Show(ctx *app.ShowQuizContext) error {
	// QuizController_Show: start_implement

	// Put your logic here

	quiz := new(types.Quiz)
	quiz.Title = ctx.Params.Get("quizTitle")

	_, err := c.db.GetQuiz(quiz)
	if err != nil {
		log.Errorf("[controller/quiz] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get quiz info from db",
		})
	}

	log.Printf("[controller/quiz] returned information for coindrop quiz: %v\n", quiz.Title)

	res := &app.Quiz{
		QuizObject: quiz,
	}
	return ctx.OK(res)
	// QuizController_Show: end_implement
}
