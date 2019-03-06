package controllers

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// WebhooksController implements the webhooks resource.
type WebhooksController struct {
	*goa.Controller
	db *db.DB
}

// NewWebhooksController creates a webhooks controller.
func NewWebhooksController(service *goa.Service, dbs *db.DB) *WebhooksController {
	return &WebhooksController{
		Controller: service.NewController("WebhooksController"),
		db:         dbs,
	}
}

// Typeform runs the typeform action.
func (c *WebhooksController) Typeform(ctx *app.TypeformWebhooksContext) error {
	// WebhooksController_Typeform: start_implement

	// Put your logic here
	fmt.Println("TODO")
	return nil
	// WebhooksController_Typeform: end_implement
}
