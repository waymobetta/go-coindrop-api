package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/reddit"
	"github.com/waymobetta/go-coindrop-api/types"
)

// VerifyredditController implements the verifyreddit resource.
type VerifyredditController struct {
	*goa.Controller
	db *db.DB
}

// NewVerifyredditController creates a verifyreddit controller.
func NewVerifyredditController(service *goa.Service, dbs *db.DB) *VerifyredditController {
	return &VerifyredditController{
		Controller: service.NewController("VerifyredditController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *VerifyredditController) Show(ctx *app.ShowVerifyredditContext) error {
	// VerifyredditController_Show: start_implement

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
		log.Errorf("[controller/verifyreddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	log.Printf("[controller/verifyreddit] returned verification information for coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.Reddituser{
		Verification: &app.Verification{
			PostedVerificationCode:    user.Social.Reddit.Verification.PostedVerificationCode,
			ConfirmedVerificationCode: user.Social.Reddit.Verification.ConfirmedVerificationCode,
			Verified:                  user.Social.Reddit.Verification.Verified,
		},
	}
	return ctx.OK(res)
	// VerifyredditController_Show: end_implement
}

// Update runs the update action.
func (c *VerifyredditController) Update(ctx *app.UpdateVerifyredditContext) error {
	// VerifyredditController_Update: start_implement

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
		log.Errorf("[controller/verifyreddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	// retrieve verification code from coindrop verification subreddit
	err = authSession.GetRecentPostsFromSubreddit(user)
	if err != nil {
		log.Errorf("[controller/verifyreddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "verification code does not match",
		})
	}

	// update verification code for user in db
	_, err = c.db.UpdateRedditVerificationCode(user)
	if err != nil {
		log.Errorf("[controller/verifyreddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	log.Printf("[controller/verifyreddit] successfully verified reddit account for coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.Reddituser{
		Verification: &app.Verification{
			Verified: user.Social.Reddit.Verification.Verified,
		},
	}

	return ctx.OK(res)
	// VerifyredditController_Update: end_implement
}
