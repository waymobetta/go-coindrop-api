package types

import (
	"time"

	"github.com/waymobetta/go-coindrop-api/services/keybase"
)

// Users info
type Users struct {
	Users []User `json:"users"`
}

// User info
type User struct {
	ID                string  `json:"id"`
	UserID            string  `json:"user_id"`
	CognitoAuthUserID string  `json:"cognito_auth_user_id"`
	Wallet            *Wallet `json:"wallet"`
	Social            *Social `json:"social"`
}

// Profile info
type Profile struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Wallet ...
type Wallet struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Address  string `json:"address"`
	Type     string `json:"type"`
	Verified bool   `json:"verified"`
}

// Social ...
type Social struct {
	Reddit        *Reddit          `json:"reddit"`
	Keybase       *keybase.Keybase `json:"keybase"`
	StackOverflow *StackOverflow   `json:"stackoverflow"`
}

/// REDDIT

// Reddit info
type Reddit struct {
	ID           string         `json:"id"`
	Username     string         `json:"username"`
	LinkKarma    int            `json:"link_karma"`
	CommentKarma int            `json:"comment_karma"`
	Trophies     []string       `json:"trophies"`
	Subreddits   map[string]int `json:"subreddits"`
	Verification *Verification  `json:"verification"`
}

/// STACK OVERFLOW

// StackOverflow struct contains all essential info for Stack User
type StackOverflow struct {
	ID                string        `json:"id"`
	ExchangeAccountID int           `json:"exchange_account_id"`
	StackUserID       int           `json:"stack_user_id"`
	DisplayName       string        `json:"display_name"`
	Accounts          []string      `json:"accounts"`
	Communities       []Community   `json:"communities"`
	Verification      *Verification `json:"verification"`
}

// Verification struct contains all info for validation of account
type Verification struct {
	PostedVerificationCode    string `json:"posted_verification_code"`
	ConfirmedVerificationCode string `json:"confirmed_verification_code"`
	Verified                  bool   `json:"verified"`
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
	Author          string `json:"author"`
	BadgeData       *Badge `json:"badge_id"`
	Description     string `json:"description"`
	ID              string `json:"id"`
	Assigned        bool   `json:"assigned"`
	Completed       bool   `json:"completed"`
	Title           string `json:"title"`
	Token           string `json:"token"`
	TokenAllocation int    `json:"token_allocation"`
	Type            string `json:"type"`
	LogoURL         string `json:"logo_url"`
	QuizID          string `json:"quiz_id"`
}

// UserTask struct contains info about tasks for a specific user
type UserTask struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	TaskID    string `json:"task_id"`
	Completed bool   `json:"completed"`
}

// ListData struct contains info about a user's assigned and completed tasks
type ListData struct {
	AssignedTasks  []string `json:"assigned_tasks"`
	CompletedTasks []string `json:"completed_tasks"`
}

// TaskUser struct contains necessary info for helping manage user's task assignment/completions
type TaskUser struct {
	ID         string `json:"id"`
	TaskID     string `json:"task_id"`
	UserID     string `json:"user_id"`
	AuthUserID string `json:"auth_user_id"`
	Title      string `json:"title"`
}

// TaskUser2 ...
type TaskUser2 struct {
	ID     string `json:"id"`
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
}

// Quiz struct contains info about a quiz
type Quiz struct {
	ID     string        `json:"id"`
	Title  string        `json:"title"`
	UserID string        `json:"user_id"`
	Fields []*QuizFields `json:"fields"`
}

// QuizFields struct contains question and answer info about a quiz
type QuizFields struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// QuizResults struct contains info about a user's quiz results
type QuizResults struct {
	QuizID             string `json:"quiz_id"`
	TypeformFormID     string `json:"typeform_form_id"`
	TypeformQuizURL    string `json:"typeform_quiz_url"`
	UserID             string `json:"user_id"`
	QuestionsCorrect   int    `json:"questions_correct"`
	QuestionsIncorrect int    `json:"questions_incorrect"`
	QuizTaken          bool   `json:"quiz_taken"`
}

// AllQuizResults struct contains a slice of QuizResults structs
type AllQuizResults struct {
	QuizResults []QuizResults `json:"quiz_results_list"`
}

// Badge struct contains info about a badge
type Badge struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
}

// Transaction struct contains info about a send tx based on task completion
type Transaction struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
	Hash   string `json:"hash"`
}

// WalletVerification contains info used for Wallet verification purposes
type WalletVerification struct {
	Address   string `json:"address"`
	Message   string `json:"msg"`
	Signature string `json:"sig"`
	Version   string `json:"version"`
}

// TypeformWebHookResponse struct contains info from the TypeformWebHook response
type TypeformWebHookResponse struct {
	EventID      string `json:"event_id"`
	EventType    string `json:"event_type"`
	FormResponse struct {
		FormID      string    `json:"form_id"`
		Token       string    `json:"token"`
		LandedAt    time.Time `json:"landed_at"`
		SubmittedAt time.Time `json:"submitted_at"`
		Hidden      struct {
			Name   string `json:"name"`
			UserID string `json:"user_id"`
		} `json:"hidden"`
		Calculated struct {
			Score int `json:"score"`
		} `json:"calculated"`
		Definition struct {
			ID     string `json:"id"`
			Title  string `json:"title"`
			Fields []struct {
				ID         string `json:"id"`
				Title      string `json:"title"`
				Type       string `json:"type"`
				Ref        string `json:"ref"`
				Properties struct {
				} `json:"properties"`
				Choices []struct {
					ID    string `json:"id"`
					Label string `json:"label"`
				} `json:"choices"`
			} `json:"fields"`
		} `json:"definition"`
		Answers []struct {
			Type   string `json:"type"`
			Choice struct {
				Label string `json:"label"`
			} `json:"choice"`
			Field struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Ref  string `json:"ref"`
			} `json:"field"`
		} `json:"answers"`
	} `json:"form_response"`
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
