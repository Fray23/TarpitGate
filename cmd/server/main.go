package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
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
	http.ListenAndServe(":8080", tarpit.TarpitMiddleware(proxy))
}
