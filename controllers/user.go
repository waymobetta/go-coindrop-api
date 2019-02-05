package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	db *db.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, db *db.DB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here

	res := &app.User{
		ID:   0,
		Name: "",
	}
	return ctx.OK(res)
	// UserController_Create: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UserController_Show: end_implement
}
