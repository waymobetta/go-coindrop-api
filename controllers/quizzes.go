package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// QuizzesController implements the quiz resource.
type QuizzesController struct {
	*goa.Controller
	db *db.DB
}

// NewQuizzesController creates a quiz controller.
func NewQuizzesController(service *goa.Service, dbs *db.DB) *QuizzesController {
	return &QuizzesController{
		Controller: service.NewController("QuizzesController"),
		db:         dbs,
	}
}

// Create runs the create action.
func (c *QuizzesController) Create(ctx *app.CreateQuizzesContext) error {
	// QuizzesController_Create: start_implement

	// Put your logic here

	quiz := &types.Quiz{
		Title: ctx.Payload.Title,
	}

	var err error
	quiz, err = c.db.AddQuiz(quiz)
	if err != nil {
		log.Errorf("[controller/quiz] failed to add quiz: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve quiz",
		})
	}

	log.Printf("[controller/quiz] returned information for coindrop quiz: %v\n", quiz.Title)

	return ctx.OK(&app.Quiz{
		ID:     quiz.ID,
		Title:  quiz.Title,
		UserID: "",
		Fields: app.QuizFieldsCollection{},
	})
	// QuizzesController_Create: end_implement
}

// List runs the list action.
func (c *QuizzesController) List(ctx *app.ListQuizzesContext) error {

	// QuizController_List: start_implement

	// Put your logic here

	quizzes, err := c.db.GetQuizzes()
	if err != nil {
		log.Errorf("[controller/quiz] failed to get quizzes: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve quizzes",
		})
	}

	log.Printf("[controller/quiz] returned information for coindrop quizzes. count: %v\n", len(quizzes))

	var appQuizzes app.QuizCollection
	for _, quiz := range quizzes {
		appQuizzes = append(appQuizzes, &app.Quiz{
			ID:     quiz.ID,
			Title:  quiz.Title,
			UserID: "",
			Fields: app.QuizFieldsCollection{},
		})
	}

	return ctx.OK(appQuizzes)
	// QuizController_List: end_implement
}

// Show runs the show action.
func (c *QuizzesController) Show(ctx *app.ShowQuizzesContext) error {
	// QuizController_Show: start_implement

	// Put your logic here

	quizID := ctx.QuizID

	quiz, err := c.db.GetQuiz(quizID)
	if err != nil {
		log.Errorf("[controller/quiz] failed to get quiz: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve quiz",
		})
	}

	log.Printf("[controller/quiz] returned information for coindrop quiz: %v\n", quiz.Title)

	return ctx.OK(&app.Quiz{
		ID:     quiz.ID,
		Title:  quiz.Title,
		UserID: "",
		Fields: app.QuizFieldsCollection{},
	})
	// QuizController_Show: end_implement
}
