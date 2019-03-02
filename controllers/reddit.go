package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
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

	user := &db.User2{
		Social: &db.Social{
			Reddit: &db.Reddit{
				Verification: &verify.Verification2{},
			},
		}}
	user.AuthUserID = ctx.Params.Get("userId")

	_, err := c.db.GetRedditUser2(user)
	if err != nil {
		log.Errorf("[controller/reddit] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get reddit user info from db",
		})
	}

	res := &app.Reddituser{
		Username:          user.Social.Reddit.Username,
		LinkKarma:         user.Social.Reddit.LinkKarma,
		CommentKarma:      user.Social.Reddit.CommentKarma,
		Subreddits:        user.Social.Reddit.Subreddits,
		Trophies:          user.Social.Reddit.Trophies,
		AccountCreatedUTC: user.Social.Reddit.AccountCreatedUTC,
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
	// RedditController_Update: start_implement

	// Put your logic here

	// user := new(db.User2)
	// user.AuthUserID = ctx.Params.Get("userId")

	// _, err := c.db.AddRedditUser2(user)
	// if err != nil {
	// 	log.Errorf("[controller/reddit] %v", err)
	// 	return ctx.NotFound(&app.StandardError{
	// 		Code:    400,
	// 		Message: "could not get reddit user info from db",
	// 	})
	// }

	// fmt.Println(user)

	// res := &app.Reddituser{
	// 	Username:          user.Social.Reddit.Username,
	// 	LinkKarma:         user.Social.Reddit.LinkKarma,
	// 	CommentKarma:      user.Social.Reddit.CommentKarma,
	// 	Subreddits:        user.Social.Reddit.Subreddits,
	// 	Trophies:          user.Social.Reddit.Trophies,
	// 	AccountCreatedUTC: user.Social.Reddit.AccountCreatedUTC,
	// 	Verification: &app.Verification{
	// 		PostedVerificationCode:    user.Social.Reddit.Verification.PostedVerificationCode,
	// 		ConfirmedVerificationCode: user.Social.Reddit.Verification.ConfirmedVerificationCode,
	// 		Verified:                  user.Social.Reddit.Verification.Verified,
	// 	},
	// }

	res := &app.Reddituser{}
	return ctx.OK(res)
	// RedditController_Update: end_implement
}
