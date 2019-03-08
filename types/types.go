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

// Wallet ...
type Wallet struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Address string `json:"address"`
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
	ID           string        `json:"id"`
	Username     string        `json:"username"`
	LinkKarma    int           `json:"link_karma"`
	CommentKarma int           `json:"comment_karma"`
	Trophies     []string      `json:"trophies"`
	Subreddits   []string      `json:"subreddits"`
	Verification *Verification `json:"verification"`
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
	Recipients  int    `json:"recipients"`
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
			Name          string `json:"name"`
			Walletaddress string `json:"walletaddress"`
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
