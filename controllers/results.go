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

// Create runs the show action.
func (c *ResultsController) Create(ctx *app.CreateResultsContext) error {
	// ResultsController_Create: start_implement

	// Put your logic here

	results := &types.QuizResults{
		QuizID:             ctx.Payload.QuizID,
		UserID:             ctx.Payload.UserID,
		QuestionsCorrect:   ctx.Payload.QuestionsCorrect,
		QuestionsIncorrect: ctx.Payload.QuestionsIncorrect,
	}

	_, err := c.db.AddQuizResults(results)
	if err != nil {
		log.Errorf("[controller/results] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not store quiz results",
		})
	}

	return ctx.OK(nil)
	// ResultsController_Create: end_implement
}

// Show runs the show action.
func (c *ResultsController) Show(ctx *app.ShowResultsContext) error {
	// ResultsController_Show: start_implement

	// Put your logic here

	quizID := ctx.QuizID
	userID := ctx.Params.Get("userId")

	quizResults, err := c.db.GetQuizResults(quizID, userID)
	if err != nil {
		log.Errorf("[controller/results] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve quiz results",
		})
	}

	if quizResults == nil {
		log.Errorf("[controller/results] no results for quiz ID; %s", quizID)
		return ctx.OK(nil)
	}

	log.Printf("[controller/results] returned all quiz results for coindrop user: %v\n", userID)

	res := &app.Results{
		QuizID:             quizResults.QuizID,
		UserID:             quizResults.UserID,
		QuestionsCorrect:   quizResults.QuestionsCorrect,
		QuestionsIncorrect: quizResults.QuestionsIncorrect,
	}
	return ctx.OK(res)
	// ResultsController_Show: end_implement
}

// List runs the show action.
func (c *ResultsController) List(ctx *app.ListResultsContext) error {
	// ResultsController_List: start_implement

	// Put your logic here

	userID := ctx.Params.Get("userId")
	quizResults, err := c.db.GetAllQuizResults(userID)
	if err != nil {
		log.Errorf("[controller/results] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve quiz results",
		})
	}

	if quizResults == nil {
		log.Error("[controller/results] no results\n")
		return ctx.OK(nil)
	}

	log.Printf("[controller/results] returned all quiz results: count %v\n", len(quizResults))

	var resp app.ResultsCollection

	for _, results := range quizResults {
		resp = append(resp, &app.Results{
			QuizID:             results.QuizID,
			UserID:             results.UserID,
			QuestionsCorrect:   results.QuestionsCorrect,
			QuestionsIncorrect: results.QuestionsIncorrect,
		})
	}

	return ctx.OK(resp)
	// ResultsController_List: end_implement
}
