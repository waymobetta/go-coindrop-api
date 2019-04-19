package controllers

import (
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

	userID := ctx.Params.Get("userId")
	if userID == "" {
		userID = ctx.Value("authUserID").(string)
	}

	user := &types.User{
		UserID: userID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				Verification: &types.Verification{},
			},
		},
	}

	_, err := c.db.GetRedditUser(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user reddit info from db",
		})
	}

	var communityCollection app.CommunityCollection

	for name, rep := range user.Social.Reddit.Subreddits {
		communityCollection = append(communityCollection, &app.Community{
			Name:       name,
			Reputation: rep,
		})
	}

	res := &app.Reddituser{
		ID:           user.Social.Reddit.ID,
		Username:     user.Social.Reddit.Username,
		LinkKarma:    user.Social.Reddit.LinkKarma,
		CommentKarma: user.Social.Reddit.CommentKarma,
		Subreddits:   communityCollection,
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

	user := &types.User{
		UserID: ctx.Payload.UserID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				Username:     ctx.Payload.Username,
				LinkKarma:    0,
				CommentKarma: 0,
				Subreddits:   map[string]int{},
				Trophies:     []string{},
				Verification: &types.Verification{
					PostedVerificationCode:    "",
					ConfirmedVerificationCode: "",
					Verified:                  false,
				},
			},
		},
	}

	user, err := c.db.AddRedditUser(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not add user's reddit listing to db",
		})
	}

	var communityCollection app.CommunityCollection

	log.Printf("[controller/reddit] added Reddit information for coindrop user: %v\n", user.UserID)

	res := &app.Reddituser{
		UserID:       user.UserID,
		Username:     user.Social.Reddit.Username,
		LinkKarma:    user.Social.Reddit.LinkKarma,
		CommentKarma: user.Social.Reddit.CommentKarma,
		Trophies:     []string{},
		Subreddits:   communityCollection,
		Verification: &app.Verification{
			PostedVerificationCode:    "",
			ConfirmedVerificationCode: "",
			Verified:                  false,
		},
	}

	return ctx.OK(res)
	// RedditController_Update: end_implement
}

// Display runs the display action.
func (c *RedditController) Display(ctx *app.DisplayRedditContext) error {
	// RedditController_Display: start_implement

	// Put your logic here

	userID := ctx.Params.Get("userId")
	if userID == "" {
		userID = ctx.Value("authUserID").(string)
	}

	user := &types.User{
		UserID: userID,
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

	log.Printf("[controller/reddit] returned verification information for coindrop user: %v\n", user.UserID)

	res := &app.Verification{
		PostedVerificationCode:    user.Social.Reddit.Verification.PostedVerificationCode,
		ConfirmedVerificationCode: user.Social.Reddit.Verification.ConfirmedVerificationCode,
		Verified:                  user.Social.Reddit.Verification.Verified,
	}

	return ctx.OK(res)
	// RedditController_Display: end_implement
}

// Verify runs the verify action.
func (c *RedditController) Verify(ctx *app.VerifyRedditContext) error {
	// RedditController_Verify: start_implement

	// Put your logic here

	user := &types.User{
		UserID: ctx.Payload.UserID,
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
			Message: "could not update user verification info in db",
		})
	}

	log.Printf("[controller/reddit] successfully verified reddit account for coindrop user: %v\n", user.UserID)

	res := &app.Verification{
		PostedVerificationCode:    user.Social.Reddit.Verification.PostedVerificationCode,
		ConfirmedVerificationCode: user.Social.Reddit.Verification.ConfirmedVerificationCode,
		Verified:                  user.Social.Reddit.Verification.Verified,
	}

	return ctx.OK(res)
	// RedditController_Verify: end_implement
}
