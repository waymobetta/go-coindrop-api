package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// WalletController implements the wallet resource.
type WalletController struct {
	*goa.Controller
	db *db.DB
}

// NewWalletController creates a wallet controller.
func NewWalletController(service *goa.Service, db *db.DB) *WalletController {
	return &WalletController{
		Controller: service.NewController("WalletController"),
		db:         db,
	}
}

// Show runs the show action.
func (c *WalletController) Show(ctx *app.ShowWalletContext) error {
	// WalletController_Show: start_implement

	// Put your logic here

	user := new(db.User)
	user.AuthUserID = ctx.Params.Get("userId")

	// return a user's wallet using the AWS cognito user ID as the key
	_, err := c.db.GetWallet(user)
	if err != nil {
		log.Errorf("[controller/user] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get wallet from db",
		})
	}

	log.Printf("[controller/user] returned wallet for coindrop user: %v\n", user.AuthUserID)

	res := &app.Wallet{
		WalletAddress: user.WalletAddress,
	}

	return ctx.OK(res)
	// WalletController_Show: end_implement
}

// Update runs the update action.
func (c *WalletController) Update(ctx *app.UpdateWalletContext) error {
	// WalletController_Update: start_implement

	// Put your logic here

	return nil
	// WalletController_Update: end_implement
}
