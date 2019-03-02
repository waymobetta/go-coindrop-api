package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// VerifyredditController implements the verifyreddit resource.
type VerifyredditController struct {
	*goa.Controller
	db *db.DB
}

// NewVerifyredditController creates a verifyreddit controller.
func NewVerifyredditController(service *goa.Service, dbs *db.DB) *VerifyredditController {
	return &VerifyredditController{
		Controller: service.NewController("VerifyredditController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *VerifyredditController) Show(ctx *app.ShowVerifyredditContext) error {
	// VerifyredditController_Show: start_implement

	// Put your logic here

	res := &app.Reddituser{}
	return ctx.OK(res)
	// VerifyredditController_Show: end_implement
}

// Update runs the update action.
func (c *VerifyredditController) Update(ctx *app.UpdateVerifyredditContext) error {
	// VerifyredditController_Update: start_implement

	// Put your logic here

	res := &app.Reddituser{}
	return ctx.OK(res)
	// VerifyredditController_Update: end_implement
}
