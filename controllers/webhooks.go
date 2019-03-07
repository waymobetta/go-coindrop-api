package controllers

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
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
	var formID string
	var correct int
	var answersCount int
	var userID string

	if ctx.Payload.FormResponse != nil {
		form := *ctx.Payload.FormResponse
		if form.Calculated != nil {
			calc := *form.Calculated
			if calc.Score != nil {
				correct = *calc.Score
			}
		}

		if form.FormID != nil {
			formID = *form.FormID
		}

		if form.Answers != nil {
			answersCount = len(form.Answers.([]interface{}))
		}

		if form.Hidden != nil {
			if form.Hidden != nil {
				hidden := *form.Hidden
				if hidden.UserID != nil {
					userID = *hidden.UserID
				}
			}
		}
	}

	incorrect := answersCount - correct
	if incorrect < 0 {
		incorrect = 0
	}

	results := &types.QuizResults{
		TypeformFormID:     formID,
		UserID:             userID,
		QuestionsCorrect:   correct,
		QuestionsIncorrect: incorrect,
	}

	log.Print("[controller/hooks] input data\n")
	spew.Dump(results)

	_, err := c.db.AddQuizResults(results)
	if err != nil {
		log.Errorf("[controller/webhooks] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not store quiz results",
		})
	}

	//fmt.Println(ctx.Payload.FormResponse)
	return nil
	// WebhooksController_Typeform: end_implement
}
