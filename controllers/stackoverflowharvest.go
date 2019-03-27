package controllers

import (
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

// Update runs the update action.
func (c *StackoverflowharvestController) Update(ctx *app.UpdateStackoverflowharvestContext) error {
	// StackoverflowharvestController_Update: start_implement

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
			Message: "could not retrieve Stack Overflow About info",
		})
	}

	// set accounts as empty slice
	user.Social.StackOverflow.Accounts = []string{}

	log.Println("[controller/stackoverflowharvest] retrieving Stack Overflow associated accounts info")
	// get list of trophies user has been awarded
	err = stackoverflow.GetAssociatedAccounts(user)
	if err != nil {
		log.Errorf("[controller/stackoverflowharvest] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not retrieve Stack Overflow Associated Account info",
		})
	}

	user = &types.User{
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				DisplayName:       aboutProfile.Items[0].DisplayName,
				ExchangeAccountID: aboutProfile.Items[0].AccountID,
				Accounts:          user.Social.StackOverflow.Accounts,
				// Communities:       user.Social.StackOverflow.Communities,
				Verification: &types.Verification{
					PostedVerificationCode:    "",
					ConfirmedVerificationCode: "",
					Verified:                  false,
				},
			},
		},
	}

	log.Println("[controller/stackoverflowharvest] retrieving Stack Overflow associated accounts info")
	_, err = c.db.UpdateStackAboutInfo(user)
	if err != nil {
		log.Errorf("[controller/stackoverflowharvest] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update Stack Overflow user listing in db",
		})
	}

	log.Println("[controller/stackoverflowharvest] retrieving Stack Overflow associated accounts info")

	res := &app.Stackoverflowuser{
		DisplayName:       user.Social.StackOverflow.DisplayName,
		ExchangeAccountID: user.Social.StackOverflow.ExchangeAccountID,
		Accounts:          user.Social.StackOverflow.Accounts,
	}
	return ctx.OK(res)
	// StackoverflowharvestController_Update: end_implement
}
