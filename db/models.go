package db

import (
	"time"

	"github.com/waymobetta/go-coindrop-api/services/keybase"
	"github.com/waymobetta/go-coindrop-api/verify"
)

// Users info
type Users struct {
	Users []User `json:"users"`
}

// User info
type User struct {
	ID            int             `json:"id"`
	AuthUserID    string          `json:"auth_user_id"`
	WalletAddress string          `json:"wallet_address"`
	Reddit        Reddit          `json:"reddit"`
	Keybase       keybase.Keybase `json:"keybase"`
	StackOverflow StackOverflow   `json:"stackoverflow"`
}

type User2 struct {
	ID            string  `json:"id"`
	AuthUserID    string  `json:"auth_user_id"`
	WalletAddress string  `json:"wallet_address"`
	Social        *Social `json:"social"`
}

/// SOCIAL
type Social struct {
	Reddit        *Reddit          `json:"reddit"`
	Keybase       *keybase.Keybase `json:"keybase"`
	StackOverflow *StackOverflow   `json:"stackoverflow"`
}

/// REDDIT

// Reddit info
type Reddit struct {
	ID                string                `json:"id"`
	Username          string                `json:"username"`
	LinkKarma         int                   `json:"link_karma"`
	CommentKarma      int                   `json:"comment_karma"`
	AccountCreatedUTC float64               `json:"account_created_utc"`
	Trophies          []string              `json:"trophies"`
	Subreddits        []string              `json:"subreddits"`
	Verification      *verify.Verification2 `json:"verification"`
}

/// STACK OVERFLOW

// StackOverflowData struct contains all essential info for Stack User
type StackOverflow struct {
	ID                string                `json:"id"`
	ExchangeAccountID int                   `json:"exchange_account_id"`
	UserID            int                   `json:"user_id"`
	DisplayName       string                `json:"display_name"`
	Accounts          []string              `json:"accounts"`
	Communities       []Community           `json:"communities"`
	Verification      *verify.Verification2 `json:"verification"`
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
	ID         int      `json:"id"`
	AuthUserID string   `json:"auth_user_id"`
	Assigned   string   `json:"assigned"`
	Completed  string   `json:"completed"`
	ListData   ListData `json:"list_data"`
}

type UserTask2 struct {
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
	AuthUserID string `json:"auth_user_id"`
	Title      string `json:"title"`
}

// Quiz struct contains info about a quiz
type Quiz struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	AuthUserID string    `json:"auth_user_id"`
	QuizInfo   *QuizInfo `json:"quiz_info"`
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
	HasTakenQuiz       bool   `json:"has_taken_quiz"`
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
		Definition  struct {
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
				AllowMultipleSelections bool `json:"allow_multiple_selections,omitempty"`
			} `json:"fields"`
		} `json:"definition"`
		Answers []struct {
			Type   string `json:"type"`
			Choice struct {
				Label string `json:"label"`
			} `json:"choice,omitempty"`
			Field struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Ref  string `json:"ref"`
			} `json:"field"`
			Choices struct {
				Labels []string `json:"labels"`
			} `json:"choices,omitempty"`
		} `json:"answers"`
	} `json:"form_response"`
}
