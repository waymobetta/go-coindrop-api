package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/services/stackoverflow"
	"github.com/waymobetta/go-coindrop-api/types"
)

var (
	noVerifError = errors.New("verification code does not match")
)

// StackoverflowController implements the stackoverflow resource.
type StackoverflowController struct {
	*goa.Controller
	db *db.DB
}

// NewStackoverflowController creates a stackoverflow controller.
func NewStackoverflowController(service *goa.Service, dbs *db.DB) *StackoverflowController {
	return &StackoverflowController{
		Controller: service.NewController("StackoverflowController"),
		db:         dbs,
	}
}

// Display runs the display action.
func (c *StackoverflowController) Display(ctx *app.DisplayStackoverflowContext) error {
	// StackoverflowController_Display: start_implement

	// Put your logic here

	user := &types.User{
		CognitoAuthUserID: ctx.Params.Get("userId"),
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				Verification: &types.Verification{},
			},
		},
	}

	user.CognitoAuthUserID = ctx.Params.Get("userId")

	_, err := c.db.GetUserStackOverfloVerification(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	log.Printf("[controller/stackoverflow] returned verification information for coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.Stackoverflowuser{
		Verification: &app.Verification{
			PostedVerificationCode:    user.Social.StackOverflow.Verification.PostedVerificationCode,
			ConfirmedVerificationCode: user.Social.StackOverflow.Verification.ConfirmedVerificationCode,
			Verified:                  user.Social.StackOverflow.Verification.Verified,
		},
	}
	return ctx.OK(res)
	// StackoverflowController_Display: end_implement
}

// Show runs the show action.
func (c *StackoverflowController) Show(ctx *app.ShowStackoverflowContext) error {
	// StackoverflowController_Show: start_implement

	// Put your logic here

	user := &types.User{
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				Verification: &types.Verification{},
			},
		},
	}

	user.CognitoAuthUserID = ctx.Params.Get("userId")

	_, err := c.db.AddStackUser(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user stack overflow info from db",
		})
	}

	res := &app.Stackoverflowuser{
		DisplayName:       user.Social.StackOverflow.DisplayName,
		StackUserID:       user.Social.StackOverflow.StackUserID,
		ExchangeAccountID: user.Social.StackOverflow.ExchangeAccountID,
		Accounts:          user.Social.StackOverflow.Accounts,
		Verification: &app.Verification{
			PostedVerificationCode:    user.Social.StackOverflow.Verification.PostedVerificationCode,
			ConfirmedVerificationCode: user.Social.StackOverflow.Verification.ConfirmedVerificationCode,
			Verified:                  user.Social.StackOverflow.Verification.Verified,
		},
	}
	return ctx.OK(res)
	// StackoverflowController_Show: end_implement
}

// Update runs the update action.
func (c *StackoverflowController) Update(ctx *app.UpdateStackoverflowContext) error {
	// StackoverflowController_Update: start_implement

	// Put your logic here

	user := &types.User{
		CognitoAuthUserID: ctx.Value("authUserID").(string),
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				StackUserID:       0,
				ExchangeAccountID: 0,
				DisplayName:       "",
				Accounts:          []string{},
				Verification: &types.Verification{
					PostedVerificationCode:    "",
					ConfirmedVerificationCode: "",
					Verified:                  false,
				},
			},
		},
	}

	_, err := c.db.UpdateStackAboutInfo(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not update create stack overflow info listing in db",
		})
	}

	log.Printf("[controller/stackoverflow] successfully verified stack overflow account for coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.Stackoverflowuser{}
	return ctx.OK(res)
	// StackoverflowController_Update: end_implement
}

// Verify runs the verify action.
func (c *StackoverflowController) Verify(ctx *app.VerifyStackoverflowContext) error {
	// StackoverflowController_Verify: start_implement

	// Put your logic here

	user := &types.User{
		CognitoAuthUserID: ctx.Payload.UserID,
		Social: &types.Social{
			StackOverflow: &types.StackOverflow{
				Verification: &types.Verification{},
			},
		},
	}

	_, err := c.db.GetUserStackOverfloVerification(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	// simulate scraping profile
	// user.Social.StackOverflow.Verification.PostedVerificationCode = ""

	// check Stack Overflow for matching verification code
	_, err = stackoverflow.GetProfileByUserID(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get stack user profile",
		})
	}

	// secondary validation to see if codes match
	if !strings.Contains(user.Social.StackOverflow.Verification.PostedVerificationCode, user.Social.StackOverflow.Verification.ConfirmedVerificationCode) {
		log.Errorf("[controller/stackoverflow] %v", noVerifError)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "verification code does not match",
		})
	}

	// if no error, update verification field values
	user.Social.StackOverflow.Verification.Verified = true

	fmt.Println("\n")
	fmt.Printf("UPDATED USER: %v\n\n", user)

	_, err = c.db.UpdateStackVerificationCode(user)
	if err != nil {
		log.Errorf("[controller/stackoverflow] %v", err)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "could not get user verification info from db",
		})
	}

	res := &app.Stackoverflowuser{
		Verification: &app.Verification{
			Verified: user.Social.StackOverflow.Verification.Verified,
		},
	}
	return ctx.OK(res)
	// StackoverflowController_Verify: end_implement
}
