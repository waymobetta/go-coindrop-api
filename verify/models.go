package verify

// VerificationData struct contains all info for validation of account
type VerificationData struct {
	PostedVerificationCode string `json:"posted_verification_code"`
	StoredVerificationCode string `json:"confirmed_verification_code"`
	IsVerified             bool   `json:"is_verified"`
}

type Verification2 struct {
	PostedVerificationCode    string `json:"posted_verification_code"`
	ConfirmedVerificationCode string `json:"confirmed_verification_code"`
	Verified                  bool   `json:"verified"`
}
