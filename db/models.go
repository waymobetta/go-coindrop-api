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
	ID                int                 `json:"id"`
	AuthUserID        string              `json:"auth_user_id"`
	WalletAddress     string              `json:"wallet_address"`
	RedditData        RedditData          `json:"reddit_data"`
	KeybaseData       keybase.KeybaseData `json:"keybase_data"`
	StackOverflowData StackOverflowData   `json:"stackoverflow_data"`
}

/// REDDIT

// RedditData info
type RedditData struct {
	ID                int                     `json:"id"`
	Username          string                  `json:"username"`
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
	ID                int                     `json:"id"`
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

// Tasks struct contains info about tasks
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// Task struct contains info about a task
type Task struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	Author          string `json:"author"`
	Description     string `json:"description"`
	Token           string `json:"token"`
	TokenAllocation int    `json:"token_allocation"`
	BadgeData       Badge  `json:"badge_data"`
}

// UserTask struct contains info about a task for a specific user
type UserTask struct {
	ID             int      `json:"id"`
	AuthUserID     string   `json:"auth_user_id"`
	AssignedTasks  []string `json:"assigned_tasks"`
	CompletedTasks []string `json:"completed_tasks"`
}

// Quiz struct contains info about a quiz
type Quiz struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	AuthUserID string   `json:"auth_user_id"`
	QuizInfo   QuizInfo `json:"quiz_info"`
}

// QuizInfo struct contains the list of QuizData objects
type QuizInfo struct {
	QuizData []QuizData `json:"quiz_data"`
}

// QuizData struct contains question and answer info about a quiz
type QuizData struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// QuizResults struct contains info about a user's quiz results
type QuizResults struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	AuthUserID         string `json:"auth_user_id"`
	QuestionsCorrect   int    `json:"questions_correct"`
	QuestionsIncorrect int    `json:"questions_incorrect"`
}

// Badge struct contains info about a badge
type Badge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Recipients  int    `json:"recipients"`
}
