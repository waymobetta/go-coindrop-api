package main

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
)

// WalletController implements the wallet resource.
type WalletController struct {
	*goa.Controller
}

// NewWalletController creates a wallet controller.
func NewWalletController(service *goa.Service) *WalletController {
	return &WalletController{Controller: service.NewController("WalletController")}
}

// Show runs the show action.
func (c *WalletController) Show(ctx *app.ShowWalletContext) error {
	// WalletController_Show: start_implement

	// Put your logic here

	res := &app.Wallet{}
	return ctx.OK(res)
	// WalletController_Show: end_implement
}
