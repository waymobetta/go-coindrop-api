package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// TransactionsController implements the transactions resource.
type TransactionsController struct {
	*goa.Controller
	db *db.DB
}

// NewTransactionsController creates a transactions controller.
func NewTransactionsController(service *goa.Service, dbs *db.DB) *TransactionsController {
	return &TransactionsController{
		Controller: service.NewController("TransactionsController"),
		db:         dbs,
	}
}

// List runs the list action.
func (c *TransactionsController) List(ctx *app.ListTransactionsContext) error {
	// TransactionsController_List: start_implement

	// Put your logic here

	userID := ctx.Params.Get("userId")
	// Note: if query string `userId` is empty,
	// then use user ID from auth token
	if userID == "" {
		userID = ctx.Value("authUserID").(string)
	}

	transactions, err := c.db.GetUserTransactions(userID)
	if err != nil {
		log.Errorf("[controller/transactions] failed to get user transactions: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get user's transactions",
		})
	}

	var t app.TransactionCollection

	for _, transaction := range transactions {
		t = append(t, &app.Transaction{
			ID:     transaction.ID,
			UserID: transaction.UserID,
			TaskID: transaction.TaskID,
			Hash:   transaction.Hash,
		})
	}

	log.Printf("[controller/transactions] returned transactions for coindrop user: %v\n", userID)

	res := &app.Transactions{
		Transactions: t,
	}

	return ctx.OK(res)
	// TransactionsController_List: end_implement
}
