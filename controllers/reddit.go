package controllers

import (
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
	"github.com/waymobetta/go-coindrop-api/verify"
)

// RedditController implements the reddit resource.
type RedditController struct {
	*goa.Controller
	db *db.DB
}

// NewRedditController creates a reddit controller.
func NewRedditController(service *goa.Service, dbs *db.DB) *RedditController {
	return &RedditController{
		Controller: service.NewController("RedditController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *RedditController) Show(ctx *app.ShowRedditContext) error {
	// RedditController_Show: start_implement

	// Put your logic here

	user := &types.User{
		Social: &types.Social{
			Reddit: &types.Reddit{
				Verification: &verify.Verification2{},
			},
		}}
	user.CognitoAuthUserID = ctx.Params.Get("userId")

	_, err := c.db.GetRedditUser(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user reddit info from db",
		})
	}

	res := &app.Reddituser{
		Username:     user.Social.Reddit.Username,
		LinkKarma:    user.Social.Reddit.LinkKarma,
		CommentKarma: user.Social.Reddit.CommentKarma,
		Subreddits:   user.Social.Reddit.Subreddits,
		Trophies:     user.Social.Reddit.Trophies,
		Verification: &app.Verification{
			PostedVerificationCode:    user.Social.Reddit.Verification.PostedVerificationCode,
			ConfirmedVerificationCode: user.Social.Reddit.Verification.ConfirmedVerificationCode,
			Verified:                  user.Social.Reddit.Verification.Verified,
		},
	}
	return ctx.OK(res)
	// RedditController_Show: end_implement
}

// Update runs the update action.
func (c *RedditController) Create(ctx *app.CreateRedditContext) error {
	// RedditController_Create: start_implement

	// Put your logic here

	// TODO:
	// 1. fix to prevent creating duplicates
	// 2. fix SQL statement to join auth table for user ID

	user := &types.User{
		CognitoAuthUserID: ctx.Payload.UserID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				Username:     ctx.Payload.Username,
				LinkKarma:    0,
				CommentKarma: 0,
				Subreddits:   []string{},
				Trophies:     []string{},
				Verification: &verify.Verification2{
					PostedVerificationCode:    "",
					ConfirmedVerificationCode: "",
					Verified:                  false,
				},
			},
		},
	}

	user.UserID = "a1f75b14-a475-4dee-a1ea-bc4c0d391e7e"

	_, err := c.db.AddRedditUser(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update create reddit info listing in db",
		})
	}

	fmt.Println(user)

	res := &app.Reddituser{}

	return ctx.OK(res)
	// RedditController_Update: end_implement
}
