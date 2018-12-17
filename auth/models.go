package auth

// Credentials is a struct of a user
type Credentials struct {
	ID       int    `json:"id"`
	Password string `json:"password", db:"password"`
	Email    string `json:"email", db:"email"`
}
