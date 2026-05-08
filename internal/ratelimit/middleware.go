package ratelimit

import (
	"net/http"
	"tarpit/internal/detector"
)


type RateLimit struct {
	Limiter *TokenBucketLimiter
}

func NewRateLimit(capacity, refillRate float64) *RateLimit {
	return &RateLimit{
		Limiter: NewLimiter(capacity, refillRate),
	}
}

func (rl *RateLimit) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fingerprint := detector.Extract(r)

		if !rl.Limiter.AllowIP(fingerprint.IP) {
			w.Header().Set("Retry-After", "1")
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})

}
