package gognito

import (
	"fmt"
	"net/http"
	"strings"
)

// AuthMiddleware method adds JWT middleware authentication to the API routes
func (s *ServiceUser) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]interface{})

		notAuth := []string{"/"}

		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json", s.Region, s.UserPoolID)

		// key to verify token against
		jwk := getJWK(jwkURL)

		tokenHeader := r.Header.Get("Authorization")
		// if token missing, return 403
		if tokenHeader == "" {
			response = message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application/json")
			respond(w, response)
			return
		}

		// token comes in format `Bearer {token-body}`
		split := strings.Split(tokenHeader, " ")

		if len(split) != 2 {
			response = message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application/json")
			respond(w, response)
			return
		}

		tokenString := split[1]

		token, err := validateToken(tokenString, s.Region, s.UserPoolID, jwk)
		if err != nil || !token.Valid {
			response = message(false, "Invalid Token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application/json")
			respond(w, response)
		} else {
			next.ServeHTTP(w, r)
			return
		}
	})
}
