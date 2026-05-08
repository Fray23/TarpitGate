package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"tarpit/internal/ratelimit"
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

	limiter := ratelimit.NewRateLimit(10, 10)
	handler := limiter.Middleware(tarpit.TarpitMiddleware(proxy))
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
