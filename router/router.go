package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/waymobetta/go-coindrop-api/auth"
	"github.com/waymobetta/go-coindrop-api/logger"
)

// NewRouter method creates a custom new mux router
func NewRouter() *mux.Router {

	region := os.Getenv("REACT_APP_AWS_COINDROP_COGNITO_REGION")
	userPoolID := os.Getenv("REACT_APP_AWS_COINDROP_COGNITO_USER_POOL_ID")

	su := &auth.ServiceUser{
		Region:     region,
		UserPoolID: userPoolID,
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Use(su.AuthMiddleware)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
