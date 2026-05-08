package tarpit

import (
	"log"
	"net/http"
	"tarpit/internal/detector"
	"time"
)

func TarpitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fingerprint := detector.Extract(r)

		score := fingerprint.Analyze()

		if score.Total >= 40 {
			select {
			case <-r.Context().Done():
				log.Println("bot disconnected")
				return
			case <-time.After(1 * time.Second):
			}
		}
		next.ServeHTTP(w, r)
	})
}
