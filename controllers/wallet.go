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

	log.Println(ctx.Params.Get("userId"))

	walletAddress := "hello world"

	res := &app.Wallet{
		WalletAddress: walletAddress,
	}

	return ctx.OK(res)
	// WalletController_Show: end_implement
}
