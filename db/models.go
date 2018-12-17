package db

import (
	"github.com/waymobetta/go-coindrop-api/services/keybase"
	"github.com/waymobetta/go-coindrop-api/verify"
)

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
	WalletAddress     string              `json:"wallet_address"`
	RedditData        RedditData          `json:"reddit_data"`
	KeybaseData       keybase.KeybaseData `json:"keybase_data"`
	StackOverflowData StackOverflowData   `json:"stackoverflow_data"`
}

/// REDDIT

// RedditData info
type RedditData struct {
	Username          string                  `json:"reddit_username"`
	LinkKarma         int                     `json:"link_karma"`
	CommentKarma      int                     `json:"comment_karma"`
	AccountCreatedUTC float64                 `json:"account_created_utc"`
	Trophies          []string                `json:"trophies"`
	Subreddits        []string                `json:"subreddits"`
	VerificationData  verify.VerificationData `json:"verification_data"`
}

/// STACK OVERFLOW

// StackOverflowData struct contains all essential info for Stack User
type StackOverflowData struct {
	ExchangeAccountID int                     `json:"exchange_account_id"`
	UserID            int                     `json:"user_id"`
	DisplayName       string                  `json:"display_name"`
	Accounts          []string                `json:"accounts"`
	Communities       []Community             `json:"communities"`
	VerificationData  verify.VerificationData `json:"verification_data"`
}

// Community struct contains info about the communities
type Community struct {
	Name          string         `json:"community_name"`
	Reputation    int            `json:"community_reputation"`
	QuestionCount int            `json:"community_question_count"`
	AnswerCount   int            `json:"community_answer_count"`
	BadgeCounts   map[string]int `json:"community_badge_counts"`
}