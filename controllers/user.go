package controllers

import (
	log "github.com/sirupsen/logrus"

	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/types"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	db *db.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, dbs *db.DB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		db:         dbs,
	}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	// initialize new user struct object
	user := new(types.User2)
	user.CognitoAuthUserID = ctx.Payload.CognitoAuthUserID

	// insert the AWS cognito user ID into the coindrop_auth table
	newUser, err := c.db.AddUserID(user)
	if err != nil {
		log.Errorf("[controller/user] %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not insert user to db",
		})
	}

	log.Printf("[controller/user] successfully added coindrop user: %v\n", user.CognitoAuthUserID)

	res := &app.User{
		ID:                "",
		CognitoAuthUserID: &newUser.CognitoAuthUserID,
		WalletAddress:     &newUser.Wallet.Address,
	}

	return ctx.OK(res)
	// UserController_Create: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	userID := ctx.UserID

	user, err := c.db.GetUser(userID)
	if err != nil {
		log.Errorf("[controller/user] failed to get user: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve user data",
		})
	}

	if user == nil {
		log.Errorf("[controller/user] user not found: %v", userID)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "user not found",
		})
	}

	res := &app.User{
		ID:                user.ID,
		CognitoAuthUserID: &user.CognitoAuthUserID,
		WalletAddress:     &user.Wallet.Address,
	}
	return ctx.OK(res)
	// UserController_Show: end_implement
}

// List runs the List action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here
	cognitoUserID := ctx.Params.Get("cognitoAuthUserId")
	userID, err := c.db.GetUserIDByCognitoUserID(cognitoUserID)
	if err != nil {
		log.Errorf("[controller/user] failed to get user id: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve user data",
		})
	}

	user, err := c.db.GetUser(userID)
	if err != nil {
		log.Errorf("[controller/user] failed to get user: %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not retrieve user data",
		})
	}

	if user == nil {
		log.Errorf("[controller/user] user not found: %v", userID)
		return ctx.NotFound(&app.StandardError{
			Code:    400,
			Message: "user not found",
		})
	}

	res := &app.User{
		ID:                user.ID,
		CognitoAuthUserID: &user.CognitoAuthUserID,
		WalletAddress:     &user.Wallet.Address,
	}
	return ctx.OK(res)
	// UserController_List: end_implement
}
