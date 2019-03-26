package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
	"google.golang.org/appengine/log"
)

// BadgesController implements the badges resource.
type BadgesController struct {
	*goa.Controller
	db *db.DB
}

// NewBadgesController creates a badges controller.
func NewBadgesController(service *goa.Service, dbs *db.DB) *BadgesController {
	return &BadgesController{
		Controller: service.NewController("BadgesController"),
		db:         dbs,
	}
}

// Create runs the create action.
func (c *BadgesController) Create(ctx *app.CreateBadgesContext) error {
	// BadgesController_Create: start_implement

	// Put your logic here

	badge := &types.Badge{
		Name:        ctx.Payload.Name,
		Description: ctx.Payload.Description,
	}

	err := c.db.AddBadge(badge)
	if err != nil {
		log.Errorf("[controller/badges] failed to add badge to db: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not add badge to db",
		})
	}

	res := &app.Badge{
		ID:          badge.ID,
		Name:        badge.Name,
		Description: badge.Description,
	}
	return ctx.OK(res)
	// BadgesController_Create: end_implement
}

// List runs the list action.
// Returns all badges for specific user
func (c *BadgesController) List(ctx *app.ListBadgesContext) error {
	// BadgesController_List: start_implement

	// Put your logic here
	userID := ctx.Params.Get("userId")
	// Note: if query string `userId` is empty,
	// then use user ID from auth token
	if userID == "" {
		userID = ctx.Value("authUserID").(string)
	}

	badges, err := c.db.GetUserBadges(userID)
	if err != nil {
		log.Errorf("[controller/badges] failed to get user badges: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get user's badges",
		})
	}

	var b app.BadgeCollection

	for _, badge := range badges {
		b = append(b, &app.Badge{
			ID:          badge.ID,
			Name:        badge.Name,
			Description: badge.Description,
		})
	}

	log.Printf("[controller/badges] returned badges for coindrop user: %v\n", userID)

	res := &app.Badges{
		Badges: b,
	}
	return ctx.OK(res)
	// BadgesController_List: end_implement
}

// Show runs the show action.
// Returns all badges
func (c *BadgesController) Show(ctx *app.ShowBadgesContext) error {
	// BadgesController_Show: start_implement

	// Put your logic here

	badges, err := c.db.GetBadges()
	if err != nil {
		log.Errorf("[controller/badges] failed to get badges: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get badges",
		})
	}

	var badgeCollection app.BadgeCollection

	for _, badge := range badges {
		badgeCollection = append(badgeCollection, &app.Badge{
			ID:          badge.ID,
			Name:        badge.Name,
			Description: badge.Description,
		})
	}

	return ctx.OK(badgeCollection)
	// BadgesController_Show: end_implement
}
