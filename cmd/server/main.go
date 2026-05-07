package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"tarpit/internal/rate_limit"
	"tarpit/internal/tarpit"
)

func main() {
	fmt.Println("start tarpit server")
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:9090"
		},
	}

	limiter := ratelimit.NewLimiter(10, 10)
	rate_limit := tarpit.RateLimit{Limiter: limiter}
	handler := rate_limit.Middleware(tarpit.TarpitMiddleware(proxy))
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
