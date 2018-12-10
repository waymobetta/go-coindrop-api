package goreddit

import (
	"github.com/jzelinskie/geddit"
	"github.com/waymobetta/go-coindrop-api/gokeybase"
	"github.com/waymobetta/go-coindrop-api/gostackoverflow"
)

// REDDIT

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
	WalletAddress     string                            `json:"wallet_address"`
	RedditData        RedditData                        `json:"reddit_data"`
	TwoFAData         TwoFAData                         `json:"twofa_data"`
	KeybaseData       gokeybase.KeybaseData             `json:"keybase_data"`
	StackOverflowData gostackoverflow.StackOverflowData `json:"stackoverflow_data"`
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
