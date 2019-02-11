package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/throttled/throttled"
	"github.com/throttled/throttled/store/memstore"
)

// RateLimitHandler rate limit handler
func RateLimitHandler(router interface{}) http.Handler {
	store, err := memstore.New(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(40),
		MaxBurst: 10,
	}

	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)

	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}
	switch v := router.(type) {
	case *mux.Router:
		return httpRateLimiter.RateLimit(v)
	case http.Handler:
		return httpRateLimiter.RateLimit(v)
	default:
		log.Fatal("handler not supported")
	}

	return mux.NewRouter()
}
