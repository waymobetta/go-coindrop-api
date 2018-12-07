package main

import "github.com/jzelinskie/geddit"

// TODO:
// organize into packages

// LOGIN

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
	WalletAddress     string            `json:"wallet_address"`
	RedditData        RedditData        `json:"reddit_data"`
	TwoFAData         TwoFAData         `json:"twofa_data"`
	KeybaseData       KeybaseData       `json:"keybase_data"`
	StackOverflowData StackOverflowData `json:"stackoverflow_data"`
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

// STACK OVERFLOW

// StackOverflowData struct contains all essential info for Stack User
type StackOverflowData struct {
	ExchangeAccountID int              `json:"exchange_account_id"`
	UserID            int              `json:"user_id"`
	DisplayName       string           `json:"display_name"`
	Communities       []Community      `json:"communities"`
	VerificationData  VerificationData `json:"verification_data"`
}

// VerificationData struct contains all info for validation of account
type VerificationData struct {
	PostedVerificationCode string `json:"posted_verification_code"`
	StoredVerificationCode string `json:"stored_verification_code"`
	IsVerified             bool   `json:"is_verified"`
}

// Community struct contains info about the communities
type Community struct {
	Name          string         `json:"community_name"`
	Reputation    int            `json:"community_reputation"`
	QuestionCount int            `json:"community_question_count"`
	AnswerCount   int            `json:"community_answer_count"`
	BadgeCounts   map[string]int `json:"community_badge_counts"`
}

// AboutProfileResponse struct contins profile info
type AboutProfileResponse struct {
	Items []struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		ViewCount               int    `json:"view_count"`
		DownVoteCount           int    `json:"down_vote_count"`
		UpVoteCount             int    `json:"up_vote_count"`
		AnswerCount             int    `json:"answer_count"`
		QuestionCount           int    `json:"question_count"`
		AccountID               int    `json:"account_id"`
		IsEmployee              bool   `json:"is_employee"`
		LastModifiedDate        int    `json:"last_modified_date"`
		LastAccessDate          int    `json:"last_access_date"`
		ReputationChangeYear    int    `json:"reputation_change_year"`
		ReputationChangeQuarter int    `json:"reputation_change_quarter"`
		ReputationChangeMonth   int    `json:"reputation_change_month"`
		ReputationChangeWeek    int    `json:"reputation_change_week"`
		ReputationChangeDay     int    `json:"reputation_change_day"`
		Reputation              int    `json:"reputation"`
		CreationDate            int    `json:"creation_date"`
		UserType                string `json:"user_type"`
		UserID                  int    `json:"user_id"`
		AboutMe                 string `json:"about_me"`
		Location                string `json:"location"`
		WebsiteURL              string `json:"website_url"`
		Link                    string `json:"link"`
		ProfileImage            string `json:"profile_image"`
		DisplayName             string `json:"display_name"`
	} `json:"items"`
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}

// AssociatedCommunitiesResponse struct contains info from
type AssociatedCommunitiesResponse struct {
	Items []struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		QuestionCount  int    `json:"question_count"`
		AnswerCount    int    `json:"answer_count"`
		LastAccessDate int    `json:"last_access_date"`
		CreationDate   int    `json:"creation_date"`
		AccountID      int    `json:"account_id"`
		Reputation     int    `json:"reputation"`
		UserID         int    `json:"user_id"`
		SiteURL        string `json:"site_url"`
		SiteName       string `json:"site_name"`
	} `json:"items"`
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}
