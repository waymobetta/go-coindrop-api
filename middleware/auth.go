package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	authpkg "github.com/waymobetta/go-coindrop-api/auth"
)

// ErrAuthFailed means it wasn't able to authenticate the user making the request.
var ErrAuthFailed = goa.NewErrorClass("auth_failed", 401)

// ErrInternalServerError means the server erred
var ErrInternalServerError = goa.NewErrorClass("internal_server_error", 500)

// Auth authenticates user
func Auth(auth *authpkg.Auth) goa.Middleware {
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
			jwt := bearer[1]
			token, err := auth.ParseJWT(jwt)
			if err != nil {
				log.Errorf("[middleware/auth] jwt parse error: %v\n", err)
				return ErrAuthFailed("Authentication failed")
			}
			if !token.Valid {
				log.Error("[middleware/auth] jwt invalid token\n")
				return ErrAuthFailed("Authentication failed")
			}

			// Then call the next handler:
			return h(ctx, rw, req)
		}
	}
}
