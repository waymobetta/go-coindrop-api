package auth

// Credentials is a struct of a user
type Credentials struct {
	ID     int    `json:"id", db:"id"`
	UserID string `json:"user_id", db:"user_id"`
}
