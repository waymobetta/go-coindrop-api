package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// StackoverflowharvestController implements the stackoverflowharvest resource.
type StackoverflowharvestController struct {
	*goa.Controller
	db *db.DB
}

// NewStackoverflowharvestController creates a stackoverflowharvest controller.
func NewStackoverflowharvestController(service *goa.Service, dbs *db.DB) *StackoverflowharvestController {
	return &StackoverflowharvestController{
		Controller: service.NewController("StackoverflowharvestController"),
		db:         dbs,
	}
}

// Update runs the update action.
func (c *StackoverflowharvestController) Update(ctx *app.UpdateStackoverflowharvestContext) error {
	// StackoverflowharvestController_Update: start_implement

	// Put your logic here

	res := &app.Stackoverflowuser{}
	return ctx.OK(res)
	// StackoverflowharvestController_Update: end_implement
}
