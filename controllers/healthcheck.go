package controllers

import (
	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/db"
)

// HealthcheckController implements the healthcheck resource.
type HealthcheckController struct {
	*goa.Controller
	db *db.DB
}

// NewHealthcheckController creates a healthcheck controller.
func NewHealthcheckController(service *goa.Service, dbs *db.DB) *HealthcheckController {
	return &HealthcheckController{
		Controller: service.NewController("HealthcheckController"),
		db:         dbs,
	}
}

// Show runs the show action.
func (c *HealthcheckController) Show(ctx *app.ShowHealthcheckContext) error {
	// HealthcheckController_Show: start_implement

	// Put your logic here
	err := c.db.Ping()
	if err != nil {
		log.Errorf("[controller/healthcheck] %v", err)
		return ctx.InternalServerError(&app.StandardError{
			Code:    500,
			Message: "could not ping db",
		})
	}

	res := &app.Healthcheck{
		Status: "OK",
	}
	return ctx.OK(res)
	// HealthcheckController_Show: end_implement
}
