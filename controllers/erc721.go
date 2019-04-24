package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	ethsvc "github.com/waymobetta/go-coindrop-api/services/ethereum"
	"google.golang.org/appengine/log"
)

// Erc721Controller implements the erc721 resource.
type Erc721Controller struct {
	*goa.Controller
	db *db.DB
}

// NewErc721Controller creates a erc721 controller.
func NewErc721Controller(service *goa.Service, dbs *db.DB) *Erc721Controller {
	return &Erc721Controller{
		Controller: service.NewController("Erc721Controller"),
		db:         dbs,
	}
}

// Assign runs the assign action.
func (c *Erc721Controller) Assign(ctx *app.AssignErc721Context) error {
	// Erc721Controller_Assign: start_implement

	// Put your logic here

	badgeId := ctx.Payload.BadgeID
	userId := ctx.Payload.UserID

	tokenId, err := ethsvc.MintToken(badgeId)
	if err != nil {
		log.Errorf("[controller/erc721] failed to mint ERC721: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not mint ERC721",
		})
	}

	err = c.db.AssignERC721ToUser(
		tokenId,
		badgeId,
		userId,
	)
	if err != nil {
		log.Errorf("[controller/erc721] failed to assign ERC721 to user in db: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not assign ERC721 to user in db",
		})
	}

	res := &app.Erc721{
		ContractAddress: "",
		TokenID:         tokenId,
		TotalMinted:     0,
	}
	return ctx.OK(res)
	// Erc721Controller_Assign: end_implement
}
