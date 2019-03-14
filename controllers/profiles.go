package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// ProfilesController implements the profiles resource.
type ProfilesController struct {
	*goa.Controller
	db *db.DB
}

// NewProfilesController creates a profiles controller.
func NewProfilesController(service *goa.Service, dbs *db.DB) *ProfilesController {
	return &ProfilesController{
		Controller: service.NewController("ProfilesController"),
		db:         dbs,
	}
}

// Create runs the create action.
func (c *ProfilesController) Create(ctx *app.CreateProfilesContext) error {
	// ProfilesController_Create: start_implement

	// Put your logic here
	userID := ctx.Value("authUserID").(string)
	p := new(types.Profile)
	p.UserID = userID
	if ctx.Payload.Name != "" {
		p.Name = ctx.Payload.Name
	}
	if ctx.Payload.Username != "" {
		p.Username = ctx.Payload.Username
	}
	p, err := c.db.UpsertProfile(p)
	if err != nil {
		log.Errorf("[controller/profiles] failed to upsert profile: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not upsert profile",
		})
	}

	return ctx.OK(&app.Profile{
		Name:     p.Name,
		Username: p.Username,
	})
	// ProfilesController_Create: end_implement
}

// List runs the list action.
func (c *ProfilesController) List(ctx *app.ListProfilesContext) error {
	// ProfilesController_List: start_implement

	// Put your logic here
	userID := ctx.Value("authUserID").(string)
	p, err := c.db.GetProfile(userID)
	if err != nil {
		log.Errorf("[controller/profiles] failed to retrieve profile: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve profile",
		})
	}

	return ctx.OK(&app.Profile{
		Name:     p.Name,
		Username: p.Username,
	})
	// ProfilesController_List: end_implement
}

// Show runs the show action.
func (c *ProfilesController) Show(ctx *app.ShowProfilesContext) error {
	// ProfilesController_Show: start_implement

	// Put your logic here
	userID := ctx.Value("authUserID").(string)
	p, err := c.db.GetProfile(userID)
	if err != nil {
		log.Errorf("[controller/profiles] failed to retrieve profile: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve profile",
		})
	}

	return ctx.OK(&app.Profile{
		Name:     p.Name,
		Username: p.Username,
	})
	// ProfilesController_Show: end_implement
}

// Update runs the update action.
func (c *ProfilesController) Update(ctx *app.UpdateProfilesContext) error {
	// ProfilesController_Update: start_implement

	// Put your logic here
	userID := ctx.Value("authUserID").(string)
	p := new(types.Profile)
	p.UserID = userID
	if ctx.Payload.Name != "" {
		p.Name = ctx.Payload.Name
	}
	if ctx.Payload.Username != "" {
		p.Username = ctx.Payload.Username
	}
	p, err := c.db.UpdateProfile(p)
	if err != nil {
		log.Errorf("[controller/profiles] failed to update profile: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not update profile",
		})
	}

	return ctx.OK(&app.Profile{
		Name:     p.Name,
		Username: p.Username,
	})
	// ProfilesController_Update: end_implement
}
