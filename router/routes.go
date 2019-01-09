package router

import (
	"github.com/waymobetta/go-coindrop-api/auth"
	"github.com/waymobetta/go-coindrop-api/handlers"
)

// Routes is a struct object containing a slice of Routes
type Routes []Route

var routes = Routes{
	Route{
		"HandleIndex",
		"GET",
		"/api/v1/test",
		handlers.HandleIndex,
	},
	Route{
		"GetID",
		"POST",
		"/api/v1/getid",
		auth.GetID,
	},
	Route{
		"AddUserID",
		"POST",
		"/api/v1/adduserid",
		auth.AddUserID,
	},
	Route{
		"UsersGet",
		"GET",
		"/api/v1/getusers",
		handlers.UsersGet,
	},
	Route{
		"UserRemove",
		"POST",
		"/api/v1/removereddituser",
		handlers.UserRemove,
	},
	Route{
		"UserAdd",
		"POST",
		"/api/v1/addreddituser",
		handlers.UserAdd,
	},
	Route{
		"UserGet",
		"POST",
		"/api/v1/getreddituser",
		handlers.UserGet,
	},
	Route{
		"UpdateWallet",
		"POST",
		"/api/v1/updatewallet",
		handlers.WalletUpdate,
	},
	Route{
		"RedditUpdate",
		"POST",
		"/api/v1/updateredditinfo",
		handlers.RedditUpdate,
	},
	Route{
		"UpdateRedditVerificationCode",
		"POST",
		"/api/v1/updateredditverificationcode",
		handlers.UpdateRedditVerificationCode,
	},
	Route{
		"GenerateRedditVerificationCode",
		"POST",
		"/api/v1/generateredditverificationcode",
		handlers.GenerateRedditVerificationCode,
	},
	Route{
		"ValidateRedditVerificationCode",
		"POST",
		"/api/v1/validateredditverificationcode",
		handlers.ValidateRedditVerificationCode,
	},
	Route{
		"StackUserAdd",
		"POST",
		"/api/v1/addstackuser",
		handlers.StackUserAdd,
	},
	Route{
		"StackUserGet",
		"POST",
		"/api/v1/getstackuser",
		handlers.StackUserGet,
	},
	Route{
		"StackUserUpdate",
		"POST",
		"/api/v1/updatestackinfo",
		handlers.StackUserUpdate,
	},
	Route{
		"GenerateStackVerificationCode",
		"POST",
		"/api/v1/generatestackverificationcode",
		handlers.GenerateStackVerificationCode,
	},
	Route{
		"ValidateStackVerificationCode",
		"POST",
		"/api/v1/validatestackverificationcode",
		handlers.ValidateStackVerificationCode,
	},
	Route{
		"TasksGet",
		"GET",
		"/api/v1/gettasks",
		handlers.TasksGet,
	},
}
