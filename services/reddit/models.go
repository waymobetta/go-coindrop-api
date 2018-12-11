package reddit

import (
	"github.com/jzelinskie/geddit"
)

// AuthSessions info
type AuthSessions struct {
	OAuthSession  *geddit.OAuthSession `json:"oauthsession"`
	NoAuthSession *geddit.Session      `json:"noauthsession"`
}
