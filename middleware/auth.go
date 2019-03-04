package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	authpkg "github.com/waymobetta/go-coindrop-api/auth"
	"github.com/waymobetta/go-coindrop-api/db"
)

// ErrAuthFailed means it wasn't able to authenticate the user making the request.
var ErrAuthFailed = goa.NewErrorClass("auth_failed", 401)

// ErrInternalServerError means the server erred
var ErrInternalServerError = goa.NewErrorClass("internal_server_error", 500)

// Auth authenticates user
func Auth(auth *authpkg.Auth, dbs *db.DB) goa.Middleware {
	usersMap := make(map[string]string)

	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Use ctx, rw and req - for example:
			// newctx := context.WithValue(ctx, "key", "value")

			// Example of setting header:
			// rw.Header().Set("X-Custom", "foo")

			authHeader := req.Header.Get("Authorization")
			bearer := strings.Split(authHeader, " ")
			if len(bearer) < 2 {
				log.Errorln("[middleware/auth] Auth token is missing")
				return ErrAuthFailed("Authentication failed")
			}
			jwtToken := bearer[1]
			token, err := auth.ParseJWT(jwtToken)
			if err != nil {
				log.Errorf("[middleware/auth] jwt parse error: %v\n", err)
				return ErrAuthFailed("Authentication failed")
			}
			if !token.Valid {
				log.Error("[middleware/auth] jwt invalid token\n")
				return ErrAuthFailed("Authentication failed")
			}

			icognitoUserID, err := auth.GetClaim(token, "sub")
			if err != nil {
				log.Errorf("[middleware/auth] could not retrieve token claim: %v\n", err)
				return ErrAuthFailed("Authentication failed")
			}
			cognitoUserID := icognitoUserID.(string)

			if cognitoUserID == "" {
				log.Error("[middleware/auth] could not retrieve user ID from token\n")
				return ErrAuthFailed("Authentication failed")
			}

			userID, ok := usersMap[cognitoUserID]
			if !ok {
				userID, err = dbs.GetUserIDByCognitoUserID(cognitoUserID)
				if err != nil {
					log.Errorf("[middleware/auth] could not retrieve user ID; %v\n", err)
					return ErrInternalServerError("Failed to get user")
				}
				usersMap[cognitoUserID] = userID
			}

			log.Printf("[middleware/auth] cognito auth ID: %s, user ID: %s\n", cognitoUserID, userID)

			ctx = context.WithValue(ctx, "authCognitoUserID", cognitoUserID)
			ctx = context.WithValue(ctx, "authUserID", userID)

			// Then call the next handler:
			return h(ctx, rw, req)
		}
	}
}
