package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
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

// List runs the list action.
func (c *TargetingController) List(ctx *app.ListTargetingContext) error {
	// TargetingController_List: start_implement

	// Put your logic here

	var r app.ReddituserCollection

	userSlice, err := c.db.GetRedditUsersAndSubs()
	if err != nil {
		log.Errorf("[controller/targeting] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not get list of reddit users and subs",
		})
	}

	for _, user := range userSlice {
		var s app.CommunityCollection
		for name, rep := range user.Social.Reddit.Subreddits {
			s = append(s, &app.Community{
				Name:       name,
				Reputation: rep,
			})
		}
		r = append(r, &app.Reddituser{
			UserID:     user.UserID,
			Subreddits: s,
		})
	}

	res := &app.Reddittargeting{
		Users: r,
	}
	return ctx.OK(res)
	// TargetingController_List: end_implement
}

// Set runs the set action.
func (c *TargetingController) Set(ctx *app.SetTargetingContext) error {
	// TargetingController_Set: start_implement

	// Put your logic here

	taskId := ctx.Payload.TaskID
	users := ctx.Payload.Users

	var userSlice []string

	err := json.Unmarshal([]byte(users), &userSlice)
	if err != nil {
		log.Errorf("[controller/targeting] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not unmarshal users string payload",
		})
	}

	for _, user := range userSlice {
		userTask := &types.UserTask{
			UserID: user,
			TaskID: taskId,
		}
		_, err := c.db.AddUserTask(userTask)
		if err != nil {
			log.Errorf("[controller/targeting] error: %v", err)
			return ctx.InternalServerError(&app.StandardError{
				Code:    500,
				Message: "could not assign task to user",
			})
		}
	}

	log.Printf("[controller/targeting] successfully assigned %v coindrop users to task id: %v\n", len(userSlice), taskId)

	return nil
	// TargetingController_Set: end_implement
}
