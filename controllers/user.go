package controllers

import (
	log "github.com/sirupsen/logrus"

	"github.com/goadesign/goa"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	db *db.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, db *db.DB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	// initialize new user struct object
	user := new(db.User)
	user.AuthUserID = ctx.Payload.CognitoAuthUserID

	// insert the AWS cognito user ID into the coindrop_auth table
	_, err := c.db.AddUserID(user)
	if err != nil {
		log.Errorf("[controller/user] %v", err)
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not insert user to db",
		})
	}

	log.Printf("[controller/user] successfully added coindrop user: %v\n", user.AuthUserID)

	res := &app.User{
		ID:                0,
		CognitoAuthUserID: &user.AuthUserID,
		WalletAddress:     &user.WalletAddress,
	}
	return ctx.OK(res)
	// UserController_Create: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UserController_Show: end_implement
}
