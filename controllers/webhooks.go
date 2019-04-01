package controllers

import (
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	ethsvc "github.com/waymobetta/go-coindrop-api/services/ethereum"
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
			hidden := *form.Hidden
			if hidden.UserID != nil {
				userID = *hidden.UserID
			}
		}
	}

	incorrect := answersCount - correct
	if incorrect < 0 {
		incorrect = 0
	}

	results := &types.QuizResults{
		QuizID:             "",
		TypeformFormID:     formID,
		UserID:             userID,
		QuestionsCorrect:   correct,
		QuestionsIncorrect: incorrect,
		QuizTaken:          true,
	}

	log.Print("[controller/webhooks] input data\n")

	fmt.Printf("Adding quiz results: \nQuiz ID: %s\nTypeformID: %s\nUserID: %s\nCorrect: %v\nIncorrect: %v\nTaken: %v\n", results.QuizID, results.TypeformFormID, results.UserID, results.QuestionsCorrect, results.QuestionsIncorrect, results.QuizTaken)

	_, err := c.db.AddQuizResults(results)
	if err != nil {
		log.Errorf("[controller/webhooks] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not store quiz results",
		})
	}

	_, err = c.db.MarkUserTaskCompletedFromQuiz(results)
	if err != nil {
		log.Errorf("[controller/webhooks] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could mark user task complete",
		})
	}

	// TODO:
	// add badge grant, token send

	wallet, err := c.db.GetWallet(results.UserID, "eth")
	if err != nil {
		log.Errorf("[controller/webhooks] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get user wallet",
		})
	}

	ethAmountInWei := int64(5000000000000000000)

	tx, err := ethsvc.SendEther(wallet.Address, ethAmountInWei)
	if err != nil {
		log.Errorf("[controller/webhooks] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not send ether",
		})
	}

	// // 1 correct answer = 100 token
	// tokenMultiplier := 100
	// tokenAmount := results.QuestionsCorrect * tokenMultiplier

	// // if token 9 decimals
	// // default: 18
	// tokenAmountInWei := fmt.Sprintf("%v000000000", tokenAmount)
	// recipientAddress := wallet.Address

	// tx, err := ethsvc.SendToken()
	// if err != nil {
	// 	log.Errorf("[controller/webhooks] %v", err)
	// 	return ctx.InternalServerError(&app.StandardError{
	// 		Code:    500,
	// 		Message: "could not send token",
	// 	})
	// }

	// log.Printf("sent %v token to %s\n", tokenAmount, wallet.Address)
	log.Printf("https://rinkeby.etherscan.io/tx/%s\n", tx)

	return nil
	// WebhooksController_Typeform: end_implement
}
