package main

import "github.com/jzelinskie/geddit"

// Users info
type Users struct {
	Users []User `json:"users"`
}

// User info
type User struct {
	Info Info `json:"info"`
}

// Info info
type Info struct {
	WalletAddress string      `json:"wallet_address"`
	RedditData    RedditData  `json:"reddit_data"`
	TwoFAData     TwoFAData   `json:"twofa_data"`
	KeybaseData   KeybaseData `json:"keybase_data"`
}

// AuthSessions info
type AuthSessions struct {
	OAuthSession  *geddit.OAuthSession `json:"oauthsession"`
	NoAuthSession *geddit.Session      `json:"noauthsession"`
}

// RedditData info
type RedditData struct {
	Username          string   `json:"reddit_username"`
	LinkKarma         int      `json:"link_karma"`
	CommentKarma      int      `json:"comment_karma"`
	AccountCreatedUTC float64  `json:"account_created_utc"`
	Trophies          []string `json:"trophies"`
	Subreddits        []string `json:"subreddits"`
}

// TwoFAData info
type TwoFAData struct {
	PostedTwoFACode string `json:"posted_twofa_code"`
	StoredTwoFACode string `json:"stored_twofa_code"`
	IsValidated     bool   `json:"is_validated"`
}

// KeybaseData profile info
type KeybaseData struct {
	Bio                string `json:"bio"`
	Location           string `json:"location"`
	FullName           string `json:"full_name"`
	GithubUsername     string `json:"github_username"`
	TwitterUsername    string `json:"twitter_username"`
	HackerNewsUsername string `json:"hackernews_username"`
}

// Credentials is a struct of a user
type Credentials struct {
	Password string `json:"password", db:"password"`
	Email    string `json:"email", db:"email"`
}
