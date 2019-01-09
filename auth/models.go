package auth

// Credentials is a struct of a user
type Credentials struct {
	ID         int    `json:"id", db:"id"`
	AuthUserID string `json:"auth_user_id", db:"auth_user_id"`
}
