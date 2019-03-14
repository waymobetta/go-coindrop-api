package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
)

// StackoverflowharvestController implements the stackoverflowharvest resource.
type StackoverflowharvestController struct {
	*goa.Controller
}

// NewStackoverflowharvestController creates a stackoverflowharvest controller.
func NewStackoverflowharvestController(service *goa.Service) *StackoverflowharvestController {
	return &StackoverflowharvestController{Controller: service.NewController("StackoverflowharvestController")}
}

// Update runs the update action.
func (c *StackoverflowharvestController) Update(ctx *app.UpdateStackoverflowharvestContext) error {
	// StackoverflowharvestController_Update: start_implement

	// Put your logic here

	res := &app.Stackoverflowuser{}
	return ctx.OK(res)
	// StackoverflowharvestController_Update: end_implement
}
