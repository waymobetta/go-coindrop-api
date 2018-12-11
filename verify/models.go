package verify

// VerificationData struct contains all info for validation of account
type VerificationData struct {
	PostedVerificationCode string `json:"posted_verification_code"`
	StoredVerificationCode string `json:"stored_verification_code"`
	IsVerified             bool   `json:"is_verified"`
}
