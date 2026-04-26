package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
	"tarpit/internal/tarpit"
)



func main() {
	proxy := &httputil.ReverseProxy{
	    Director: func(req *http.Request) {
	        req.URL.Scheme = "http"
	        req.URL.Host = "localhost:9090"
	    },
	}
	http.ListenAndServe(":8080", tarpit.TarpitMiddleware(proxy))
}
