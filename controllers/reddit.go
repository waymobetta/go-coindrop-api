package controllers

import (
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/reddit"
	"github.com/waymobetta/go-coindrop-api/types"
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
				Verification: &types.Verification{},
			},
		},
	}

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
func (c *RedditController) Update(ctx *app.UpdateRedditContext) error {
	// RedditController_Create: start_implement

	// Put your logic here

	// TODO:
	// 1. fix to prevent creating duplicates
	// 2. fix SQL statement to join auth table for user ID

	user := &types.User{
		CognitoAuthUserID: ctx.Value("authUserID").(string),
		Social: &types.Social{
			Reddit: &types.Reddit{
				Username:     ctx.Payload.Username,
				LinkKarma:    0,
				CommentKarma: 0,
				Subreddits:   []string{},
				Trophies:     []string{},
				Verification: &types.Verification{
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

// Display runs the display action.
func (c *RedditController) Display(ctx *app.DisplayRedditContext) error {
	// RedditController_Display: start_implement

	// Put your logic here

	user := &types.User{
		CognitoAuthUserID: ctx.Params.Get("userId"),
		Social: &types.Social{
			Reddit: &types.Reddit{
				Verification: &types.Verification{},
			},
		},
	}

	_, err := c.db.GetUserRedditVerification(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	log.Printf("[controller/reddit] returned verification information for coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.Reddituser{
		Verification: &app.Verification{
			PostedVerificationCode:    user.Social.Reddit.Verification.PostedVerificationCode,
			ConfirmedVerificationCode: user.Social.Reddit.Verification.ConfirmedVerificationCode,
			Verified:                  user.Social.Reddit.Verification.Verified,
		},
	}

	return ctx.OK(res)
	// RedditController_Display: end_implement
}

// Verify runs the verify action.
func (c *RedditController) Verify(ctx *app.VerifyRedditContext) error {
	// RedditController_Verify: start_implement

	// Put your logic here

	user := &types.User{
		CognitoAuthUserID: ctx.Payload.UserID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				Verification: &types.Verification{},
			},
		},
	}

	// initializes reddit OAuth sessions
	authSession, err := reddit.NewRedditAuth()
	if err != nil {
		log.Errorf("[controller/reddit] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not start reddit auth session",
		})
	}

	// get previously stored verification info from db
	_, err = c.db.GetUserRedditVerification(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	// retrieve verification code from coindrop verification subreddit
	err = authSession.GetRecentPostsFromSubreddit(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "verification code does not match",
		})
	}

	// update verification code for user in db
	_, err = c.db.UpdateRedditVerificationCode(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	log.Printf("[controller/reddit] successfully verified reddit account for coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.Reddituser{
		Verification: &app.Verification{
			Verified: user.Social.Reddit.Verification.Verified,
		},
	}

	return ctx.OK(res)
	// RedditController_Verify: end_implement
}
