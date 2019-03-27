package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// WalletsController implements the wallet resource.
type WalletsController struct {
	*goa.Controller
	db *db.DB
}

// NewWalletsController creates a wallet controller.
func NewWalletsController(service *goa.Service, dbs *db.DB) *WalletsController {
	return &WalletsController{
		Controller: service.NewController("WalletsController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *WalletsController) Show(ctx *app.ShowWalletsContext) error {
	// WalletsController_Show: start_implement

	// Put your logic here
	userID := ctx.Value("authUserID").(string)

	// return a user's wallets using the AWS cognito user ID as the key
	wallets, err := c.db.GetWallets(userID)
	if err != nil {
		log.Errorf("[controller/wallet] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get wallets from db",
		})
	}

	var w app.WalletCollection

	for _, wallet := range wallets {
		w = append(w, &app.Wallet{
			Address:    wallet.Address,
			WalletType: wallet.Type,
		})
	}

	log.Printf("[controller/wallet] returned wallets for coindrop user: %v", userID)

	return ctx.OK(&app.Wallets{
		Wallets: w,
	})
	// WalletsController_Show: end_implement
}

// Update runs the update action.
func (c *WalletsController) Update(ctx *app.UpdateWalletsContext) error {
	// WalletsController_Update: start_implement

	// Put your logic here
	// userID := ctx.Value("authUserID").(string)

	wallet := &types.Wallet{
		Address: ctx.Payload.WalletAddress,
		Type:    ctx.Payload.WalletType,
	}

	walletStored := true

	// TODO:
	// needs better error handling

	_, err := c.db.GetWallet(userID, wallet.Type)
	if err != nil {
		walletStored = false
	}

	if walletStored {
		// update wallet in db
		_, err := c.db.UpdateWallet(
			userID,
			wallet.Address,
			wallet.Type,
		)
		if err != nil {
			log.Errorf("[controller/wallet] %v", err)
			return ctx.BadRequest(&app.StandardError{
				Code:    400,
				Message: "could not update wallet in db",
			})
		}
	} else {
		_, err := c.db.AddWallet(
			userID,
			wallet.Address,
			wallet.Type,
		)
		if err != nil {
			log.Errorf("[controller/wallet] %v", err)
			return ctx.BadRequest(&app.StandardError{
				Code:    400,
				Message: "could not insert wallet in db",
			})
		}
	}

	log.Printf("[controller/wallet] successfully updated wallet for coindrop user: %v\n", userID)

	res := &app.Wallet{
		Address:    wallet.Address,
		WalletType: wallet.Type,
	}

	return ctx.OK(res)
	// WalletsController_Update: end_implement
}
