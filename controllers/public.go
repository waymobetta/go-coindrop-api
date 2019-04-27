package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// PublicController implements the public resource.
type PublicController struct {
	*goa.Controller
	db *db.DB
}

// NewPublicController creates a public controller.
func NewPublicController(service *goa.Service, dbs *db.DB) *PublicController {
	return &PublicController{
		Controller: service.NewController("PublicController"),
		db:         dbs,
	}
}

// Display runs the display action.
func (c *PublicController) Display(ctx *app.DisplayPublicContext) error {
	// PublicController_Display: start_implement

	// Put your logic here

	tokenId := ctx.Erc721TokenID

	erc721Lookup, err := c.db.GetTaskAndBadgeBy721Id(tokenId)
	if err != nil {
		log.Errorf("[controller/public] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get task + badge info from the supplied ERC-721 token ID",
		})
	}

	log.Printf("[controller/public] successfully returned task + badge info for ERC-721 token ID: %s", tokenId)

	res := &app.Erc721lookup{
		Task: &app.Task{
			Author:      erc721Lookup.Task.Author,
			Title:       erc721Lookup.Task.Title,
			Type:        erc721Lookup.Task.Type,
			Description: erc721Lookup.Task.Description,
			LogoURL:     erc721Lookup.Task.LogoURL,
			Badge: &app.Badge{
				Name:        erc721Lookup.Task.BadgeData.Name,
				Description: erc721Lookup.Task.BadgeData.Description,
				LogoURL:     erc721Lookup.Task.BadgeData.LogoURL,
			},
		},
		Erc721: &app.Erc721{
			TokenID:         tokenId,
			ContractAddress: erc721Lookup.ERC721.ContractAddress,
		},
	}
	return ctx.OK(res)
	// PublicController_Display: end_implement
}

// Show runs the show action.
func (c *PublicController) Show(ctx *app.ShowPublicContext) error {
	// PublicController_Show: start_implement

	// Put your logic here

	publicUser := &types.Public{
		RedditUsername: ctx.RedditUsername,
	}

	var badgeCollection app.BadgeCollection

	badges, err := c.db.GetBadgesByRedditUsername(publicUser.RedditUsername)
	if err != nil {
		log.Errorf("[controller/public] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get badges from the supplied reddit username",
		})
	}

	for _, badge := range badges {
		badgeCollection = append(badgeCollection, &app.Badge{
			Name:        badge.Name,
			Project:     badge.Project,
			Description: badge.Description,
			LogoURL:     badge.LogoURL,
		})
	}

	log.Printf("[controller/public] successfully returned badges for Reddit user: %v\n", publicUser.RedditUsername)

	res := &app.Publicbadges{
		RedditUsername: publicUser.RedditUsername,
		Badges:         badgeCollection,
	}
	return ctx.OK(res)
	// PublicController_Show: end_implement
}
