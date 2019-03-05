package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
)

// QuizzesController implements the quizzes resource.
type QuizzesController struct {
	*goa.Controller
}

// NewQuizzesController creates a quizzes controller.
func NewQuizzesController(service *goa.Service) *QuizzesController {
	return &QuizzesController{Controller: service.NewController("QuizzesController")}
}

// Create runs the create action.
func (c *QuizzesController) Create(ctx *app.CreateQuizzesContext) error {
	// QuizzesController_Create: start_implement

	// Put your logic here

	res := &app.Quiz{}
	return ctx.OK(res)
	// QuizzesController_Create: end_implement
}

// List runs the list action.
func (c *QuizzesController) List(ctx *app.ListQuizzesContext) error {
	// QuizzesController_List: start_implement

	// Put your logic here

	res := app.QuizCollection{}
	return ctx.OK(res)
	// QuizzesController_List: end_implement
}

// Show runs the show action.
func (c *QuizzesController) Show(ctx *app.ShowQuizzesContext) error {
	// QuizzesController_Show: start_implement

	// Put your logic here

	res := &app.Quiz{}
	return ctx.OK(res)
	// QuizzesController_Show: end_implement
}
