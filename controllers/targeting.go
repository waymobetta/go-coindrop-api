package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// TargetingController implements the targeting resource.
type TargetingController struct {
	*goa.Controller
	db *db.DB
}

// NewTargetingController creates a targeting controller.
func NewTargetingController(service *goa.Service, dbs *db.DB) *TargetingController {
	return &TargetingController{
		Controller: service.NewController("TargetingController"),
		db:         dbs,
	}
}

// Display runs the display action.
func (c *TargetingController) Display(ctx *app.DisplayTargetingContext) error {
	// TargetingController_Display: start_implement

	// Put your logic here

	project := ctx.Params.Get("project")
	// threshold := ctx.Params.Get("threshold")

	threshold := 10

	users, err := c.db.GetEligibleRedditUsersAcrossSingleSub(project, threshold)
	if err != nil {
		log.Errorf("[controller/targeting] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get eligible users for subreddit",
		})
	}

	userBytes, err := json.Marshal(users)
	if err != nil {
		log.Errorf("[controller/targeting] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get marshall users into string",
		})
	}

	usersString := fmt.Sprintf("%s", userBytes)

	res := &app.Targeting{
		Users: usersString,
	}
	return ctx.OK(res)
	// TargetingController_Display: end_implement
}

// Set runs the set action.
func (c *TargetingController) Set(ctx *app.SetTargetingContext) error {
	// TargetingController_Set: start_implement

	// Put your logic here

	// POST
	type setTargetingPayload struct {
		// Task ID
		TaskID *string `form:"taskId,omitempty" json:"taskId,omitempty" yaml:"taskId,omitempty" xml:"taskId,omitempty"`
		// List of users
		Users *string `form:"users,omitempty" json:"users,omitempty" yaml:"users,omitempty" xml:"users,omitempty"`
	}

	return nil
	// TargetingController_Set: end_implement
}
