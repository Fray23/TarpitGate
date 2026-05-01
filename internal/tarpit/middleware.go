package tarpit

import (
	"fmt"
	"log"
	"net/http"
	"tarpit/internal/detector"
	"time"
)

func TarpitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fingerprint := detector.Extract(r)
		fmt.Println(fingerprint)
		fmt.Println(fingerprint.Analyze())

		select {
		case <-r.Context().Done():
			log.Println("bot disconnected")
			return
		case <-time.After(1 * time.Second):
			// return
			next.ServeHTTP(w, r)
		}
		next.ServeHTTP(w, r)
	})
}
