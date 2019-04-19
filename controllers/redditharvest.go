package controllers

import (
	"encoding/json"
	"fmt"

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

// UpdateAbout runs the updateAbout action.
func (c *RedditharvestController) UpdateAbout(ctx *app.UpdateAboutRedditharvestContext) error {
	// RedditharvestController_UpdateAbout: start_implement

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
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not start reddit auth session",
		})
	}

	log.Println("[controllers/reddit/harvest] retrieving Reddit About info")
	// get general about info for user
	if err := authSession.GetAboutInfo(user); err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Reddit About info",
		})
	}

	user = &types.User{
		UserID: ctx.Payload.UserID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				LinkKarma:    user.Social.Reddit.LinkKarma,
				CommentKarma: user.Social.Reddit.CommentKarma,
			},
		},
	}

	_, err = c.db.UpdateRedditKarmaInfo(user)
	if err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update Reddit user listing in db",
		})
	}

	log.Printf("[controllers/reddit/harvest] successfully harvested Reddit About info for user: %s", user.UserID)

	res := &app.Reddituser{
		Username:     ctx.Payload.Username,
		LinkKarma:    user.Social.Reddit.LinkKarma,
		CommentKarma: user.Social.Reddit.CommentKarma,
	}
	return ctx.OK(res)
	// RedditharvestController_UpdateAbout: end_implement
}

// UpdateTrophies runs the updateTrophies action.
func (c *RedditharvestController) UpdateTrophies(ctx *app.UpdateTrophiesRedditharvestContext) error {
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
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not start reddit auth session",
		})
	}

	log.Println("[controllers/reddit/harvest] retrieving Reddit Trophy info")
	// get list of trophies user has been awarded
	if err := authSession.GetRedditUserTrophies(user); err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Reddit Trophy info",
		})
	}

	user = &types.User{
		UserID: ctx.Payload.UserID,
		Social: &types.Social{
			Reddit: &types.Reddit{
				Trophies: user.Social.Reddit.Trophies,
			},
		},
	}

	_, err = c.db.UpdateRedditTrophyInfo(user)
	if err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update Reddit user listing in db",
		})
	}

	log.Printf("[controllers/reddit/harvest] successfully harvested Reddit Trophy info for user: %s", user.UserID)

	res := &app.Reddituser{
		Username: ctx.Payload.Username,
		Trophies: user.Social.Reddit.Trophies,
	}
	return ctx.OK(res)
	// RedditharvestController_UpdateTrophies: end_implement
}

// UpdateSubmittedInfo runs the updateSubmittedInfo action.
func (c *RedditharvestController) UpdateSubmittedInfo(ctx *app.UpdateSubmittedInfoRedditharvestContext) error {
	// RedditharvestController_UpdateSubmittedInfo: start_implement

	// Put your logic here

	user := &types.User{
		UserID: ctx.Payload.UserID,
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
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not start reddit auth session",
		})
	}

	log.Println("[controllers/reddit/harvest] retrieving Reddit Submitted info")
	// get slice of subreddits user is subscribed to based on activity
	subs, err := authSession.GetRawSubmittedInfo(user)
	if err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Reddit Submitted info",
		})
	}

	specMap := make(map[string]int)

	for _, sub := range subs {
		specMap[sub.Subreddit] = sub.Ups * 10
	}

	mapString, err := json.Marshal(specMap)
	if err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "error marshalling community map",
		})
	}

	specMapString := fmt.Sprintf("%s", mapString)

	var communityCollection app.CommunityCollection

	for name, rep := range specMap {
		communityCollection = append(communityCollection, &app.Community{
			Name:       name,
			Reputation: rep,
		})
	}

	err = c.db.UpdateRedditSubInfo(specMapString, user.UserID)
	if err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update Reddit user listing in db",
		})
	}

	log.Printf("[controllers/reddit/harvest] successfully harvested Reddit Submitted info for user: %s", user.UserID)

	res := &app.Reddituser{
		Username:   ctx.Payload.Username,
		Subreddits: communityCollection,
	}
	return ctx.OK(res)
	// RedditharvestController_UpdateSubmittedInfo: end_implement
}
