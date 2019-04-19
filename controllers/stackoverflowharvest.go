package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/stackoverflow"
	"github.com/waymobetta/go-coindrop-api/types"
)

// StackoverflowharvestController implements the stackoverflowharvest resource.
type StackoverflowharvestController struct {
	*goa.Controller
	db *db.DB
}

// NewStackoverflowharvestController creates a stackoverflowharvest controller.
func NewStackoverflowharvestController(service *goa.Service, dbs *db.DB) *StackoverflowharvestController {
	return &StackoverflowharvestController{
		Controller: service.NewController("StackoverflowharvestController"),
		db:         dbs,
	}
}

// UpdateProfile runs the updateProfile action.
func (c *StackoverflowharvestController) UpdateProfile(ctx *app.UpdateProfileStackoverflowharvestContext) error {
	// StackoverflowharvestController_UpdateProfile: start_implement

	// Put your logic here
	user := &types.User{
		UserID: ctx.Payload.UserID,
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				StackUserID:  ctx.Payload.StackUserID,
				Verification: &types.Verification{},
			},
		},
	}

	log.Println("[controller/stackoverflowharvest] retrieving Stack Overflow About info")
	// get general about info for user

	aboutProfile, err := stackoverflow.GetProfileByUserID(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Stack Overflow Profile info",
		})
	}

	user.Social.StackOverflow.DisplayName = aboutProfile.Items[0].DisplayName
	user.Social.StackOverflow.ExchangeAccountID = aboutProfile.Items[0].AccountID

	err = c.db.UpdateStackProfileInfo(user)
	if err != nil {
		log.Errorf("[controller/stackoverflowharvest] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update Stack Overflow user listing in db",
		})
	}

	log.Println("[controller/stackoverflowharvest] successfully retrieved Stack Overflow profile info")

	res := &app.Stackoverflowuser{
		DisplayName:       user.Social.StackOverflow.DisplayName,
		ExchangeAccountID: user.Social.StackOverflow.ExchangeAccountID,
		StackUserID:       user.Social.StackOverflow.StackUserID,
	}
	return ctx.OK(res)
	// StackoverflowharvestController_UpdateProfile: end_implement
}

// UpdateCommunities runs the updateCommunities action.
func (c *StackoverflowharvestController) UpdateCommunities(ctx *app.UpdateCommunitiesStackoverflowharvestContext) error {
	// StackoverflowharvestController_UpdateCommunities: start_implement

	// Put your logic here

	user := &types.User{
		UserID: ctx.Payload.UserID,
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				StackUserID:  ctx.Payload.StackUserID,
				Verification: &types.Verification{},
			},
		},
	}

	log.Println("[controller/stackoverflowharvest] retrieving Stack Overflow exchange account id")
	stackExchangeId, err := c.db.GetExchangeIdByStackId(
		user.Social.StackOverflow.StackUserID,
		user.UserID,
	)
	if err != nil {
		log.Errorf("[controller/stackoverflowharvest] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Stack Overflow Associated Account info",
		})
	}

	user.Social.StackOverflow.ExchangeAccountID = stackExchangeId

	log.Println("[controller/stackoverflowharvest] retrieving Stack Overflow associated accounts info")
	// get list of trophies user has been awarded
	accounts, err := stackoverflow.GetAssociatedAccounts(user)
	if err != nil {
		log.Errorf("[controller/stackoverflowharvest] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Stack Overflow Associated Account info",
		})
	}

	communityMap := make(map[string]int)

	for _, community := range accounts {
		communityMap[community.Name] = community.Reputation
	}

	mapString, err := json.Marshal(communityMap)
	if err != nil {
		log.Errorf("[controller/reddit/harvest] error: %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "error marshalling community map",
		})
	}

	communityMapString := fmt.Sprintf("%s", mapString)

	var communityCollection app.CommunityCollection

	for _, community := range accounts {
		communityCollection = append(communityCollection, &app.Community{
			Name:       community.Name,
			Reputation: community.Reputation,
		})
	}

	err = c.db.UpdateStackCommunityInfo(communityMapString, user.UserID)
	if err != nil {
		log.Errorf("[controller/stackoverflowharvest] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update Stack Overflow user listing in db",
		})
	}

	log.Println("[controller/stackoverflowharvest] successfully retrieved Stack Overflow associated accounts info")

	res := &app.Stackoverflowuser{
		StackUserID: user.Social.StackOverflow.StackUserID,
		Accounts:    communityCollection,
	}
	return ctx.OK(res)
	// StackoverflowharvestController_UpdateCommunities: end_implement
}
