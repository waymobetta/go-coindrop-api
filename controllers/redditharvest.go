package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/reddit"
	"github.com/waymobetta/go-coindrop-api/types"
)

// RedditharvestController implements the redditharvest resource.
type RedditharvestController struct {
	*goa.Controller
	db *db.DB
}

// NewRedditharvestController creates a redditharvest controller.
func NewRedditharvestController(service *goa.Service, dbs *db.DB) *RedditharvestController {
	return &RedditharvestController{
		Controller: service.NewController("RedditharvestController"),
		db:         dbs,
	}
}

// Update runs the update action.
func (c *RedditharvestController) Update(ctx *app.UpdateRedditharvestContext) error {
	// RedditharvestController_Update: start_implement

	// Put your logic here

	user := &types.User{
		Social: &types.Social{
			Reddit: &types.Reddit{
				Username:     ctx.Payload.Username,
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

	log.Println("[controllers/reddit] retrieving Reddit About info")
	// get general about info for user
	if err := authSession.GetAboutInfo(user); err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Reddit About info",
		})
	}

	log.Println("[controllers/reddit] retrieving Reddit Trophy info")
	// get list of trophies user has been awarded
	if err := authSession.GetRedditUserTrophies(user); err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Reddit Trophy info",
		})
	}

	log.Println("[controllers/reddit] retrieving Reddit Submitted info")
	// get slice of subreddits user is subscribed to based on activity
	if err := authSession.GetSubmittedInfo(user); err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Reddit Submitted info",
		})
	}

	user = &types.User{
		CognitoAuthUserID: ctx.Payload.UserID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				LinkKarma:    user.Social.Reddit.LinkKarma,
				CommentKarma: user.Social.Reddit.CommentKarma,
				Trophies:     user.Social.Reddit.Trophies,
				Subreddits:   user.Social.Reddit.Subreddits,
				Verification: &types.Verification{
					PostedVerificationCode:    "",
					ConfirmedVerificationCode: "",
					Verified:                  false,
				},
			},
		},
	}

	_, err = c.db.UpdateRedditInfo(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update create reddit info listing in db",
		})
	}

	res := &app.Reddituser{
		Username:     ctx.Payload.Username,
		LinkKarma:    user.Social.Reddit.LinkKarma,
		CommentKarma: user.Social.Reddit.CommentKarma,
		Trophies:     user.Social.Reddit.Trophies,
		Subreddits:   user.Social.Reddit.Subreddits,
	}
	return ctx.OK(res)
	// RedditharvestController_Update: end_implement
}
