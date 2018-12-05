package main

import (
	"net/http"
)

// Route is a struct object containing route info
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a struct object containing a slice of Routes
type Routes []Route

var routes = Routes{
	Route{
		"UsersGet",
		"GET",
		"/api/v1/users",
		usersGet,
	},
	Route{
		"UserRemove",
		"POST",
		"/api/v1/removeuser",
		userRemove,
	},
	Route{
		"UserAdd",
		"POST",
		"/api/v1/adduser",
		userAdd,
	},
	Route{
		"UserGet",
		"POST",
		"/api/v1/user",
		userGet,
	},
	Route{
		"UpdateWallet",
		"POST",
		"/api/v1/updatewallet",
		walletUpdate,
	},
	Route{
		"RedditUpdate",
		"POST",
		"/api/v1/updateredditinfo",
		redditUpdate,
	},
	Route{
		"TwoFAUpdate",
		"POST",
		"/api/v1/updatetwofa",
		twoFAUpdate,
	},
	Route{
		"GenerateTwoEffEhCode",
		"POST",
		"/api/v1/generatetwoeffeh",
		generateTwoEffEhCode,
	},
	Route{
		"ValidateTwoEffEhCode",
		"POST",
		"/api/v1/validatetwoeffeh",
		validateTwoEffEhCode,
	},
	Route{
		"SignUp",
		"POST",
		"/api/v1/signup",
		signUp,
	},
	Route{
		"SignIn",
		"POST",
		"/api/v1/signin",
		signIn,
	},
}
