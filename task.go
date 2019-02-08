package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
)

// TaskController implements the task resource.
type TaskController struct {
	*goa.Controller
}

// NewTaskController creates a task controller.
func NewTaskController(service *goa.Service) *TaskController {
	return &TaskController{Controller: service.NewController("TaskController")}
}

// Show runs the show action.
func (c *TaskController) Show(ctx *app.ShowTaskContext) error {
	// TaskController_Show: start_implement

	// Put your logic here

	res := &app.Task{}
	return ctx.OK(res)
	// TaskController_Show: end_implement
}
