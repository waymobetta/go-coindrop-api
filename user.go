package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	res := &app.GoaExampleUser{}
	return ctx.OK(res)
	// UserController_Show: end_implement
}
