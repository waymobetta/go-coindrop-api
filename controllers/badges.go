package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// BadgesController implements the badges resource.
type BadgesController struct {
	*goa.Controller
	db *db.DB
}

// NewBadgesController creates a badges controller.
func NewBadgesController(service *goa.Service, dbs *db.DB) *BadgesController {
	return &BadgesController{
		Controller: service.NewController("BadgesController"),
		db:         dbs,
	}
}

// Create runs the create action.
func (c *BadgesController) Create(ctx *app.CreateBadgesContext) error {
	// BadgesController_Create: start_implement

	// Put your logic here

	res := &app.Badge{}
	return ctx.OK(res)
	// BadgesController_Create: end_implement
}

// List runs the list action.
func (c *BadgesController) List(ctx *app.ListBadgesContext) error {
	// BadgesController_List: start_implement

	// Put your logic here

	res := &app.Badges{}
	return ctx.OK(res)
	// BadgesController_List: end_implement
}

// Show runs the show action.
func (c *BadgesController) Show(ctx *app.ShowBadgesContext) error {
	// BadgesController_Show: start_implement

	// Put your logic here

	res := app.BadgeCollection{}
	return ctx.OK(res)
	// BadgesController_Show: end_implement
}
