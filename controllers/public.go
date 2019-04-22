package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
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

// Show runs the show action.
func (c *PublicController) Show(ctx *app.ShowPublicContext) error {
	// PublicController_Show: start_implement

	// Put your logic here

	redditUsername := ctx.RedditUsername

	var badgeCollection app.BadgeCollection

	badges, err := c.db.GetBadgesByRedditUsername(redditUsername)
	if err != nil {
		log.Errorf("[controller/public] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get badges from the supplied reddit username",
		})
	}

	for _, badge := range badges {
		badgeCollection = append(badgeCollection, &app.Badge{
			ID: badge.ID,
		})
	}

	res := &app.Public{
		Badges: badgeCollection,
	}
	return ctx.OK(res)
	// PublicController_Show: end_implement
}
