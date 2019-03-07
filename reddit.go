package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
)

// RedditController implements the reddit resource.
type RedditController struct {
	*goa.Controller
}

// NewRedditController creates a reddit controller.
func NewRedditController(service *goa.Service) *RedditController {
	return &RedditController{Controller: service.NewController("RedditController")}
}

// Display runs the display action.
func (c *RedditController) Display(ctx *app.DisplayRedditContext) error {
	// RedditController_Display: start_implement

	// Put your logic here

	res := &app.Reddituser{}
	return ctx.OK(res)
	// RedditController_Display: end_implement
}

// Show runs the show action.
func (c *RedditController) Show(ctx *app.ShowRedditContext) error {
	// RedditController_Show: start_implement

	// Put your logic here

	res := &app.Reddituser{}
	return ctx.OK(res)
	// RedditController_Show: end_implement
}

// Update runs the update action.
func (c *RedditController) Update(ctx *app.UpdateRedditContext) error {
	// RedditController_Update: start_implement

	// Put your logic here

	res := &app.Reddituser{}
	return ctx.OK(res)
	// RedditController_Update: end_implement
}

// Verify runs the verify action.
func (c *RedditController) Verify(ctx *app.VerifyRedditContext) error {
	// RedditController_Verify: start_implement

	// Put your logic here

	res := &app.Reddituser{}
	return ctx.OK(res)
	// RedditController_Verify: end_implement
}
