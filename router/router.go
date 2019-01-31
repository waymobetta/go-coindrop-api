package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waymobetta/go-coindrop-api/handlers"
	"github.com/waymobetta/go-coindrop-api/logger"
	"github.com/waymobetta/gognito"
)

// Config ...
type Config struct {
	Region     string
	UserPoolID string
	Handlers   *handlers.Handlers
}

// NewRouter method creates a custom new mux router
func NewRouter(config *Config) *mux.Router {
	su := &gognito.ServiceUser{
		Region:     config.Region,
		UserPoolID: config.UserPoolID,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Use(su.AuthMiddleware)

	// NOTE: routes will be moved when using Goa
	hdlrs := config.Handlers
	for _, route := range []Route{
		{
			"HandleIndex",
			"GET",
			"/api/v1/test",
			hdlrs.HandleIndex,
		},
		{
			"GetWalletAddress",
			"POST",
			"/api/v1/getwallet",
			hdlrs.WalletGet,
		},
		{
			"AddUserID",
			"POST",
			"/api/v1/adduserid",
			hdlrs.UserIDAdd,
		},
		{
			"UsersGet",
			"GET",
			"/api/v1/getusers",
			hdlrs.UsersGet,
		},
		{
			"UserRemove",
			"POST",
			"/api/v1/removereddituser",
			hdlrs.RedditUserRemove,
		},
		{
			"UserAdd",
			"POST",
			"/api/v1/addreddituser",
			hdlrs.RedditUserAdd,
		},
		{
			"UserGet",
			"POST",
			"/api/v1/getreddituser",
			hdlrs.RedditUserGet,
		},
		{
			"UpdateWallet",
			"POST",
			"/api/v1/updatewallet",
			hdlrs.WalletUpdate,
		},
		{
			"RedditUpdate",
			"POST",
			"/api/v1/updateredditinfo",
			hdlrs.RedditUpdate,
		},
		{
			"UpdateRedditVerificationCode",
			"POST",
			"/api/v1/updateredditverificationcode",
			hdlrs.UpdateRedditVerificationCode,
		},
		{
			"GenerateRedditVerificationCode",
			"POST",
			"/api/v1/generateredditverificationcode",
			hdlrs.GenerateRedditVerificationCode,
		},
		{
			"ValidateRedditVerificationCode",
			"POST",
			"/api/v1/validateredditverificationcode",
			hdlrs.ValidateRedditVerificationCode,
		},
		{
			"StackUserAdd",
			"POST",
			"/api/v1/addstackuser",
			hdlrs.StackUserAdd,
		},
		{
			"StackUserGet",
			"POST",
			"/api/v1/getstackuser",
			hdlrs.StackUserGet,
		},
		{
			"StackUserUpdate",
			"POST",
			"/api/v1/updatestackinfo",
			hdlrs.StackUserUpdate,
		},
		{
			"GenerateStackVerificationCode",
			"POST",
			"/api/v1/generatestackverificationcode",
			hdlrs.GenerateStackVerificationCode,
		},
		{
			"ValidateStackVerificationCode",
			"POST",
			"/api/v1/validatestackverificationcode",
			hdlrs.ValidateStackVerificationCode,
		},
		{
			"TasksGet",
			"GET",
			"/api/v1/gettasks",
			hdlrs.TasksGet,
		},
		{
			"TaskAdd",
			"POST",
			"/api/v1/addtask",
			hdlrs.TaskAdd,
		},
		{
			"UserTasksGet",
			"POST",
			"/api/v1/getusertasks",
			hdlrs.UserTasksGet,
		},
		{
			"UserTaskAdd",
			"POST",
			"/api/v1/addusertask",
			hdlrs.UserTaskAdd,
		},
		{
			"UserTaskComplete",
			"POST",
			"/api/v1/completeusertask",
			hdlrs.UserTaskComplete,
		},
		{
			"UserTaskAssign",
			"POST",
			"/api/v1/assignusertask",
			hdlrs.UserTaskAssign,
		},
		{
			"ResultsPost",
			"POST",
			"/api/v1/postresults",
			hdlrs.ResultsPost,
		},
		{
			"ResultsGet",
			"POST",
			"/api/v1/getresults",
			hdlrs.ResultsGet,
		},
		{
			"ActionAdd",
			"POST",
			"/api/v1/addaction",
			hdlrs.ActionAdd,
		},
		{
			"ActionGet",
			"POST",
			"/api/v1/getaction",
			hdlrs.ActionGet,
		},
		{
			"QuizAdd",
			"POST",
			"/api/v1/addquiz",
			hdlrs.QuizAdd,
		},
		{
			"QuizGet",
			"POST",
			"/api/v1/getquiz",
			hdlrs.QuizGet,
		},
	} {
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
