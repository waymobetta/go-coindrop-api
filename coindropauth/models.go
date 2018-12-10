package coindropauth

// Credentials is a struct of a user
type Credentials struct {
	Password string `json:"password", db:"password"`
	Email    string `json:"email", db:"email"`
}
