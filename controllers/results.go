package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// ResultsController implements the results resource.
type ResultsController struct {
	*goa.Controller
	db *db.DB
}

// NewResultsController creates a results controller.
func NewResultsController(service *goa.Service, dbs *db.DB) *ResultsController {
	return &ResultsController{
		Controller: service.NewController("ResultsController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *ResultsController) Show(ctx *app.ShowResultsContext) error {
	// ResultsController_Show: start_implement

	// Put your logic here

	quizResults := new(types.QuizResults)
	quizResults.AuthUserID = ctx.Params.Get("userId")

	allQuizResults := new(types.AllQuizResults)

	_, err := c.db.GetAllQuizResults(quizResults, allQuizResults)
	if err != nil {
		log.Errorf("[controller/results] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get quiz results from db",
		})
	}

	log.Printf("[controller/results] returned all quiz results for coindrop user: %v\n", quizResults.AuthUserID)

	res := &app.Results{
		QuizResultsList: allQuizResults.QuizResults,
	}
	return ctx.OK(res)
	// ResultsController_Show: end_implement
}
