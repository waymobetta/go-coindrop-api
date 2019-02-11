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
func NewWalletController(service *goa.Service, dbs *db.DB) *WalletController {
	return &WalletController{
		Controller: service.NewController("WalletController"),
		db:         dbs,
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
		log.Errorf("[controller/wallet] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get wallet from db",
		})
	}

	log.Printf("[controller/wallet] returned wallet for coindrop user: %v\n", user.AuthUserID)

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

	user := new(db.User)
	user.AuthUserID = ctx.Payload.CognitoAuthUserID
	user.WalletAddress = ctx.Payload.WalletAddress

	_, err := c.db.UpdateWallet(user)
	if err != nil {
		log.Errorf("[controller/wallet] %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not update wallet in db",
		})
	}

	log.Printf("[controller/wallet] successfully updated wallet for coindrop user: %v\n", user.AuthUserID)

	return ctx.OK(nil)
	// WalletController_Update: end_implement
}
